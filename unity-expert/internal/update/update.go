package update

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// UpdateOptions holds parsed flags for the update command.
type UpdateOptions struct {
	CheckOnly bool
}

// Orchestrator manages the update pipeline.
type Orchestrator struct {
	currentVersion string
	opts           UpdateOptions
	githubOwner    string
	githubRepo     string
}

// NewOrchestrator creates a new update Orchestrator.
func NewOrchestrator(currentVersion string, opts UpdateOptions) *Orchestrator {
	return &Orchestrator{
		currentVersion: currentVersion,
		opts:           opts,
		githubOwner:    "Ulysses-Alv",
		githubRepo:     "Unity-Agent-Expert",
	}
}

// Execute runs the full update pipeline.
// Pipeline steps:
// 1. CheckGentleAI() → abort if managed env
// 2. FetchLatestRelease() → GET api.github.com/.../releases/latest
// 3. CompareVersions(current, latest) → semver comparison
// 4. SelectAsset(release, GOOS, GOARCH) → matching zip asset
// 5. [if --check-only] → report available version → exit 0
// 6. DownloadAsset(assetURL, tempPath) → HTTP GET → write temp file
// 7. ExtractBinary(zipPath, destDir, binaryName) → archive/zip → extract binary
// 8. ReplaceBinary(currentPath, newPath) → atomic swap with .old backup
func (o *Orchestrator) Execute() error {
	// Step 1: Check gentle-ai environment
	if err := CheckGentleAI(); err != nil {
		return err
	}

	// Step 2: Fetch latest release from GitHub
	release, err := FetchLatestRelease(o.githubOwner, o.githubRepo)
	if err != nil {
		return err
	}

	// Parse the latest version from tag name
	latestVersionStr := ParseTagName(release.TagName)

	// Step 3: Compare versions
	// If current is "dev" or not valid semver, treat it as older than any release
	if IsDevVersion(o.currentVersion) {
		fmt.Printf("[INFO] Current version is '%s' (dev build), checking for updates...\n", o.currentVersion)
	} else {
		currentVer, err := ParseVersion(o.currentVersion)
		if err != nil {
			return fmt.Errorf("failed to parse current version: %w", err)
		}

		latestVer, err := ParseVersion(latestVersionStr)
		if err != nil {
			return fmt.Errorf("failed to parse latest version: %w", err)
		}

		cmp := CompareVersions(currentVer, latestVer)
		if cmp >= 0 {
			// Already on latest or newer
			fmt.Printf("Ya tenés la última versión (v%s)\n", latestVersionStr)
			return nil
		}

		fmt.Printf("[INFO] New version available: v%s (current: v%s)\n", latestVersionStr, o.currentVersion)
	}

	// Step 4: Select the correct asset for this platform
	asset, err := SelectAsset(release, runtime.GOOS, runtime.GOARCH)
	if err != nil {
		return err
	}

	// Step 5: If check-only mode, report and exit
	if o.opts.CheckOnly {
		fmt.Printf("Nueva versión disponible: v%s\n", latestVersionStr)
		fmt.Println("Ejecutá 'unity-agent-expert update' para actualizar.")
		return nil
	}

	// Get current binary path
	currentPath, err := GetCurrentBinaryPath()
	if err != nil {
		return fmt.Errorf("failed to determine current binary path: %w", err)
	}

	// Setup temp paths
	tempZipPath := GetTempFilePath(currentPath, ".zip")
	tempDir := filepath.Dir(currentPath)

	// Cleanup function
	cleanup := func() {
		os.Remove(tempZipPath)
	}

	defer cleanup()

	// Step 6: Download the asset
	fmt.Printf("[INFO] Descargando actualización desde %s...\n", asset.DownloadURL)
	if err := DownloadAsset(asset.DownloadURL, tempZipPath); err != nil {
		return err
	}

	// Step 7: Extract the binary
	fmt.Println("[INFO] Extrayendo binario...")
	extractedPath, err := ExtractBinary(tempZipPath, tempDir, "unity-agent-expert")
	if err != nil {
		return err
	}

	// Step 8: Replace the binary
	fmt.Println("[INFO] Reemplazando binario...")
	if err := ReplaceBinary(currentPath, extractedPath); err != nil {
		return err
	}

	fmt.Printf("Actualizado a v%s.\n", latestVersionStr)
	fmt.Println("Ejecutá install para instalar los agentes.")

	return nil
}
