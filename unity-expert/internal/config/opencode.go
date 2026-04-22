package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

// DefaultOpenCodePath returns the default path to opencode.json.
func DefaultOpenCodePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return "~/.config/opencode/opencode.json"
	}
	return filepath.Join(home, ".config", "opencode", "opencode.json")
}

// DefaultUnityExpertPath returns the default base path for unity-expert data.
func DefaultUnityExpertPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return "~/.unity-expert"
	}
	return filepath.Join(home, ".unity-expert")
}

// DefaultBackupsPath returns the default backups directory path.
func DefaultBackupsPath() string {
	return filepath.Join(DefaultUnityExpertPath(), "backups")
}

// DefaultStatePath returns the default state.json path.
func DefaultStatePath() string {
	return filepath.Join(DefaultUnityExpertPath(), "state.json")
}

// ReadOpenCodeConfig reads and parses the opencode.json file.
func ReadOpenCodeConfig(path string) (map[string]interface{}, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// WriteOpenCodeConfig writes the opencode.json file atomically.
func WriteOpenCodeConfig(path string, config map[string]interface{}) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	// Atomic write: write to temp file then rename
	dir := filepath.Dir(path)
	tmp, err := os.CreateTemp(dir, "opencode-*.json")
	if err != nil {
		return err
	}
	tmpPath := tmp.Name()

	if _, err := tmp.Write(data); err != nil {
		tmp.Close()
		os.Remove(tmpPath)
		return err
	}
	if err := tmp.Close(); err != nil {
		os.Remove(tmpPath)
		return err
	}

	if err := os.Rename(tmpPath, path); err != nil {
		os.Remove(tmpPath)
		return err
	}

	return nil
}

// ReadState reads and parses the state.json file.
func ReadState(path string) (*State, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var state State
	if err := json.Unmarshal(data, &state); err != nil {
		return nil, err
	}

	return &state, nil
}

// WriteState writes the state.json file atomically.
func WriteState(path string, state *State) error {
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}

	// Ensure directory exists
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Atomic write: write to temp file then rename
	tmp, err := os.CreateTemp(dir, "state-*.json")
	if err != nil {
		return err
	}
	tmpPath := tmp.Name()

	if _, err := tmp.Write(data); err != nil {
		tmp.Close()
		os.Remove(tmpPath)
		return err
	}
	if err := tmp.Close(); err != nil {
		os.Remove(tmpPath)
		return err
	}

	if err := os.Rename(tmpPath, path); err != nil {
		os.Remove(tmpPath)
		return err
	}

	return nil
}

// EnsureDir ensures a directory exists, creating it if necessary.
func EnsureDir(path string) error {
	return os.MkdirAll(path, 0755)
}

// MergeAgents merges embedded agents into existing opencode config.
// Merge strategy:
//   - Non-unity agents: ALWAYS preserved (never overwritten)
//   - Unity-* agents without model: get preset model (not overwritten)
//   - Unity-* agents with existing model: preserved (unless force=true)
//   - New unity-* agents: always added
//   - force=true: overwrite everything
//   - agents filter: if non-empty, only process listed agents
func MergeAgents(existing map[string]interface{}, embedded map[string]interface{}, providerModels map[string]string, force bool, agents []string) map[string]interface{} {
	if existing["agent"] == nil {
		existing["agent"] = make(map[string]interface{})
	}
	agentMap := existing["agent"].(map[string]interface{})

	for agentName, agentData := range embedded {
		// Filter to specific agents if requested
		if len(agents) > 0 {
			found := false
			for _, a := range agents {
				if a == agentName {
					found = true
					break
				}
			}
			if !found {
				continue
			}
		}

		existingAgent, exists := agentMap[agentName]

		if !exists {
			// New agent — add it with preset model if available
			if model, ok := providerModels[agentName]; ok {
				if agentDataMap, ok := agentData.(map[string]interface{}); ok {
					agentDataMap["model"] = model
				}
			}
			agentMap[agentName] = agentData
			continue
		}

		// Agent exists — check if it should be preserved
		if !strings.HasPrefix(agentName, "unity-") {
			// Non-unity agent — always preserve
			continue
		}

		// Unity agent — check force flag
		if force {
			// Force: overwrite with preset model
			if model, ok := providerModels[agentName]; ok {
				if agentDataMap, ok := agentData.(map[string]interface{}); ok {
					agentDataMap["model"] = model
				}
			}
			agentMap[agentName] = agentData
			continue
		}

		// No force: check if existing agent has a model
		existingAgentMap, ok := existingAgent.(map[string]interface{})
		if !ok {
			// Corrupted entry — overwrite
			if model, ok := providerModels[agentName]; ok {
				if agentDataMap, ok := agentData.(map[string]interface{}); ok {
					agentDataMap["model"] = model
				}
			}
			agentMap[agentName] = agentData
			continue
		}

		if _, hasModel := existingAgentMap["model"]; hasModel {
			// Existing unity agent has model — preserve
			continue
		}

		// Existing unity agent has NO model — apply preset
		if model, ok := providerModels[agentName]; ok {
			existingAgentMap["model"] = model
		}
		// Don't overwrite the agent entry, just ensure model is set
	}

	return existing
}
