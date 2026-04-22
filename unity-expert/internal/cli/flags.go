package cli

import (
	"flag"
	"fmt"
	"strings"
)

// InstallOptions holds parsed flags for the install command.
type InstallOptions struct {
	Provider string
	DryRun   bool
	Force    bool
	Agents   string // CSV of agent names
}

// SyncOptions holds parsed flags for the sync command.
type SyncOptions struct {
	Provider string
	DryRun   bool
	Force    bool
}

// UpdateOptions holds parsed flags for the update command.
type UpdateOptions struct {
	CheckOnly bool
}

var validProviders = map[string]bool{
	"opencode": true,
	"claude":   true,
	"gpt":      true,
	"gemini":   true,
	"custom":   true,
}

// ParseInstallFlags parses command-line flags for the install command.
func ParseInstallFlags(args []string) (InstallOptions, error) {
	// Check for help flag
	for _, arg := range args {
		if arg == "-h" || arg == "--help" {
			printInstallUsage()
			return InstallOptions{}, flag.ErrHelp
		}
	}

	fs := flag.NewFlagSet("install", flag.ContinueOnError)
	fs.Usage = func() {
		fmt.Println("Usage: unity-agent-expert install [flags]")
		fmt.Println("Flags:")
		fs.PrintDefaults()
	}

	provider := fs.String("provider", "opencode", "Model provider preset: opencode, claude, gpt, gemini, custom")
	dryRun := fs.Bool("dry-run", false, "Preview changes without applying")
	force := fs.Bool("force", false, "Overwrite existing Unity agents")
	agents := fs.String("agents", "", "Comma-separated list of specific agents to install")

	if err := fs.Parse(args); err != nil {
		return InstallOptions{}, fmt.Errorf("invalid flags: %w", err)
	}

	if !validProviders[*provider] {
		return InstallOptions{}, fmt.Errorf(
			"invalid --provider '%s'. Known providers: opencode, claude, gpt, gemini, custom",
			*provider)
	}

	return InstallOptions{
		Provider: *provider,
		DryRun:   *dryRun,
		Force:    *force,
		Agents:   *agents,
	}, nil
}

// ParseSyncFlags parses command-line flags for the sync command.
func ParseSyncFlags(args []string) (SyncOptions, error) {
	// Check for help flag
	for _, arg := range args {
		if arg == "-h" || arg == "--help" {
			printSyncUsage()
			return SyncOptions{}, flag.ErrHelp
		}
	}

	fs := flag.NewFlagSet("sync", flag.ContinueOnError)
	fs.Usage = func() {
		fmt.Println("Usage: unity-agent-expert sync [flags]")
		fmt.Println("Flags:")
		fs.PrintDefaults()
	}

	provider := fs.String("provider", "", "Update model provider preset (optional)")
	dryRun := fs.Bool("dry-run", false, "Preview changes without applying")
	force := fs.Bool("force", false, "Force overwrite even if no drift detected")

	if err := fs.Parse(args); err != nil {
		return SyncOptions{}, fmt.Errorf("invalid flags: %w", err)
	}

	if *provider != "" && !validProviders[*provider] {
		return SyncOptions{}, fmt.Errorf(
			"invalid --provider '%s'. Known providers: opencode, claude, gpt, gemini, custom",
			*provider)
	}

	return SyncOptions{
		Provider: *provider,
		DryRun:   *dryRun,
		Force:    *force,
	}, nil
}

// ParseUpdateFlags parses command-line flags for the update command.
func ParseUpdateFlags(args []string) (UpdateOptions, error) {
	// Check for help flag
	for _, arg := range args {
		if arg == "-h" || arg == "--help" {
			printUpdateUsage()
			return UpdateOptions{}, flag.ErrHelp
		}
	}

	fs := flag.NewFlagSet("update", flag.ContinueOnError)
	fs.Usage = func() {
		fmt.Println("Usage: unity-agent-expert update [flags]")
		fmt.Println("Flags:")
		fs.PrintDefaults()
	}

	checkOnly := fs.Bool("check-only", false, "Check for updates without downloading")

	// Also register -c as an alias for --check-only
	// Since we can't easily add aliases, we manually handle -c in parsing

	if err := fs.Parse(args); err != nil {
		return UpdateOptions{}, fmt.Errorf("invalid flags: %w", err)
	}

	// Check for -c flag manually since flag package doesn't support aliases well
	for _, arg := range args {
		if arg == "-c" || arg == "--c" {
			*checkOnly = true
		}
	}

	return UpdateOptions{
		CheckOnly: *checkOnly,
	}, nil
}

// ParseAgentList parses a CSV string into a slice of agent names.
func ParseAgentList(csv string) []string {
	if csv == "" {
		return nil
	}
	parts := strings.Split(csv, ",")
	agents := make([]string, 0, len(parts))
	for _, p := range parts {
		trimmed := strings.TrimSpace(p)
		if trimmed != "" {
			agents = append(agents, trimmed)
		}
	}
	return agents
}
