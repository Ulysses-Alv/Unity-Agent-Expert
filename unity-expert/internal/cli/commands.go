package cli

import (
	"fmt"

	"github.com/Ulysses-Alv/unity-expert/internal/install"
)

// PrintUsage prints the CLI usage information.
func PrintUsage() {
	fmt.Println("Usage: unity-expert <command> [flags]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  install    Install Unity agents, prompts, and skills")
	fmt.Println("  sync       Idempotently re-apply configuration")
	fmt.Println("  doctor     Check installation health")
	fmt.Println("  version    Display version information")
	fmt.Println()
	fmt.Println("Run 'unity-expert help' for detailed usage information.")
}

// printInstallUsage prints detailed install command usage.
func printInstallUsage() {
	fmt.Println("Usage: unity-expert install [flags]")
	fmt.Println()
	fmt.Println("Installs Unity agents, prompts, and skills into the OpenCode configuration.")
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Println("  --provider string   Model provider preset: opencode, claude, gpt, gemini, custom (default: opencode)")
	fmt.Println("  --dry-run          Preview changes without applying")
	fmt.Println("  --force            Overwrite existing Unity agents")
	fmt.Println("  --agents string    Comma-separated list of specific agents to install")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  unity-expert install --provider claude")
	fmt.Println("  unity-expert install --provider gpt --dry-run")
	fmt.Println("  unity-expert install --agents unity-6000-expert,unity-graphics-expert")
}

// printSyncUsage prints detailed sync command usage.
func printSyncUsage() {
	fmt.Println("Usage: unity-expert sync [flags]")
	fmt.Println()
	fmt.Println("Idempotently re-applies configuration without reinstalling binaries.")
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Println("  --provider string   Update model provider preset (optional)")
	fmt.Println("  --dry-run           Preview changes without applying")
	fmt.Println("  --force             Force overwrite even if no drift detected")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  unity-expert sync")
	fmt.Println("  unity-expert sync --provider claude --dry-run")
}

// printDoctorUsage prints detailed doctor command usage.
func printDoctorUsage() {
	fmt.Println("Usage: unity-expert doctor")
	fmt.Println()
	fmt.Println("Checks the health of the Unity Agent Expert installation.")
	fmt.Println()
	fmt.Println("Exit codes:")
	fmt.Println("  0   All checks passed")
	fmt.Println("  1   One or more checks failed")
}

// RunInstall is the entry point for the install command.
func RunInstall(args []string) error {
	opts, err := ParseInstallFlags(args)
	if err != nil {
		return err
	}

	fmt.Printf("[INFO] Install command called with provider=%s, dryRun=%v, force=%v, agents=%s\n",
		opts.Provider, opts.DryRun, opts.Force, opts.Agents)

	// Build install options for pipeline
	installOpts := install.InstallOptions{
		Provider: opts.Provider,
		DryRun:   opts.DryRun,
		Force:    opts.Force,
		Agents:   ParseAgentList(opts.Agents),
	}

	// Create and execute pipeline
	pipeline := install.NewPipeline(installOpts)
	return pipeline.Execute()
}
