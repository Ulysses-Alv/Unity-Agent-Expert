package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	configpkg "github.com/Ulysses-Alv/unity-expert/internal/config"
	"github.com/Ulysses-Alv/unity-expert/internal/embed"
)

// DriftResult holds the result of a drift detection.
type DriftResult struct {
	HasDrift      bool
	MissingAgents []string // agents in state but not in opencode.json
	ExtraAgents   []string // agents in opencode.json but not in state
	ModelChanges  map[string]ModelChange
}

// ModelChange represents a model assignment change.
type ModelChange struct {
	Agent    string
	OldModel string
	NewModel string
}

// RunSync performs an idempotent sync of the Unity configuration.
func RunSync(args []string) error {
	opts, err := ParseSyncFlags(args)
	if err != nil {
		return err
	}

	fmt.Printf("[INFO] Sync command called with provider=%s, dryRun=%v, force=%v\n",
		opts.Provider, opts.DryRun, opts.Force)

	if opts.DryRun {
		fmt.Println("[DRY RUN] Would sync Unity configuration")
		return nil
	}

	// Read state.json
	state, err := configpkg.ReadState(configpkg.DefaultStatePath())
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("[WARN] No state.json found. Run 'unity-agent-expert install' first.")
			return fmt.Errorf("state.json not found: please run install first")
		}
		return fmt.Errorf("failed to read state.json: %w", err)
	}

	// Read opencode.json
	configPath := configpkg.DefaultOpenCodePath()
	configData, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("failed to read opencode.json: %w", err)
	}

	var opencodeConfig map[string]interface{}
	if err := json.Unmarshal(configData, &opencodeConfig); err != nil {
		return fmt.Errorf("failed to parse opencode.json: %w", err)
	}

	// Detect drift
	drift, err := detectDrift(state, opencodeConfig)
	if err != nil {
		return fmt.Errorf("drift detection failed: %w", err)
	}

	if drift.HasDrift && !opts.Force {
		fmt.Println("[WARN] Drift detected between state.json and opencode.json:")
		if len(drift.MissingAgents) > 0 {
			fmt.Printf("  Missing agents: %s\n", strings.Join(drift.MissingAgents, ", "))
		}
		if len(drift.ExtraAgents) > 0 {
			fmt.Printf("  Extra agents: %s\n", strings.Join(drift.ExtraAgents, ", "))
		}
		for agent, change := range drift.ModelChanges {
			fmt.Printf("  Model changed for %s: %s -> %s\n", agent, change.OldModel, change.NewModel)
		}
		fmt.Println("[INFO] Use --force to ignore drift and re-apply anyway")
		return fmt.Errorf("drift detected: run with --force to override")
	}

	if drift.HasDrift {
		fmt.Println("[INFO] Drift detected but --force specified, re-applying configuration...")
	}

	// Determine provider to use
	provider := opts.Provider
	if provider == "" {
		provider = state.Provider
	}
	if provider == "" {
		provider = "opencode"
	}

	// Read embedded presets.json
	presetsData, err := embed.EmbeddedPresetsJSON()
	if err != nil {
		return fmt.Errorf("failed to read presets.json: %w", err)
	}

	var presets map[string]struct {
		Models map[string]string `json:"models"`
	}
	if err := json.Unmarshal(presetsData, &presets); err != nil {
		return fmt.Errorf("failed to parse presets.json: %w", err)
	}

	// Get models for the selected provider
	providerModels := make(map[string]string)
	if providerPreset, ok := presets[provider]; ok {
		providerModels = providerPreset.Models
	}

	// Update model assignments in opencode.json (update only, don't re-add agents)
	if err := syncModelAssignments(opencodeConfig, state.InstalledAgents, providerModels); err != nil {
		return fmt.Errorf("failed to sync model assignments: %w", err)
	}

	// Write updated opencode.json
	if err := configpkg.WriteOpenCodeConfig(configPath, opencodeConfig); err != nil {
		return fmt.Errorf("failed to write opencode.json: %w", err)
	}

	// Update state.json with new provider and model assignments
	state.Provider = provider
	state.ModelAssignments = make(map[string]string)
	for _, agent := range state.InstalledAgents {
		if model, ok := providerModels[agent]; ok {
			state.ModelAssignments[agent] = model
		}
	}

	if err := configpkg.WriteState(configpkg.DefaultStatePath(), state); err != nil {
		return fmt.Errorf("failed to write state.json: %w", err)
	}

	fmt.Println("[OK] Sync completed successfully")
	return nil
}

// detectDrift compares state.json against opencode.json to find differences.
func detectDrift(state *configpkg.State, opencodeConfig map[string]interface{}) (*DriftResult, error) {
	result := &DriftResult{
		ModelChanges: make(map[string]ModelChange),
	}

	// Build set of current unity agents from opencode.json
	currentAgents := make(map[string]string) // agent name -> model
	if agentMap, ok := opencodeConfig["agent"].(map[string]interface{}); ok {
		for name, data := range agentMap {
			if strings.HasPrefix(name, "unity-") {
				if agentData, ok := data.(map[string]interface{}); ok {
					if model, ok := agentData["model"].(string); ok {
						currentAgents[name] = model
					} else {
						currentAgents[name] = ""
					}
				}
			}
		}
	}

	// Check for missing agents (in state but not in opencode.json)
	for _, agent := range state.InstalledAgents {
		if _, exists := currentAgents[agent]; !exists {
			result.MissingAgents = append(result.MissingAgents, agent)
			result.HasDrift = true
		}
	}

	// Check for extra agents (in opencode.json but not in state)
	for agent := range currentAgents {
		found := false
		for _, stateAgent := range state.InstalledAgents {
			if agent == stateAgent {
				found = true
				break
			}
		}
		if !found {
			result.ExtraAgents = append(result.ExtraAgents, agent)
			result.HasDrift = true
		}
	}

	// Check for model assignment changes
	for agent, currentModel := range currentAgents {
		if storedModel, exists := state.ModelAssignments[agent]; exists {
			if storedModel != currentModel {
				result.ModelChanges[agent] = ModelChange{
					Agent:    agent,
					OldModel: storedModel,
					NewModel: currentModel,
				}
				result.HasDrift = true
			}
		}
	}

	return result, nil
}

// syncModelAssignments updates model assignments for existing Unity agents.
func syncModelAssignments(opencodeConfig map[string]interface{}, installedAgents []string, providerModels map[string]string) error {
	agentMap, ok := opencodeConfig["agent"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("opencode.json missing agent map")
	}

	// Create set of installed agents for quick lookup
	installedSet := make(map[string]bool)
	for _, agent := range installedAgents {
		installedSet[agent] = true
	}

	// Update model assignments only for installed agents
	for agentName := range agentMap {
		if !strings.HasPrefix(agentName, "unity-") {
			continue
		}

		// Only update if this agent was installed by us
		if !installedSet[agentName] {
			continue
		}

		// Apply provider model if available
		if model, ok := providerModels[agentName]; ok {
			agentData, ok := agentMap[agentName].(map[string]interface{})
			if !ok {
				continue
			}
			agentData["model"] = model
		}
	}

	return nil
}
