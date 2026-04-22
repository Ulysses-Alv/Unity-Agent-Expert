package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestMergeAgentsIntoOpenCodeConfig(t *testing.T) {
	// Create a temporary directory for test files
	tmpDir := t.TempDir()
	opencodePath := filepath.Join(tmpDir, "opencode.json")
	agentsPath := filepath.Join(tmpDir, "agents.json")

	// Existing opencode.json with some agents
	existingConfig := map[string]interface{}{
		"agent": map[string]interface{}{
			"existing-agent": map[string]interface{}{
				"description": "Existing agent",
				"model":       "opencode-go/legacy-model",
			},
			"unity-existing": map[string]interface{}{
				"description": "Already installed unity agent",
				"model":       "opencode-go/old-model",
			},
		},
		"version": "1.0",
	}

	existingData, err := json.Marshal(existingConfig)
	if err != nil {
		t.Fatalf("Failed to marshal existing config: %v", err)
	}
	if err := os.WriteFile(opencodePath, existingData, 0644); err != nil {
		t.Fatalf("Failed to write opencode.json: %v", err)
	}

	// agents.json to merge
	agentsToMerge := map[string]interface{}{
		"unity-new-agent": map[string]interface{}{
			"description": "New Unity agent",
			"model":       "opencode-go/new-model",
		},
		"unity-another": map[string]interface{}{
			"description": "Another Unity agent",
			"model":       "opencode-go/another-model",
		},
	}

	agentsData, err := json.Marshal(agentsToMerge)
	if err != nil {
		t.Fatalf("Failed to marshal agents: %v", err)
	}
	if err := os.WriteFile(agentsPath, agentsData, 0644); err != nil {
		t.Fatalf("Failed to write agents.json: %v", err)
	}

	// Read and parse
	configData, err := os.ReadFile(opencodePath)
	if err != nil {
		t.Fatalf("Failed to read opencode.json: %v", err)
	}

	var opencodeConfig map[string]interface{}
	if err := json.Unmarshal(configData, &opencodeConfig); err != nil {
		t.Fatalf("Failed to unmarshal opencode.json: %v", err)
	}

	var embeddedAgents map[string]interface{}
	if err := json.Unmarshal(agentsData, &embeddedAgents); err != nil {
		t.Fatalf("Failed to unmarshal agents.json: %v", err)
	}

	// Ensure agent map exists
	if opencodeConfig["agent"] == nil {
		opencodeConfig["agent"] = make(map[string]interface{})
	}
	agentMap := opencodeConfig["agent"].(map[string]interface{})

	// Merge agents
	for agentName, agentData := range embeddedAgents {
		agentMap[agentName] = agentData
	}

	// Verify existing agents are preserved
	if _, exists := agentMap["existing-agent"]; !exists {
		t.Error("Expected existing-agent to be preserved after merge")
	}
	if _, exists := agentMap["unity-existing"]; !exists {
		t.Error("Expected unity-existing to be preserved after merge")
	}

	// Verify new agents are added
	if _, exists := agentMap["unity-new-agent"]; !exists {
		t.Error("Expected unity-new-agent to be added")
	}
	if _, exists := agentMap["unity-another"]; !exists {
		t.Error("Expected unity-another to be added")
	}

	// Verify version field is preserved (unknown field)
	if version, ok := opencodeConfig["version"].(string); !ok || version != "1.0" {
		t.Error("Expected version field to be preserved")
	}
}

func TestProviderPresetApplication(t *testing.T) {
	// Test that provider presets correctly override agent models
	presets := map[string]struct {
		Models map[string]string `json:"models"`
	}{
		"claude": {
			Models: map[string]string{
				"unity-6000-expert": "anthropic/claude-sonnet-4-20250514",
				"unity-graphics-expert": "anthropic/claude-sonnet-4-20250514",
			},
		},
		"gpt": {
			Models: map[string]string{
				"unity-6000-expert": "openai/gpt-4o",
				"unity-graphics-expert": "openai/gpt-4o-mini",
			},
		},
	}

	// Embedded agents
	embeddedAgents := map[string]interface{}{
		"unity-6000-expert": map[string]interface{}{
			"description": "Orchestrator",
			"model":       "opencode-go/minimax-m2.7",
		},
		"unity-graphics-expert": map[string]interface{}{
			"description": "Graphics expert",
			"model":       "opencode-go/mimo-v2-pro",
		},
	}

	// Apply clau   de preset
	provider := "claude"
	providerModels := make(map[string]string)
	if providerPreset, ok := presets[provider]; ok {
		providerModels = providerPreset.Models
	}

	for agentName, agentData := range embeddedAgents {
		if model, ok := providerModels[agentName]; ok {
			if agentDataMap, ok := agentData.(map[string]interface{}); ok {
				agentDataMap["model"] = model
			}
		}
	}

	// Verify models were overridden
	unity6000 := embeddedAgents["unity-6000-expert"].(map[string]interface{})
	if model, ok := unity6000["model"].(string); !ok || model != "anthropic/claude-sonnet-4-20250514" {
		t.Errorf("Expected claude model for unity-6000-expert, got: %v", unity6000["model"])
	}

	graphics := embeddedAgents["unity-graphics-expert"].(map[string]interface{})
	if model, ok := graphics["model"].(string); !ok || model != "anthropic/claude-sonnet-4-20250514" {
		t.Errorf("Expected claude model for unity-graphics-expert, got: %v", graphics["model"])
	}
}

func TestPreservingUnknownFields(t *testing.T) {
	// Test that when we marshal/unmarshal, unknown fields are preserved
	// This simulates how Raw field works in AgentConfig

	originalJSON := `{
		"agent": {
			"unity-test": {
				"description": "Test agent",
				"model": "opencode-go/test",
				"customField": "should be preserved",
				"nested": {
					"key": "value"
				}
			}
		},
		"unknownTopLevel": "should also be preserved"
	}`

	var result map[string]interface{}
	if err := json.Unmarshal([]byte(originalJSON), &result); err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	// Re-marshal and unmarshal again to simulate round-trip
	output, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}

	var roundTripped map[string]interface{}
	if err := json.Unmarshal(output, &roundTripped); err != nil {
		t.Fatalf("Failed to unmarshal round-tripped: %v", err)
	}

	// Check unknown top-level field is preserved
	if _, ok := roundTripped["unknownTopLevel"]; !ok {
		t.Error("Expected unknownTopLevel to be preserved")
	}

	// Check nested custom field
	agentMap := roundTripped["agent"].(map[string]interface{})
	unityTest := agentMap["unity-test"].(map[string]interface{})
	if _, ok := unityTest["customField"]; !ok {
		t.Error("Expected customField to be preserved in agent")
	}
	if nested, ok := unityTest["nested"].(map[string]interface{}); !ok || nested["key"] != "value" {
		t.Error("Expected nested object to be preserved")
	}
}

func TestUTF8NoBOMOutput(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "test.json")

	config := map[string]interface{}{
		"agent": map[string]interface{}{
			"test-agent": map[string]interface{}{
				"description": "Test with unicode: 日本語, émojis: 🎮, special: ñ",
			},
		},
	}

	// Write using our WriteOpenCodeConfig function
	if err := WriteOpenCodeConfig(configPath, config); err != nil {
		t.Fatalf("WriteOpenCodeConfig failed: %v", err)
	}

	// Read back and verify UTF-8 without BOM
	data, err := os.ReadFile(configPath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	// Check for BOM (UTF-8 BOM is 0xEF 0xBB 0xBF)
	if len(data) >= 3 && data[0] == 0xEF && data[1] == 0xBB && data[2] == 0xBF {
		t.Error("File should not have UTF-8 BOM")
	}

	// Verify content is valid JSON
	var parsed map[string]interface{}
	if err := json.Unmarshal(data, &parsed); err != nil {
		t.Errorf("Output should be valid JSON: %v", err)
	}

	// Verify Unicode is preserved
	agentMap := parsed["agent"].(map[string]interface{})
	testAgent := agentMap["test-agent"].(map[string]interface{})
	desc := testAgent["description"].(string)
	if desc == "" {
		t.Error("Description should not be empty")
	}
}

func TestMergePreservesExistingAgents(t *testing.T) {
	// Create opencode config with existing agents
	existingConfig := map[string]interface{}{
		"agent": map[string]interface{}{
			"pre-existing": map[string]interface{}{
				"description": "Pre-existing agent",
				"model":       "some-model",
			},
			"unity-pre-existing": map[string]interface{}{
				"description": "Unity pre-existing",
				"model":       "unity-model-v1",
			},
		},
	}

	// New agents to merge
	newAgents := map[string]interface{}{
		"unity-new": map[string]interface{}{
			"description": "New unity agent",
			"model":       "new-model",
		},
	}

	// Ensure agent map exists
	if existingConfig["agent"] == nil {
		existingConfig["agent"] = make(map[string]interface{})
	}
	agentMap := existingConfig["agent"].(map[string]interface{})

	// Merge - should not overwrite existing
	for agentName, agentData := range newAgents {
		if _, exists := agentMap[agentName]; !exists {
			agentMap[agentName] = agentData
		}
	}

	// Verify pre-existing is still there
	if _, exists := agentMap["pre-existing"]; !exists {
		t.Error("pre-existing agent should be preserved")
	}
	if _, exists := agentMap["unity-pre-existing"]; !exists {
		t.Error("unity-pre-existing agent should be preserved")
	}

	// Verify new agent is added
	if _, exists := agentMap["unity-new"]; !exists {
		t.Error("unity-new agent should be added")
	}

	// Verify original models are NOT overwritten (test the logic of not overwriting)
	existingAgent := agentMap["pre-existing"].(map[string]interface{})
	if model, ok := existingAgent["model"].(string); !ok || model != "some-model" {
		t.Error("pre-existing agent model should not be changed")
	}
}

func TestMergeAgents_PreservesNonUnityAgents(t *testing.T) {
	existing := map[string]interface{}{
		"agent": map[string]interface{}{
			"custom-agent": map[string]interface{}{
				"description": "My custom agent",
				"model":       "my-custom-model",
			},
		},
	}

	embedded := map[string]interface{}{
		"custom-agent": map[string]interface{}{
			"description": "Updated description",
			"model":       "new-model",
		},
	}

	providerModels := map[string]string{
		"custom-agent": "preset-model",
	}

	result := MergeAgents(existing, embedded, providerModels, false, nil)
	agentMap := result["agent"].(map[string]interface{})

	customAgent := agentMap["custom-agent"].(map[string]interface{})
	if customAgent["model"] != "my-custom-model" {
		t.Error("Non-unity agent model should be preserved, not overwritten")
	}
	if customAgent["description"] != "My custom agent" {
		t.Error("Non-unity agent description should be preserved")
	}
}

func TestMergeAgents_PreservesUnityAgentWithModel(t *testing.T) {
	existing := map[string]interface{}{
		"agent": map[string]interface{}{
			"unity-6000-expert": map[string]interface{}{
				"description": "My unity agent",
				"model":       "my-custom-unity-model",
			},
		},
	}

	embedded := map[string]interface{}{
		"unity-6000-expert": map[string]interface{}{
			"description": "Updated unity agent",
			"model":       "preset-unity-model",
		},
	}

	providerModels := map[string]string{
		"unity-6000-expert": "preset-unity-model",
	}

	result := MergeAgents(existing, embedded, providerModels, false, nil)
	agentMap := result["agent"].(map[string]interface{})

	unityAgent := agentMap["unity-6000-expert"].(map[string]interface{})
	if unityAgent["model"] != "my-custom-unity-model" {
		t.Error("Unity agent with existing model should be preserved")
	}
}

func TestMergeAgents_PreservesUnityAgentWithModel_ForceOverwrites(t *testing.T) {
	existing := map[string]interface{}{
		"agent": map[string]interface{}{
			"unity-6000-expert": map[string]interface{}{
				"description": "My unity agent",
				"model":       "my-custom-unity-model",
			},
		},
	}

	embedded := map[string]interface{}{
		"unity-6000-expert": map[string]interface{}{
			"description": "Updated unity agent",
			"model":       "preset-unity-model",
		},
	}

	providerModels := map[string]string{
		"unity-6000-expert": "preset-unity-model",
	}

	result := MergeAgents(existing, embedded, providerModels, true, nil)
	agentMap := result["agent"].(map[string]interface{})

	unityAgent := agentMap["unity-6000-expert"].(map[string]interface{})
	if unityAgent["model"] != "preset-unity-model" {
		t.Error("With force=true, unity agent should be overwritten with preset model")
	}
}

func TestMergeAgents_UnityAgentWithoutModel_GetsPresetModel(t *testing.T) {
	existing := map[string]interface{}{
		"agent": map[string]interface{}{
			"unity-6000-expert": map[string]interface{}{
				"description": "Unity agent without model",
				// No model set
			},
		},
	}

	embedded := map[string]interface{}{
		"unity-6000-expert": map[string]interface{}{
			"description": "Unity agent from embedded",
			"model":       "default-model",
		},
	}

	providerModels := map[string]string{
		"unity-6000-expert": "preset-claude-model",
	}

	result := MergeAgents(existing, embedded, providerModels, false, nil)
	agentMap := result["agent"].(map[string]interface{})

	unityAgent := agentMap["unity-6000-expert"].(map[string]interface{})
	if unityAgent["model"] != "preset-claude-model" {
		t.Errorf("Unity agent without model should get preset model, got: %v", unityAgent["model"])
	}
}

func TestMergeAgents_NewUnityAgents_AreAdded(t *testing.T) {
	existing := map[string]interface{}{
		"agent": map[string]interface{}{},
	}

	embedded := map[string]interface{}{
		"unity-new-agent": map[string]interface{}{
			"description": "Brand new agent",
			"model":       "new-model",
		},
	}

	providerModels := map[string]string{
		"unity-new-agent": "preset-model",
	}

	result := MergeAgents(existing, embedded, providerModels, false, nil)
	agentMap := result["agent"].(map[string]interface{})

	if _, exists := agentMap["unity-new-agent"]; !exists {
		t.Error("New unity agent should be added")
	}

	unityAgent := agentMap["unity-new-agent"].(map[string]interface{})
	if unityAgent["model"] != "preset-model" {
		t.Errorf("New agent should get preset model, got: %v", unityAgent["model"])
	}
}

func TestMergeAgents_AgentFilter(t *testing.T) {
	existing := map[string]interface{}{
		"agent": map[string]interface{}{},
	}

	embedded := map[string]interface{}{
		"unity-agent-a": map[string]interface{}{"description": "Agent A"},
		"unity-agent-b": map[string]interface{}{"description": "Agent B"},
		"unity-agent-c": map[string]interface{}{"description": "Agent C"},
	}

	providerModels := map[string]string{}

	result := MergeAgents(existing, embedded, providerModels, false, []string{"unity-agent-a", "unity-agent-b"})
	agentMap := result["agent"].(map[string]interface{})

	if _, exists := agentMap["unity-agent-a"]; !exists {
		t.Error("unity-agent-a should be added (in filter)")
	}
	if _, exists := agentMap["unity-agent-b"]; !exists {
		t.Error("unity-agent-b should be added (in filter)")
	}
	if _, exists := agentMap["unity-agent-c"]; exists {
		t.Error("unity-agent-c should NOT be added (not in filter)")
	}
}