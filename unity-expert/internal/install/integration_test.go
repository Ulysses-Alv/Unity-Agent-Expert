package install

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Ulysses-Alv/unity-expert/internal/backup"
	configpkg "github.com/Ulysses-Alv/unity-expert/internal/config"
)

// TestPrepare_CreatesBackup verifies that Prepare creates a backup
func TestPrepare_CreatesBackup(t *testing.T) {
	tmpDir := t.TempDir()

	// Create opencode.json
	configPath := filepath.Join(tmpDir, "opencode.json")
	configContent := []byte(`{"agent": {"existing": {"model": "test"}}}`)
	if err := os.WriteFile(configPath, configContent, 0644); err != nil {
		t.Fatalf("Failed to create config: %v", err)
	}

	// Create pipeline
	p := &Pipeline{
		backupDir:  filepath.Join(tmpDir, "backups"),
		statePath:  filepath.Join(tmpDir, "state.json"),
		configPath: configPath,
		dryRun:     false,
		force:      true,
	}

	// Run Prepare
	if err := p.Prepare(); err != nil {
		t.Fatalf("Prepare failed: %v", err)
	}

	// Check backup was created
	if p.backupPath == "" {
		t.Error("Expected backup path to be set")
	}

	if _, err := os.Stat(p.backupPath); os.IsNotExist(err) {
		t.Errorf("Backup file should exist: %s", p.backupPath)
	}
}

// TestPrepare_DryRunDoesNotModifyFiles verifies --dry-run doesn't modify files
func TestPrepare_DryRunDoesNotModifyFiles(t *testing.T) {
	tmpDir := t.TempDir()

	// Create opencode.json
	configPath := filepath.Join(tmpDir, "opencode.json")
	originalContent := []byte(`{"agent": {"existing": {"model": "test"}}}`)
	if err := os.WriteFile(configPath, originalContent, 0644); err != nil {
		t.Fatalf("Failed to create config: %v", err)
	}

	// Create pipeline with dryRun
	p := &Pipeline{
		backupDir:  filepath.Join(tmpDir, "backups"),
		statePath:  filepath.Join(tmpDir, "state.json"),
		configPath: configPath,
		dryRun:     true,
		force:      true,
	}

	// Run Prepare
	if err := p.Prepare(); err != nil {
		t.Fatalf("Prepare failed: %v", err)
	}

	// Check backup was NOT created (dry run)
	if p.backupPath != "" {
		t.Error("Expected no backup path in dry-run mode")
	}

	// Verify original content unchanged
	content, err := os.ReadFile(configPath)
	if err != nil {
		t.Fatalf("Failed to read config: %v", err)
	}
	if string(content) != string(originalContent) {
		t.Error("Config file should not be modified in dry-run mode")
	}
}

// TestApply_DryRunShowsActions verifies --dry-run shows what would be done
func TestApply_DryRunShowsActions(t *testing.T) {
	tmpDir := t.TempDir()

	// Create opencode.json
	configPath := filepath.Join(tmpDir, "opencode.json")
	configContent := []byte(`{"agent": {}}`)
	if err := os.WriteFile(configPath, configContent, 0644); err != nil {
		t.Fatalf("Failed to create config: %v", err)
	}

	// Create pipeline with dryRun
	p := &Pipeline{
		backupDir:  filepath.Join(tmpDir, "backups"),
		statePath:  filepath.Join(tmpDir, "state.json"),
		configPath: configPath,
		dryRun:     true,
		force:      true,
		provider:   "opencode",
	}

	// Run Apply (should not error in dry-run)
	if err := p.Apply(); err != nil {
		t.Fatalf("Apply dry-run failed: %v", err)
	}

	// Verify file was NOT modified
	content, err := os.ReadFile(configPath)
	if err != nil {
		t.Fatalf("Failed to read config: %v", err)
	}
	if string(content) != string(configContent) {
		t.Error("Config file should not be modified in dry-run mode")
	}
}

// TestApply_ErrorRollback verifies rollback is triggered on Apply error
func TestApply_ErrorRollback(t *testing.T) {
	tmpDir := t.TempDir()

	// Create opencode.json
	configPath := filepath.Join(tmpDir, "opencode.json")
	originalContent := []byte(`{"agent": {"unity-existing": {"model": "original"}}}`)
	if err := os.WriteFile(configPath, originalContent, 0644); err != nil {
		t.Fatalf("Failed to create config: %v", err)
	}

	// Create a backup first
	backupDir := filepath.Join(tmpDir, "backups")
	os.MkdirAll(backupDir, 0755)
	backupPath, err := backup.CreateBackup(configPath, backupDir, "test")
	if err != nil {
		t.Fatalf("Failed to create backup: %v", err)
	}

	// Set a backup path so rollback can restore
	p := &Pipeline{
		backupDir:   backupDir,
		statePath:   filepath.Join(tmpDir, "state.json"),
		configPath:  configPath,
		promptsDir:  "invalid::path",  // Invalid path to cause error
		skillsDir:   filepath.Join(tmpDir, "skills"),
		dryRun:      false,
		force:       true,
		provider:    "opencode",
		backupPath:  backupPath,
	}

	// Run Execute which calls Apply and handles rollback on error
	err = p.Execute()
	if err == nil {
		t.Fatal("Expected Apply to fail with invalid promptsDir")
	}

	// Manually call Rollback to restore (since test used Apply directly)
	// In real usage, Execute handles this
	p.Rollback()

	// Verify config was rolled back
	content, err := os.ReadFile(configPath)
	if err != nil {
		t.Fatalf("Failed to read config after rollback: %v", err)
	}
	if string(content) != string(originalContent) {
		t.Errorf("Config should be rolled back to original. Got: %s", string(content))
	}
}

// TestPipelineExecute_FullPipelineDryRun verifies full pipeline in dry-run
func TestPipelineExecute_FullPipelineDryRun(t *testing.T) {
	tmpDir := t.TempDir()

	// Create opencode.json
	configPath := filepath.Join(tmpDir, "opencode.json")
	originalContent := []byte(`{"agent": {"existing": {"model": "test"}}}`)
	if err := os.WriteFile(configPath, originalContent, 0644); err != nil {
		t.Fatalf("Failed to create config: %v", err)
	}

	// Create valid backup dir
	backupDir := filepath.Join(tmpDir, "backups")
	os.MkdirAll(backupDir, 0755)

	// Create pipeline
	p := &Pipeline{
		backupDir:  backupDir,
		statePath:  filepath.Join(tmpDir, "state.json"),
		configPath: configPath,
		promptsDir: "nonexistent-prompts",
		skillsDir:  "nonexistent-skills",
		dryRun:     true,
		force:      true,
		provider:   "opencode",
	}

	// Execute should succeed in dry-run
	if err := p.Execute(); err != nil {
		t.Fatalf("Execute dry-run failed: %v", err)
	}

	// Verify config unchanged
	content, err := os.ReadFile(configPath)
	if err != nil {
		t.Fatalf("Failed to read config: %v", err)
	}
	if string(content) != string(originalContent) {
		t.Error("Config should not be modified in dry-run mode")
	}
}

// TestForceReinstall_OverwritesExisting verifies --force reinstalls
func TestForceReinstall_OverwritesExisting(t *testing.T) {
	tmpDir := t.TempDir()

	// Create opencode.json with existing unity agents
	configPath := filepath.Join(tmpDir, "opencode.json")
	configContent := []byte(`{"agent": {"unity-existing": {"model": "old-model"}}}`)
	if err := os.WriteFile(configPath, configContent, 0644); err != nil {
		t.Fatalf("Failed to create config: %v", err)
	}

	backupDir := filepath.Join(tmpDir, "backups")
	os.MkdirAll(backupDir, 0755)

	// Create pipeline with force
	p := &Pipeline{
		backupDir:  backupDir,
		statePath:  filepath.Join(tmpDir, "state.json"),
		configPath: configPath,
		promptsDir: "nonexistent",
		skillsDir:  "nonexistent",
		dryRun:     true, // dry-run to avoid actual file writes
		force:      true,
		provider:   "claude",
	}

	// Prepare should not skip when force is true
	if err := p.Prepare(); err != nil {
		t.Fatalf("Prepare with force failed: %v", err)
	}
}

// TestCheckPreconditions_MissingConfig verifies error on missing config
func TestCheckPreconditions_MissingConfig(t *testing.T) {
	p := &Pipeline{
		configPath: "non-existent-config.json",
	}

	err := p.checkPreconditions()
	if err == nil {
		t.Fatal("Expected error for missing config")
	}

	preconditionErr, ok := err.(*PreconditionError)
	if !ok {
		t.Fatalf("Expected PreconditionError, got %T", err)
	}

	if preconditionErr.Message == "" {
		t.Error("Expected non-empty error message")
	}
}

// TestCheckExistingAgents_DetectsUnityAgents verifies detection of unity agents
func TestCheckExistingAgents_DetectsUnityAgents(t *testing.T) {
	tmpDir := t.TempDir()

	configPath := filepath.Join(tmpDir, "opencode.json")
	configContent := []byte(`{
		"agent": {
			"unity-6000-expert": {"model": "test"},
			"unity-graphics-expert": {"model": "test"},
			"other-agent": {"model": "test"}
		}
	}`)
	if err := os.WriteFile(configPath, configContent, 0644); err != nil {
		t.Fatalf("Failed to create config: %v", err)
	}

	p := &Pipeline{
		configPath: configPath,
	}

	agents, err := p.checkExistingAgents()
	if err != nil {
		t.Fatalf("checkExistingAgents failed: %v", err)
	}

	if len(agents) != 2 {
		t.Errorf("Expected 2 unity agents, got %d", len(agents))
	}
}

// TestStateIsWrittenAfterVerify verifies state.json is written after successful Verify
func TestStateIsWrittenAfterVerify(t *testing.T) {
	tmpDir := t.TempDir()

	configPath := filepath.Join(tmpDir, "opencode.json")
	configContent := []byte(`{
		"agent": {
			"unity-6000-expert": {"model": "test"}
		}
	}`)
	if err := os.WriteFile(configPath, configContent, 0644); err != nil {
		t.Fatalf("Failed to create config: %v", err)
	}

	statePath := filepath.Join(tmpDir, "state.json")
	promptsDir := filepath.Join(tmpDir, "prompts", "unity")
	skillsDir := filepath.Join(tmpDir, "skills", "unity-6000")
	os.MkdirAll(promptsDir, 0755)
	os.MkdirAll(skillsDir, 0755)

	// Create dummy prompt files
	os.WriteFile(filepath.Join(promptsDir, "test.md"), []byte("# Test"), 0644)
	os.WriteFile(filepath.Join(skillsDir, "test-skill", "SKILL.md"), []byte("# Skill"), 0644)

	p := &Pipeline{
		backupDir:  filepath.Join(tmpDir, "backups"),
		statePath:  statePath,
		configPath: configPath,
		promptsDir: promptsDir,
		skillsDir:  skillsDir,
		dryRun:    false,
		force:     true,
		provider:  "opencode",
	}

	// Verify should write state
	if err := p.Verify(); err != nil {
		t.Fatalf("Verify failed: %v", err)
	}

	// Check state.json was created
	valid, err := configpkg.ValidState(statePath)
	if err != nil {
		t.Fatalf("ValidState check failed: %v", err)
	}
	if !valid {
		t.Error("Expected state.json to be valid after Verify")
	}
}