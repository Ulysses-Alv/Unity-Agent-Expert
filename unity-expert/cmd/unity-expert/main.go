package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Ulysses-Alv/unity-expert/internal/cli"
	"github.com/Ulysses-Alv/unity-expert/internal/install"
)

// Version is injected at build time via -ldflags.
var version = "dev"

func main() {
	if len(os.Args) < 2 {
		cli.PrintUsage()
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]

	// Handle --version flag regardless of command position
	if command == "--version" || command == "-v" {
		fmt.Printf("unity-agent-expert v%s\n", version)
		return
	}

	var err error
	switch command {
	case "install":
		err = cli.RunInstall(args)
	case "sync":
		err = cli.RunSync(args)
	case "doctor":
		err = cli.RunDoctor()
	case "version":
		fmt.Printf("unity-agent-expert v%s\n", version)
		return
	case "help", "--help", "-h":
		cli.PrintUsage()
		return
	default:
		fmt.Fprintf(os.Stderr, "[ERROR] Unknown command: %s\n", command)
		cli.PrintUsage()
		os.Exit(1)
	}

	if err != nil {
		// Don't print anything for help (flag.ErrHelp)
		if err == flag.ErrHelp {
			return
		}

		exitCode := 1
		switch err.(type) {
		case *install.BackupError:
			exitCode = 2
		case *install.RollbackError:
			exitCode = 2
		case *install.PreconditionError:
			exitCode = 3
		case *install.ApplyError:
			exitCode = 4
		}
		fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
		os.Exit(exitCode)
	}
}
