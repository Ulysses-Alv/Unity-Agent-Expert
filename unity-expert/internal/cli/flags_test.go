package cli

import (
	"testing"
)

func TestParseInstallFlags_ValidProviders(t *testing.T) {
	testCases := []struct {
		provider string
		wantErr  bool
	}{
		{"opencode", false},
		{"claude", false},
		{"gpt", false},
		{"gemini", false},
		{"custom", false},
	}

	for _, tc := range testCases {
		t.Run(tc.provider, func(t *testing.T) {
			opts, err := ParseInstallFlags([]string{"--provider", tc.provider})
			if tc.wantErr && err == nil {
				t.Errorf("Expected error for provider %q, got nil", tc.provider)
			}
			if !tc.wantErr && err != nil {
				t.Errorf("Unexpected error for provider %q: %v", tc.provider, err)
			}
			if !tc.wantErr && opts.Provider != tc.provider {
				t.Errorf("Expected provider %q, got %q", tc.provider, opts.Provider)
			}
		})
	}
}

func TestParseInstallFlags_InvalidProvider(t *testing.T) {
	invalidProviders := []string{"anthropic", "openai", "google", "invalid", "OPENCODE", "Claude"}

	for _, provider := range invalidProviders {
		t.Run(provider, func(t *testing.T) {
			_, err := ParseInstallFlags([]string{"--provider", provider})
			if err == nil {
				t.Errorf("Expected error for invalid provider %q, got nil", provider)
			}
		})
	}
}

func TestParseInstallFlags_DryRun(t *testing.T) {
	opts, err := ParseInstallFlags([]string{"--dry-run"})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !opts.DryRun {
		t.Error("Expected DryRun to be true")
	}
}

func TestParseInstallFlags_Force(t *testing.T) {
	opts, err := ParseInstallFlags([]string{"--force"})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !opts.Force {
		t.Error("Expected Force to be true")
	}
}

func TestParseInstallFlags_Combined(t *testing.T) {
	opts, err := ParseInstallFlags([]string{
		"--provider", "claude",
		"--dry-run",
		"--force",
		"--agents", "unity-6000-expert,unity-graphics-expert",
	})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if opts.Provider != "claude" {
		t.Errorf("Expected provider 'claude', got %q", opts.Provider)
	}
	if !opts.DryRun {
		t.Error("Expected DryRun to be true")
	}
	if !opts.Force {
		t.Error("Expected Force to be true")
	}
	if opts.Agents != "unity-6000-expert,unity-graphics-expert" {
		t.Errorf("Expected agents 'unity-6000-expert,unity-graphics-expert', got %q", opts.Agents)
	}
}

func TestParseInstallFlags_EmptyAgents(t *testing.T) {
	opts, err := ParseInstallFlags([]string{})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if opts.Agents != "" {
		t.Errorf("Expected empty agents, got %q", opts.Agents)
	}
}

func TestParseAgentList_CommaSeparated(t *testing.T) {
	result := ParseAgentList("unity-6000-expert,unity-graphics-expert,unity-physics-expert")
	expected := []string{"unity-6000-expert", "unity-graphics-expert", "unity-physics-expert"}

	if len(result) != len(expected) {
		t.Fatalf("Expected %d agents, got %d", len(expected), len(result))
	}
	for i, agent := range expected {
		if result[i] != agent {
			t.Errorf("Expected agent %q at index %d, got %q", agent, i, result[i])
		}
	}
}

func TestParseAgentList_WithSpaces(t *testing.T) {
	result := ParseAgentList("unity-6000-expert, unity-graphics-expert, unity-physics-expert")
	expected := []string{"unity-6000-expert", "unity-graphics-expert", "unity-physics-expert"}

	if len(result) != len(expected) {
		t.Fatalf("Expected %d agents, got %d", len(expected), len(result))
	}
	for i, agent := range expected {
		if result[i] != agent {
			t.Errorf("Expected agent %q at index %d, got %q", agent, i, result[i])
		}
	}
}

func TestParseAgentList_SingleAgent(t *testing.T) {
	result := ParseAgentList("unity-6000-expert")
	if len(result) != 1 {
		t.Fatalf("Expected 1 agent, got %d", len(result))
	}
	if result[0] != "unity-6000-expert" {
		t.Errorf("Expected 'unity-6000-expert', got %q", result[0])
	}
}

func TestParseAgentList_EmptyString(t *testing.T) {
	result := ParseAgentList("")
	if result != nil {
		t.Errorf("Expected nil for empty string, got %v", result)
	}
}

func TestParseAgentList_EmptyParts(t *testing.T) {
	result := ParseAgentList("unity-6000-expert,,unity-graphics-expert")
	if len(result) != 2 {
		t.Errorf("Expected 2 agents, got %d: %v", len(result), result)
	}
}

func TestParseSyncFlags_ValidProviders(t *testing.T) {
	testCases := []struct {
		provider string
		wantErr  bool
	}{
		{"opencode", false},
		{"claude", false},
		{"gpt", false},
		{"gemini", false},
		{"custom", false},
		{"", false}, // empty is allowed for sync
	}

	for _, tc := range testCases {
		t.Run(tc.provider, func(t *testing.T) {
			opts, err := ParseSyncFlags([]string{"--provider", tc.provider})
			if tc.wantErr && err == nil {
				t.Errorf("Expected error for provider %q, got nil", tc.provider)
			}
			if !tc.wantErr && err != nil {
				t.Errorf("Unexpected error for provider %q: %v", tc.provider, err)
			}
			if !tc.wantErr && opts.Provider != tc.provider {
				t.Errorf("Expected provider %q, got %q", tc.provider, opts.Provider)
			}
		})
	}
}

func TestParseSyncFlags_InvalidProvider(t *testing.T) {
	_, err := ParseSyncFlags([]string{"--provider", "invalid-provider"})
	if err == nil {
		t.Error("Expected error for invalid provider")
	}
}

func TestParseSyncFlags_DryRunAndForce(t *testing.T) {
	opts, err := ParseSyncFlags([]string{"--dry-run", "--force"})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !opts.DryRun {
		t.Error("Expected DryRun to be true")
	}
	if !opts.Force {
		t.Error("Expected Force to be true")
	}
}

func TestParseSyncFlags_NoProvider(t *testing.T) {
	opts, err := ParseSyncFlags([]string{})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if opts.Provider != "" {
		t.Errorf("Expected empty provider, got %q", opts.Provider)
	}
}