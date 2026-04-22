package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// RecoveryFromOpencode scans opencode.json to rebuild the state.
func RecoveryFromOpencode(configPath string) (*State, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read opencode.json: %w", err)
	}

	var config map[string]interface{}
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse opencode.json: %w", err)
	}

	// Find all unity agents
	var installedAgents []string
	provider := "opencode" // default

	agentMap, ok := config["agent"].(map[string]interface{})
	if ok {
		for name := range agentMap {
			if strings.HasPrefix(name, "unity-") {
				installedAgents = append(installedAgents, name)
			}
		}
	}

	state := &State{
		Version:          1,
		InstalledAt:      "", // will be filled by caller if needed
		UpdatedAt:        "", // will be filled by caller if needed
		InstalledAgents:  installedAgents,
		Provider:         provider,
		ModelAssignments: make(map[string]string),
	}

	return state, nil
}

// ValidState checks if state.json exists and has valid structure.
func ValidState(path string) (bool, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	var state State
	if err := json.Unmarshal(data, &state); err != nil {
		return false, fmt.Errorf("state.json parse error: %w", err)
	}

	// Basic validation
	if state.Version == 0 {
		return false, fmt.Errorf("state.json has invalid version")
	}

	if state.InstalledAgents == nil {
		return false, fmt.Errorf("state.json missing InstalledAgents")
	}

	return true, nil
}

// StateExists checks if state.json exists.
func StateExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// DeleteState removes the state.json file.
func DeleteState(path string) error {
	if err := os.Remove(path); err != nil {
		if os.IsNotExist(err) {
			return nil // already gone
		}
		return err
	}
	return nil
}
