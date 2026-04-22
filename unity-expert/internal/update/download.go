package update

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

// DownloadAsset downloads a file from the given URL to the destination path.
// Returns an error if the download fails.
func DownloadAsset(url, destPath string) error {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return &DownloadError{URL: url, Err: err}
	}

	req.Header.Set("Accept", "application/octet-stream")
	req.Header.Set("User-Agent", "unity-agent-expert-update-checker")

	resp, err := client.Do(req)
	if err != nil {
		return &DownloadError{URL: url, Err: err}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return &DownloadError{URL: url, Err: fmt.Errorf("HTTP %d", resp.StatusCode)}
	}

	// Create the destination file
	outFile, err := os.Create(destPath)
	if err != nil {
		return &DownloadError{URL: url, Err: fmt.Errorf("failed to create file: %w", err)}
	}
	defer outFile.Close()

	// Copy the response body to the file
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return &DownloadError{URL: url, Err: fmt.Errorf("failed to write file: %w", err)}
	}

	return nil
}

// ExtractBinary extracts the binary from a zip archive.
// The zip is expected to contain a single file (the binary itself).
// Returns the path to the extracted binary.
func ExtractBinary(zipPath, destDir, binaryName string) (string, error) {
	reader, err := zip.OpenReader(zipPath)
	if err != nil {
		return "", &ExtractError{Path: zipPath, Err: fmt.Errorf("failed to open zip: %w", err)}
	}
	defer reader.Close()

	// Ensure destination directory exists
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return "", &ExtractError{Path: zipPath, Err: fmt.Errorf("failed to create directory: %w", err)}
	}

	// Find the binary in the archive
	var extractedPath string
	found := false
	for _, f := range reader.File {
		// Check if this file looks like the binary
		if !f.FileInfo().IsDir() && (f.Name == binaryName || filepath.Base(f.Name) == binaryName) {
			found = true
			extractedPath = filepath.Join(destDir, filepath.Base(f.Name))
			if err := extractFile(f, extractedPath); err != nil {
				return "", &ExtractError{Path: zipPath, Err: err}
			}
			break
		}
	}

	// If no specific binary found, try to extract the first non-directory file
	if !found {
		for _, f := range reader.File {
			if !f.FileInfo().IsDir() {
				found = true
				extractedPath = filepath.Join(destDir, filepath.Base(f.Name))
				if err := extractFile(f, extractedPath); err != nil {
					return "", &ExtractError{Path: zipPath, Err: err}
				}
				break
			}
		}
	}

	if !found {
		return "", &ExtractError{Path: zipPath, Err: fmt.Errorf("no files found in archive")}
	}

	return extractedPath, nil
}

// extractFile extracts a single file from a zip archive.
func extractFile(f *zip.File, destPath string) error {
	reader, err := f.Open()
	if err != nil {
		return fmt.Errorf("failed to open file in archive: %w", err)
	}
	defer reader.Close()

	// Create the destination file
	outFile, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer outFile.Close()

	// Make the file executable
	if err := os.Chmod(destPath, 0755); err != nil {
		return fmt.Errorf("failed to set permissions: %w", err)
	}

	_, err = io.Copy(outFile, reader)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// ReplaceBinary atomically replaces the current binary with a new one.
// It first renames the current binary to .old, then renames the new binary to the current path.
// On failure, it attempts to restore from the .old backup.
func ReplaceBinary(currentPath, newPath string) error {
	oldPath := currentPath + ".old"

	// Remove any existing .old file first
	if _, err := os.Stat(oldPath); err == nil {
		if err := os.Remove(oldPath); err != nil {
			return &ReplaceError{Path: oldPath, Err: fmt.Errorf("failed to remove old backup: %w", err)}
		}
	}

	// Step 1: Rename current binary to .old
	if err := os.Rename(currentPath, oldPath); err != nil {
		return &ReplaceError{Path: currentPath, Err: fmt.Errorf("failed to rename current binary: %w", err)}
	}

	// Step 2: Rename new binary to current path
	if err := os.Rename(newPath, currentPath); err != nil {
		// Rollback: restore the .old file
		rollbackErr := os.Rename(oldPath, currentPath)
		if rollbackErr != nil {
			return &ReplaceError{
				Path: currentPath,
				Err:  fmt.Errorf("replace failed and rollback also failed: original error: %v, rollback error: %w", err, rollbackErr),
			}
		}
		return &ReplaceError{Path: currentPath, Err: fmt.Errorf("failed to rename new binary: %w", err)}
	}

	return nil
}

// GetCurrentBinaryPath returns the path to the currently running binary.
func GetCurrentBinaryPath() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", &ReplaceError{Path: "<unknown>", Err: fmt.Errorf("failed to get executable path: %w", err)}
	}

	// On Windows, the executable might have a .exe extension
	if runtime.GOOS == "windows" && filepath.Ext(exePath) != ".exe" {
		exePath = exePath + ".exe"
	}

	return exePath, nil
}

// GetTempFilePath returns a path for a temporary file in the same directory as the binary.
func GetTempFilePath(binaryPath, suffix string) string {
	dir := filepath.Dir(binaryPath)
	return filepath.Join(dir, fmt.Sprintf("unity-agent-expert_update_%d%s", os.Getpid(), suffix))
}
