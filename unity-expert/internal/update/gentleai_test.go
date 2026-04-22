package update

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCheckGentleAI(t *testing.T) {
	// Save original values to restore later
	originalExe := os.Args[0]
	originalHome := os.Getenv("HOME")
	defer func() {
		os.Args[0] = originalExe
		if originalHome != "" {
			os.Setenv("HOME", originalHome)
		} else {
			os.Unsetenv("HOME")
		}
	}()

	t.Run("GENTLE_AI_ENV set", func(t *testing.T) {
		os.Setenv("GENTLE_AI_ENV", "true")
		defer os.Unsetenv("GENTLE_AI_ENV")

		err := CheckGentleAI()
		if err == nil {
			t.Error("CheckGentleAI() expected GentleAIError when GENTLE_AI_ENV is set, got nil")
		}
		if _, ok := err.(*GentleAIError); !ok {
			t.Errorf("CheckGentleAI() error type = %T, want *GentleAIError", err)
		}
	})

	t.Run("GENTLE_AI_ENV not set", func(t *testing.T) {
		os.Unsetenv("GENTLE_AI_ENV")

		err := CheckGentleAI()
		if err != nil {
			t.Errorf("CheckGentleAI() unexpected error: %v", err)
		}
	})

	t.Run("binary path contains gentle-ai", func(t *testing.T) {
		os.Unsetenv("GENTLE_AI_ENV")

		// Set executable path to something containing "gentle-ai"
		os.Args[0] = filepath.Join("C:", "Users", "Test", "AppData", "gentle-ai", "bin", "unity-agent-expert.exe")

		err := CheckGentleAI()
		if err == nil {
			t.Error("CheckGentleAI() expected GentleAIError when path contains gentle-ai, got nil")
		}
	})

	t.Run("binary path does not contain gentle-ai", func(t *testing.T) {
		os.Unsetenv("GENTLE_AI_ENV")

		// Set executable path to something normal
		os.Args[0] = filepath.Join("C:", "Users", "Test", "bin", "unity-agent-expert.exe")

		err := CheckGentleAI()
		if err != nil {
			t.Errorf("CheckGentleAI() unexpected error: %v", err)
		}
	})
}

func TestIsGentleAI(t *testing.T) {
	// Save original values
	originalEnv := os.Getenv("GENTLE_AI_ENV")
	defer func() {
		if originalEnv != "" {
			os.Setenv("GENTLE_AI_ENV", originalEnv)
		} else {
			os.Unsetenv("GENTLE_AI_ENV")
		}
	}()

	t.Run("returns true when gentle-ai detected", func(t *testing.T) {
		os.Setenv("GENTLE_AI_ENV", "true")
		defer os.Unsetenv("GENTLE_AI_ENV")

		if !IsGentleAI() {
			t.Error("IsGentleAI() expected true when GENTLE_AI_ENV is set")
		}
	})

	t.Run("returns false when not gentle-ai", func(t *testing.T) {
		os.Unsetenv("GENTLE_AI_ENV")
		// Also reset executable path
		os.Args[0] = filepath.Join("C:", "test", "bin", "test.exe")

		if IsGentleAI() {
			t.Error("IsGentleAI() expected false when not gentle-ai environment")
		}
	})
}
