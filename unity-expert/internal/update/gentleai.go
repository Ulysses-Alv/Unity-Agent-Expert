package update

import (
	"os"
	"path/filepath"
	"strings"
)

// CheckGentleAI checks if the binary is running in a gentle-ai managed environment.
// Returns a GentleAIError if detected, nil otherwise.
func CheckGentleAI() error {
	// Check for GENTLE_AI_ENV environment variable
	if _, exists := os.LookupEnv("GENTLE_AI_ENV"); exists {
		return &GentleAIError{}
	}

	// Check if any of the executable paths contain gentle-ai
	pathsToCheck := []string{
		os.Args[0],
	}

	exePath, err := os.Executable()
	if err == nil {
		pathsToCheck = append(pathsToCheck, exePath)
	}

	for _, path := range pathsToCheck {
		lowerPath := strings.ToLower(path)
		if strings.Contains(lowerPath, "gentle-ai") {
			return &GentleAIError{}
		}

		// Also check home directory for gentle-ai
		home, err := os.UserHomeDir()
		if err == nil {
			gentleAIPath := filepath.Join(home, "gentle-ai")
			if strings.HasPrefix(lowerPath, strings.ToLower(gentleAIPath)) {
				return &GentleAIError{}
			}
		}
	}

	return nil
}

// IsGentleAI returns true if the current environment is detected as gentle-ai managed.
func IsGentleAI() bool {
	return CheckGentleAI() != nil
}
