package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestReadState_WriteState_Roundtrip(t *testing.T) {
	tmpDir := t.TempDir()
	statePath := filepath.Join(tmpDir, "state.json")

	// Create a state
	original := &State{
		Version:          1,
		InstalledAt:      "2024-01-15T10:30:00Z",
		UpdatedAt:        "2024-01-15T10:30:00Z",
		InstalledAgents:  []string{"unity-6000-expert", "unity-graphics-expert"},
		Provider:         "claude",
		ModelAssignments: map[string]string{
			"unity-6000-expert": "anthropic/claude-sonnet-4-20250514",
		},
	}

	// Write state
	if err := WriteState(statePath, original); err != nil {
		t.Fatalf("WriteState failed: %v", err)
	}

	// Read state
	read, err := ReadState(statePath)
	if err != nil {
		t.Fatalf("ReadState failed: %v", err)
	}

	// Verify fields
	if read.Version != original.Version {
		t.Errorf("Version mismatch: got %d, want %d", read.Version, original.Version)
	}
	if read.InstalledAt != original.InstalledAt {
		t.Errorf("InstalledAt mismatch: got %s, want %s", read.InstalledAt, original.InstalledAt)
	}
	if read.Provider != original.Provider {
		t.Errorf("Provider mismatch: got %s, want %s", read.Provider, original.Provider)
	}
	if len(read.InstalledAgents) != len(original.InstalledAgents) {
		t.Errorf("InstalledAgents length mismatch: got %d, want %d", len(read.InstalledAgents), len(original.InstalledAgents))
	}
}

func TestValidState_ValidState(t *testing.T) {
	tmpDir := t.TempDir()
	statePath := filepath.Join(tmpDir, "state.json")

	// Create a valid state
	validState := &State{
		Version:          1,
		InstalledAt:      "2024-01-15T10:30:00Z",
		UpdatedAt:        "2024-01-15T10:30:00Z",
		InstalledAgents:  []string{"unity-6000-expert"},
		Provider:         "claude",
		ModelAssignments: map[string]string{},
	}

	if err := WriteState(statePath, validState); err != nil {
		t.Fatalf("WriteState failed: %v", err)
	}

	valid, err := ValidState(statePath)
	if err != nil {
		t.Fatalf("ValidState returned error: %v", err)
	}
	if !valid {
		t.Error("Expected ValidState to return true for valid state")
	}
}

func TestValidState_NonExistentFile(t *testing.T) {
	valid, err := ValidState("nonExistentPath12345.json")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if valid {
		t.Error("Expected ValidState to return false for non-existent file")
	}
}

func TestValidState_InvalidJSON(t *testing.T) {
	tmpDir := t.TempDir()
	statePath := filepath.Join(tmpDir, "state.json")

	// Write invalid JSON
	if err := os.WriteFile(statePath, []byte("not valid json{"), 0644); err != nil {
		t.Fatalf("Failed to write file: %v", err)
	}

	valid, err := ValidState(statePath)
	if err == nil {
		t.Error("Expected error for invalid JSON")
	}
	if valid {
		t.Error("Expected ValidState to return false for invalid JSON")
	}
}

func TestValidState_InvalidVersion(t *testing.T) {
	tmpDir := t.TempDir()
	statePath := filepath.Join(tmpDir, "state.json")

	// Write state with version 0
	invalidState := &State{
		Version:          0, // Invalid
		InstalledAt:      "2024-01-15T10:30:00Z",
		UpdatedAt:        "2024-01-15T10:30:00Z",
		InstalledAgents:  []string{},
		Provider:         "claude",
		ModelAssignments: map[string]string{},
	}

	data, _ := json.Marshal(invalidState)
	os.WriteFile(statePath, data, 0644)

	valid, err := ValidState(statePath)
	if err == nil {
		t.Error("Expected error for invalid version")
	}
	if valid {
		t.Error("Expected ValidState to return false for version 0")
	}
}

func TestValidState_NilInstalledAgents(t *testing.T) {
	tmpDir := t.TempDir()
	statePath := filepath.Join(tmpDir, "state.json")

	// Write state with nil InstalledAgents
	data, _ := json.Marshal(map[string]interface{}{
		"version":          1,
		"installedAt":      "2024-01-15T10:30:00Z",
		"updatedAt":        "2024-01-15T10:30:00Z",
		"installedAgents":  nil, // Invalid
		"provider":         "claude",
		"modelAssignments": map[string]string{},
	})
	os.WriteFile(statePath, data, 0644)

	valid, err := ValidState(statePath)
	if err == nil {
		t.Error("Expected error for nil InstalledAgents")
	}
	if valid {
		t.Error("Expected ValidState to return false for nil InstalledAgents")
	}
}

func TestRecoveryFromOpencode_Basic(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "opencode.json")

	// Create a config with some unity agents
	config := map[string]interface{}{
		"agent": map[string]interface{}{
			"unity-6000-expert":     map[string]interface{}{"description": "Orchestrator"},
			"unity-graphics-expert": map[string]interface{}{"description": "Graphics"},
			"unity-physics-expert":  map[string]interface{}{"description": "Physics"},
			"other-agent":           map[string]interface{}{"description": "Not Unity"},
		},
	}

	data, _ := json.Marshal(config)
	if err := os.WriteFile(configPath, data, 0644); err != nil {
		t.Fatalf("Failed to write config: %v", err)
	}

	state, err := RecoveryFromOpencode(configPath)
	if err != nil {
		t.Fatalf("RecoveryFromOpencode failed: %v", err)
	}

	// Should have found 3 unity agents
	if len(state.InstalledAgents) != 3 {
		t.Errorf("Expected 3 installed agents, got %d", len(state.InstalledAgents))
	}

	// Should default to opencode provider
	if state.Provider != "opencode" {
		t.Errorf("Expected provider 'opencode', got %s", state.Provider)
	}

	// Verify specific agents
	expectedAgents := map[string]bool{
		"unity-6000-expert":     true,
		"unity-graphics-expert": true,
		"unity-physics-expert":  true,
	}
	for _, agent := range state.InstalledAgents {
		if !expectedAgents[agent] {
			t.Errorf("Unexpected agent found: %s", agent)
		}
	}
}

func TestRecoveryFromOpencode_NoAgents(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "opencode.json")

	config := map[string]interface{}{
		"agent": map[string]interface{}{
			"other-agent": map[string]interface{}{"description": "Not Unity"},
		},
	}

	data, _ := json.Marshal(config)
	if err := os.WriteFile(configPath, data, 0644); err != nil {
		t.Fatalf("Failed to write config: %v", err)
	}

	state, err := RecoveryFromOpencode(configPath)
	if err != nil {
		t.Fatalf("RecoveryFromOpencode failed: %v", err)
	}

	if len(state.InstalledAgents) != 0 {
		t.Errorf("Expected 0 agents, got %d", len(state.InstalledAgents))
	}
}

func TestRecoveryFromOpencode_NoAgentSection(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "opencode.json")

	config := map[string]interface{}{
		"version": "1.0",
	}

	data, _ := json.Marshal(config)
	if err := os.WriteFile(configPath, data, 0644); err != nil {
		t.Fatalf("Failed to write config: %v", err)
	}

	state, err := RecoveryFromOpencode(configPath)
	if err != nil {
		t.Fatalf("RecoveryFromOpencode failed: %v", err)
	}

	if len(state.InstalledAgents) != 0 {
		t.Errorf("Expected 0 agents for config without agent section, got %d", len(state.InstalledAgents))
	}
}

func TestStateExists(t *testing.T) {
	tmpDir := t.TempDir()
	statePath := filepath.Join(tmpDir, "state.json")

	// Should not exist initially
	if StateExists(statePath) {
		t.Error("Expected StateExists to return false for non-existent file")
	}

	// Create file
	os.WriteFile(statePath, []byte("{}"), 0644)

	// Should exist now
	if !StateExists(statePath) {
		t.Error("Expected StateExists to return true for existing file")
	}
}

func TestDeleteState(t *testing.T) {
	tmpDir := t.TempDir()
	statePath := filepath.Join(tmpDir, "state.json")

	// Create file
	os.WriteFile(statePath, []byte("{}"), 0644)

	// Delete should succeed
	if err := DeleteState(statePath); err != nil {
		t.Fatalf("DeleteState failed: %v", err)
	}

	// File should not exist
	if StateExists(statePath) {
		t.Error("Expected file to be deleted")
	}
}

func TestDeleteState_NonExistent(t *testing.T) {
	// Deleting non-existent file should not error
	err := DeleteState("nonExistent12345.json")
	if err != nil {
		t.Fatalf("DeleteState should not error for non-existent file: %v", err)
	}
}