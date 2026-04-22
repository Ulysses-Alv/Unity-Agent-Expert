package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	configpkg "github.com/Ulysses-Alv/unity-expert/internal/config"
	"github.com/Ulysses-Alv/unity-expert/internal/embed"
)

// CheckResult represents the result of a single health check.
type CheckResult struct {
	Name    string
	Status  string // "PASS", "FAIL", "WARN"
	Message string
}

// RunDoctor performs health checks on the Unity Expert installation.
func RunDoctor() error {
	fmt.Println()
	fmt.Println("Unity Expert Doctor — Installation Health Check")
	fmt.Println("================================================")
	fmt.Println()

	checks := []CheckResult{}
	passCount := 0
	warnCount := 0
	failCount := 0

	// Check 1: opencode.json exists
	result := checkOpenCodeConfig()
	checks = append(checks, result)
	switch result.Status {
	case "PASS":
		passCount++
	case "WARN":
		warnCount++
	case "FAIL":
		failCount++
	}

	// Check 2: Unity agents installed
	result = checkUnityAgents()
	checks = append(checks, result)
	switch result.Status {
	case "PASS":
		passCount++
	case "WARN":
		warnCount++
	case "FAIL":
		failCount++
	}

	// Check 3: Prompts present
	result = checkPrompts()
	checks = append(checks, result)
	switch result.Status {
	case "PASS":
		passCount++
	case "WARN":
		warnCount++
	case "FAIL":
		failCount++
	}

	// Check 4: Skills present
	result = checkSkills()
	checks = append(checks, result)
	switch result.Status {
	case "PASS":
		passCount++
	case "WARN":
		warnCount++
	case "FAIL":
		failCount++
	}

	// Check 5: State.json exists
	result = checkStateFile()
	checks = append(checks, result)
	switch result.Status {
	case "PASS":
		passCount++
	case "WARN":
		warnCount++
	case "FAIL":
		failCount++
	}

	// Print individual results
	fmt.Println()
	for _, check := range checks {
		fmt.Printf("[%s] %s: %s\n", check.Status, check.Name, check.Message)
	}

	// Print summary
	fmt.Println()
	fmt.Println("================================================")
	fmt.Printf("Summary: %d checks passed, %d warnings, %d failures\n", passCount, warnCount, failCount)
	fmt.Println("================================================")

	// Return exit code based on failures
	if failCount > 0 {
		return fmt.Errorf("health checks failed")
	}

	return nil
}

// checkOpenCodeConfig verifies opencode.json exists.
func checkOpenCodeConfig() CheckResult {
	configPath := configpkg.DefaultOpenCodePath()

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return CheckResult{
			Name:    "OpenCode Config",
			Status:  "FAIL",
			Message: "opencode.json not found at " + configPath,
		}
	}

	// Verify it's valid JSON
	data, err := os.ReadFile(configPath)
	if err != nil {
		return CheckResult{
			Name:    "OpenCode Config",
			Status:  "FAIL",
			Message: "cannot read opencode.json",
		}
	}

	var jsonRaw json.RawMessage
	if err := json.Unmarshal(data, &jsonRaw); err != nil {
		return CheckResult{
			Name:    "OpenCode Config",
			Status:  "FAIL",
			Message: "opencode.json is not valid JSON",
		}
	}

	return CheckResult{
		Name:    "OpenCode Config",
		Status:  "PASS",
		Message: "opencode.json found and valid",
	}
}

// checkUnityAgents verifies Unity agents are installed.
func checkUnityAgents() CheckResult {
	configPath := configpkg.DefaultOpenCodePath()

	data, err := os.ReadFile(configPath)
	if err != nil {
		return CheckResult{
			Name:    "Unity Agents",
			Status:  "WARN",
			Message: "cannot read opencode.json",
		}
	}

	var config map[string]interface{}
	if err := json.Unmarshal(data, &config); err != nil {
		return CheckResult{
			Name:    "Unity Agents",
			Status:  "WARN",
			Message: "cannot parse opencode.json",
		}
	}

	agentMap, ok := config["agent"].(map[string]interface{})
	if !ok {
		return CheckResult{
			Name:    "Unity Agents",
			Status:  "WARN",
			Message: "no agents found in opencode.json",
		}
	}

	var unityAgents []string
	for name := range agentMap {
		if strings.HasPrefix(name, "unity-") {
			unityAgents = append(unityAgents, name)
		}
	}

	if len(unityAgents) == 0 {
		return CheckResult{
			Name:    "Unity Agents",
			Status:  "WARN",
			Message: "no Unity agents installed",
		}
	}

	return CheckResult{
		Name:    "Unity Agents",
		Status:  "PASS",
		Message: fmt.Sprintf("%d Unity agents installed: %s", len(unityAgents), strings.Join(unityAgents, ", ")),
	}
}

// checkPrompts verifies prompts are present in prompts/unity/.
func checkPrompts() CheckResult {
	home, err := os.UserHomeDir()
	if err != nil {
		return CheckResult{
			Name:    "Prompts",
			Status:  "WARN",
			Message: "cannot determine home directory",
		}
	}

	promptsDir := filepath.Join(home, ".config", "opencode", "prompts", "unity")

	entries, err := os.ReadDir(promptsDir)
	if err != nil {
		if os.IsNotExist(err) {
			return CheckResult{
				Name:    "Prompts",
				Status:  "WARN",
				Message: "prompts/unity/ directory not found",
			}
		}
		return CheckResult{
			Name:    "Prompts",
			Status:  "WARN",
			Message: "cannot read prompts directory",
		}
	}

	// Count .md files
	mdCount := 0
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".md") {
			mdCount++
		}
	}

	if mdCount == 0 {
		return CheckResult{
			Name:    "Prompts",
			Status:  "WARN",
			Message: "no prompt files found in prompts/unity/",
		}
	}

	// Compare with embedded prompts
	embeddedPrompts, err := embed.ListPrompts()
	if err != nil {
		return CheckResult{
			Name:    "Prompts",
			Status:  "WARN",
			Message: fmt.Sprintf("%d prompts found (could not verify against embedded)", mdCount),
		}
	}

	embeddedCount := len(embeddedPrompts)
	if mdCount < embeddedCount {
		return CheckResult{
			Name:    "Prompts",
			Status:  "WARN",
			Message: fmt.Sprintf("only %d/%d prompt files present", mdCount, embeddedCount),
		}
	}

	return CheckResult{
		Name:    "Prompts",
		Status:  "PASS",
		Message: fmt.Sprintf("%d/%d prompt files present", mdCount, embeddedCount),
	}
}

// checkSkills verifies skills are present in skills/unity-6000/.
func checkSkills() CheckResult {
	home, err := os.UserHomeDir()
	if err != nil {
		return CheckResult{
			Name:    "Skills",
			Status:  "WARN",
			Message: "cannot determine home directory",
		}
	}

	skillsDir := filepath.Join(home, ".config", "opencode", "skills", "unity-6000")

	entries, err := os.ReadDir(skillsDir)
	if err != nil {
		if os.IsNotExist(err) {
			return CheckResult{
				Name:    "Skills",
				Status:  "WARN",
				Message: "skills/unity-6000/ directory not found",
			}
		}
		return CheckResult{
			Name:    "Skills",
			Status:  "WARN",
			Message: "cannot read skills directory",
		}
	}

	// Count directories
	skillCount := 0
	for _, entry := range entries {
		if entry.IsDir() {
			skillCount++
		}
	}

	if skillCount == 0 {
		return CheckResult{
			Name:    "Skills",
			Status:  "WARN",
			Message: "no skill directories found in skills/unity-6000/",
		}
	}

	// Compare with embedded skills
	embeddedSkills, err := embed.ListSkills()
	if err != nil {
		return CheckResult{
			Name:    "Skills",
			Status:  "WARN",
			Message: fmt.Sprintf("%d skills found (could not verify against embedded)", skillCount),
		}
	}

	embeddedCount := len(embeddedSkills)
	if skillCount < embeddedCount {
		return CheckResult{
			Name:    "Skills",
			Status:  "WARN",
			Message: fmt.Sprintf("only %d/%d skills present", skillCount, embeddedCount),
		}
	}

	return CheckResult{
		Name:    "Skills",
		Status:  "PASS",
		Message: fmt.Sprintf("%d/%d skills present", skillCount, embeddedCount),
	}
}

// checkStateFile verifies state.json exists.
func checkStateFile() CheckResult {
	statePath := configpkg.DefaultStatePath()

	if _, err := os.Stat(statePath); os.IsNotExist(err) {
		return CheckResult{
			Name:    "State File",
			Status:  "WARN",
			Message: "state.json not found (run install first)",
		}
	}

	// Verify it's valid JSON and valid state
	data, err := os.ReadFile(statePath)
	if err != nil {
		return CheckResult{
			Name:    "State File",
			Status:  "WARN",
			Message: "cannot read state.json",
		}
	}

	var jsonRaw json.RawMessage
	if err := json.Unmarshal(data, &jsonRaw); err != nil {
		return CheckResult{
			Name:    "State File",
			Status:  "FAIL",
			Message: "state.json is not valid JSON",
		}
	}

	var state configpkg.State
	if err := json.Unmarshal(data, &state); err != nil {
		return CheckResult{
			Name:    "State File",
			Status:  "FAIL",
			Message: "state.json has invalid structure",
		}
	}

	if state.Version == 0 {
		return CheckResult{
			Name:    "State File",
			Status:  "FAIL",
			Message: "state.json has invalid version",
		}
	}

	return CheckResult{
		Name:    "State File",
		Status:  "PASS",
		Message: fmt.Sprintf("state.json valid (v%d, %d agents)", state.Version, len(state.InstalledAgents)),
	}
}
