package install

import (
	stdembed "embed"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Ulysses-Alv/unity-expert/internal/backup"
	configpkg "github.com/Ulysses-Alv/unity-expert/internal/config"
	"github.com/Ulysses-Alv/unity-expert/internal/embed"
)

// PreconditionError is returned when a precondition check fails.
type PreconditionError struct {
	Message string
	Hint    string
}

func (e *PreconditionError) Error() string {
	return e.Message
}

// BackupError is returned when backup creation fails.
type BackupError struct {
	Path string
	Err  error
}

func (e *BackupError) Error() string {
	return fmt.Sprintf("backup failed at %s: %v", e.Path, e.Err)
}

// ApplyError is returned when an apply operation fails.
type ApplyError struct {
	Operation string
	Path      string
	Err       error
}

func (e *ApplyError) Error() string {
	return fmt.Sprintf("%s failed for %s: %v", e.Operation, e.Path, e.Err)
}

// RollbackError is returned when rollback fails.
type RollbackError struct {
	BackupPath string
	Err       error
}

func (e *RollbackError) Error() string {
	return fmt.Sprintf("rollback failed from %s: %v", e.BackupPath, e.Err)
}

// Pipeline orchestrates install/sync operations with backup and rollback.
type Pipeline struct {
	backupDir   string
	statePath   string
	configPath  string
	promptsDir  string
	skillsDir   string
	dryRun      bool
	force       bool
	provider    string
	agents      []string
	backupPath  string
	createdFiles []string // Track files created during Apply for rollback
}

// NewPipeline creates a new Pipeline instance.
func NewPipeline(opts InstallOptions) *Pipeline {
	return &Pipeline{
		backupDir:   configpkg.DefaultBackupsPath(),
		statePath:   configpkg.DefaultStatePath(),
		configPath:  configpkg.DefaultOpenCodePath(),
		promptsDir:  filepath.Join(configpkg.DefaultOpenCodePath(), "..", "prompts", "unity"),
		skillsDir:   filepath.Join(configpkg.DefaultOpenCodePath(), "..", "skills", "unity-6000"),
		dryRun:      opts.DryRun,
		force:       opts.Force,
		provider:    opts.Provider,
		agents:      opts.Agents,
		createdFiles: []string{},
	}
}

// Prepare validates preconditions and creates a backup.
func (p *Pipeline) Prepare() error {
	// Check preconditions
	if err := p.checkPreconditions(); err != nil {
		return err
	}

	// Check for existing agents (skip unless force)
	agents, err := p.checkExistingAgents()
	if err != nil {
		return err
	}
	if len(agents) > 0 && !p.force {
		fmt.Printf("[INFO] Unity agents already installed (%d found). Use --force to reinstall.\n", len(agents))
		return nil
	}

	// Create backup
	if !p.dryRun {
		backupPath, err := backup.CreateBackup(p.configPath, p.backupDir, "install")
		if err != nil {
			return &BackupError{Path: p.backupDir, Err: err}
		}
		p.backupPath = backupPath
		fmt.Printf("[INFO] Backup created at %s\n", backupPath)
	} else {
		fmt.Println("[DRY RUN] Would create backup of opencode.json")
	}

	return nil
}

// Apply performs the actual installation.
func (p *Pipeline) Apply() error {
	if p.dryRun {
		fmt.Println("[DRY RUN] Would perform the following actions:")
		fmt.Println("  1. Merge agents from agents.json into opencode.json")
		fmt.Println("  2. Apply provider preset models from presets.json")
		fmt.Println("  3. Copy prompts to prompts/unity/")
		fmt.Println("  4. Copy skills to skills/unity-6000/")
		return nil
	}

	fmt.Println("[INFO] Applying installation...")

	// 1. Read embedded agents.json
	agentsData, err := embed.EmbeddedAgentsJSON()
	if err != nil {
		return &ApplyError{Operation: "read agents.json", Path: "embedded", Err: err}
	}

	// 2. Read embedded presets.json
	presetsData, err := embed.EmbeddedPresetsJSON()
	if err != nil {
		return &ApplyError{Operation: "read presets.json", Path: "embedded", Err: err}
	}

	// Parse presets to get provider models
	var presets map[string]struct {
		Models map[string]string `json:"models"`
	}
	if err := json.Unmarshal(presetsData, &presets); err != nil {
		return &ApplyError{Operation: "parse presets", Path: "embedded", Err: err}
	}

	// Get models for the selected provider
	providerModels := make(map[string]string)
	if providerPreset, ok := presets[p.provider]; ok {
		providerModels = providerPreset.Models
	}

	// 3. Read opencode.json
	configData, err := os.ReadFile(p.configPath)
	if err != nil {
		return &ApplyError{Operation: "read opencode.json", Path: p.configPath, Err: err}
	}

	var opencodeConfig map[string]interface{}
	if err := json.Unmarshal(configData, &opencodeConfig); err != nil {
		return &ApplyError{Operation: "parse opencode.json", Path: p.configPath, Err: err}
	}

	// 4. Merge agents.json into opencode.json
	var embeddedAgents map[string]interface{}
	if err := json.Unmarshal(agentsData, &embeddedAgents); err != nil {
		return &ApplyError{Operation: "parse agents.json", Path: "embedded", Err: err}
	}

	// Merge agents, preserving user modifications
	opencodeConfig = configpkg.MergeAgents(opencodeConfig, embeddedAgents, providerModels, p.force, p.agents)

	agentMap := opencodeConfig["agent"].(map[string]interface{})

	// 5. Write updated opencode.json
	if err := configpkg.WriteOpenCodeConfig(p.configPath, opencodeConfig); err != nil {
		return &ApplyError{Operation: "write opencode.json", Path: p.configPath, Err: err}
	}
	fmt.Printf("[INFO] Updated opencode.json with %d agents\n", len(agentMap))

	// 6. Copy prompts to prompts/unity/
	if err := p.copyPrompts(); err != nil {
		return &ApplyError{Operation: "copy prompts", Path: p.promptsDir, Err: err}
	}

	// 7. Copy skills to skills/unity-6000/
	if err := p.copySkills(); err != nil {
		return &ApplyError{Operation: "copy skills", Path: p.skillsDir, Err: err}
	}

	return nil
}

// copyPrompts copies embedded prompts to the prompts/unity directory.
func (p *Pipeline) copyPrompts() error {
	// Ensure target directory exists
	if err := os.MkdirAll(p.promptsDir, 0755); err != nil {
		return fmt.Errorf("failed to create prompts directory: %w", err)
	}

	// Open the embedded prompts directory
	promptsFS := embed.EmbeddedPromptsFS()
	entries, err := promptsFS.ReadDir(".")
	if err != nil {
		return fmt.Errorf("failed to read embedded prompts: %w", err)
	}

	copied := 0
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		if !strings.HasSuffix(name, ".md") {
			continue
		}

		// Read embedded file
		data, err := promptsFS.ReadFile(name)
		if err != nil {
			return fmt.Errorf("failed to read embedded prompt %s: %w", name, err)
		}

		// Write to target
		targetPath := filepath.Join(p.promptsDir, name)
		if err := os.WriteFile(targetPath, data, 0644); err != nil {
			return fmt.Errorf("failed to write prompt %s: %w", name, err)
		}

		p.createdFiles = append(p.createdFiles, targetPath)
		copied++
	}

	fmt.Printf("[INFO] Copied %d prompt files to %s\n", copied, p.promptsDir)
	return nil
}

// copySkills copies embedded skills to the skills/unity-6000 directory.
func (p *Pipeline) copySkills() error {
	// Ensure target directory exists
	if err := os.MkdirAll(p.skillsDir, 0755); err != nil {
		return fmt.Errorf("failed to create skills directory: %w", err)
	}

	// Open the embedded skills directory
	skillsFS := embed.EmbeddedSkillsFS()

	// Read all skill folders from embedded skills
	entries, err := skillsFS.ReadDir(".")
	if err != nil {
		return fmt.Errorf("failed to read embedded skills directory: %w", err)
	}

	copied := 0
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		skillName := entry.Name()
		sourceDir := skillName

		// Copy entire directory recursively
		if err := p.copyDir(skillsFS, sourceDir, p.skillsDir, skillName); err != nil {
			return fmt.Errorf("failed to copy skill %s: %w", skillName, err)
		}

		copied++
	}

	fmt.Printf("[INFO] Copied %d skill folders to %s\n", copied, p.skillsDir)
	return nil
}

// copyDir recursively copies a directory from an embed.FS to the filesystem.
func (p *Pipeline) copyDir(efs stdembed.FS, sourceDir, targetBase, skillName string) error {
	targetDir := filepath.Join(targetBase, skillName)
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", targetDir, err)
	}
	p.createdFiles = append(p.createdFiles, targetDir)

	entries, err := efs.ReadDir(sourceDir)
	if err != nil {
		return fmt.Errorf("failed to read directory %s: %w", sourceDir, err)
	}

	for _, entry := range entries {
		sourcePath := filepath.Join(sourceDir, entry.Name())
		targetPath := filepath.Join(targetDir, entry.Name())

		if entry.IsDir() {
			if err := p.copyDir(efs, sourcePath, targetDir, entry.Name()); err != nil {
				return err
			}
		} else {
			data, err := efs.ReadFile(sourcePath)
			if err != nil {
				return fmt.Errorf("failed to read file %s: %w", sourcePath, err)
			}

			if err := os.WriteFile(targetPath, data, 0644); err != nil {
				return fmt.Errorf("failed to write file %s: %w", targetPath, err)
			}
			p.createdFiles = append(p.createdFiles, targetPath)
		}
	}

	return nil
}

// Verify checks that all expected files exist and JSON is valid.
func (p *Pipeline) Verify() error {
	if p.dryRun {
		fmt.Println("[DRY RUN] Would verify installation")
		return nil
	}

	fmt.Println("[INFO] Verifying installation...")

	// Verify opencode.json is valid JSON and has agents
	data, err := os.ReadFile(p.configPath)
	if err != nil {
		return fmt.Errorf("failed to read opencode.json: %w", err)
	}

	var cfg map[string]interface{}
	if err := json.Unmarshal(data, &cfg); err != nil {
		return fmt.Errorf("opencode.json has invalid JSON: %w", err)
	}

	agentMap, ok := cfg["agent"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("opencode.json missing agent map")
	}

	// Count unity agents
	unityAgents := []string{}
	for name := range agentMap {
		if strings.HasPrefix(name, "unity-") {
			unityAgents = append(unityAgents, name)
		}
	}

	if len(unityAgents) == 0 {
		return fmt.Errorf("no unity agents found in opencode.json")
	}
	fmt.Printf("[OK] Found %d unity agents in opencode.json\n", len(unityAgents))

	// Verify prompts directory
	promptsFS := embed.EmbeddedPromptsFS()
	promptEntries, err := promptsFS.ReadDir(".")
	if err != nil {
		return fmt.Errorf("failed to read embedded prompts: %w", err)
	}

	if len(promptEntries) == 0 {
		return fmt.Errorf("no embedded prompts found")
	}
	fmt.Printf("[OK] Embedded prompts verified (%d files)\n", len(promptEntries))

	// Verify skills directory
	skillsFS := embed.EmbeddedSkillsFS()
	skillEntries, err := skillsFS.ReadDir(".")
	if err != nil {
		return fmt.Errorf("failed to read embedded skills: %w", err)
	}

	if len(skillEntries) == 0 {
		return fmt.Errorf("no embedded skills found")
	}
	fmt.Printf("[OK] Embedded skills verified (%d folders)\n", len(skillEntries))

	// Write state.json atomically
	state := configpkg.State{
		Version:          1,
		InstalledAt:      time.Now().Format(time.RFC3339),
		UpdatedAt:        time.Now().Format(time.RFC3339),
		InstalledAgents:  unityAgents,
		Provider:         p.provider,
		ModelAssignments: make(map[string]string),
	}

	if err := configpkg.WriteState(p.statePath, &state); err != nil {
		return fmt.Errorf("failed to write state.json: %w", err)
	}

	fmt.Println("[OK] Installation verified successfully")
	return nil
}

// Rollback restores opencode.json from backup and cleans up created files.
func (p *Pipeline) Rollback() error {
	// First, restore the backup
	if p.backupPath != "" {
		if err := backup.RestoreFromBackup(p.backupPath, p.configPath); err != nil {
			return &RollbackError{BackupPath: p.backupPath, Err: err}
		}
		fmt.Printf("[OK] Restored opencode.json from backup\n")
	}

	// Delete created files (in reverse order - directories after files)
	for i := len(p.createdFiles) - 1; i >= 0; i-- {
		path := p.createdFiles[i]
		info, err := os.Stat(path)
		if err != nil {
			// File may not exist, skip
			continue
		}

		if info.IsDir() {
			os.Remove(path)
		} else {
			os.Remove(path)
		}
		fmt.Printf("[INFO] Cleaned up %s\n", path)
	}

	return nil
}

// Execute runs the full pipeline.
func (p *Pipeline) Execute() error {
	if err := p.Prepare(); err != nil {
		return err
	}

	if p.dryRun {
		return nil
	}

	if err := p.Apply(); err != nil {
		p.Rollback()
		return err
	}

	return p.Verify()
}

// checkPreconditions validates that prerequisites are met.
func (p *Pipeline) checkPreconditions() error {
	// Check opencode.json exists
	if _, err := os.Stat(p.configPath); os.IsNotExist(err) {
		home, _ := os.UserHomeDir()
		return &PreconditionError{
			Message: fmt.Sprintf("OpenCode config not found at: %s", p.configPath),
			Hint:    fmt.Sprintf("Please ensure OpenCode is installed. Visit: https://opencode.ai\nConfig expected at: %s", filepath.Join(home, ".config", "opencode", "opencode.json")),
		}
	}
	return nil
}

// checkExistingAgents returns the list of existing Unity agents.
func (p *Pipeline) checkExistingAgents() ([]string, error) {
	config, err := configpkg.ReadOpenCodeConfig(p.configPath)
	if err != nil {
		return nil, err
	}

	agents := []string{}
	if agentMap, ok := config["agent"].(map[string]interface{}); ok {
		for name := range agentMap {
			if strings.HasPrefix(name, "unity-") {
				agents = append(agents, name)
			}
		}
	}

	return agents, nil
}
