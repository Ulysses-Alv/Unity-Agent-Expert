package cli

import (
	"fmt"

	"github.com/Ulysses-Alv/unity-expert/internal/install"
)

// PrintUsage prints the CLI usage information.
func PrintUsage() {
	fmt.Println("Usage: unity-agent-expert <command> [flags]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  install    Install Unity agents, prompts, and skills")
	fmt.Println("  sync       Idempotently re-apply configuration")
	fmt.Println("  update     Check for and install updates")
	fmt.Println("  doctor     Check installation health")
	fmt.Println("  version    Display version information")
	fmt.Println()
	fmt.Println("Run 'unity-agent-expert help' for detailed usage information.")
}

// printInstallUsage prints detailed install command usage.
func printInstallUsage() {
	fmt.Println("Usage: unity-agent-expert install [flags]")
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
	fmt.Println("  unity-agent-expert install --provider claude")
	fmt.Println("  unity-agent-expert install --provider gpt --dry-run")
	fmt.Println("  unity-agent-expert install --agents unity-6000-expert,unity-graphics-expert")
}

// printSyncUsage prints detailed sync command usage.
func printSyncUsage() {
	fmt.Println("Usage: unity-agent-expert sync [flags]")
	fmt.Println()
	fmt.Println("Idempotently re-applies configuration without reinstalling binaries.")
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Println("  --provider string   Update model provider preset (optional)")
	fmt.Println("  --dry-run           Preview changes without applying")
	fmt.Println("  --force             Force overwrite even if no drift detected")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  unity-agent-expert sync")
	fmt.Println("  unity-agent-expert sync --provider claude --dry-run")
}

// printDoctorUsage prints detailed doctor command usage.
func printDoctorUsage() {
	fmt.Println("Usage: unity-agent-expert doctor")
	fmt.Println()
	fmt.Println("Checks the health of the Unity Agent Expert installation.")
	fmt.Println()
	fmt.Println("Exit codes:")
	fmt.Println("  0   All checks passed")
	fmt.Println("  1   One or more checks failed")
}

// printUpdateUsage prints detailed update command usage.
func printUpdateUsage() {
	fmt.Println("Usage: unity-agent-expert update [flags]")
	fmt.Println()
	fmt.Println("Checks GitHub for new releases and updates the binary if a newer version is available.")
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Println("  --check-only, -c   Check for updates without downloading or replacing the binary")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  unity-agent-expert update")
	fmt.Println("  unity-agent-expert update --check-only")
	fmt.Println()
	fmt.Println("Exit codes:")
	fmt.Println("  0   Success (updated, already up-to-date, or check-only)")
	fmt.Println("  1   Error (network, API, permission, no asset)")
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
