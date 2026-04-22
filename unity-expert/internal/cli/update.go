package cli

import (
	"fmt"

	"github.com/Ulysses-Alv/unity-expert/internal/update"
)

// RunUpdate is the entry point for the update command.
func RunUpdate(args []string, version string) error {
	opts, err := ParseUpdateFlags(args)
	if err != nil {
		return err
	}

	fmt.Printf("[INFO] Update command called with version=%s, checkOnly=%v\n",
		version, opts.CheckOnly)

	// Create and execute the update orchestrator
	orchestrator := update.NewOrchestrator(version, update.UpdateOptions{
		CheckOnly: opts.CheckOnly,
	})

	return orchestrator.Execute()
}
