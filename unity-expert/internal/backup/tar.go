package backup

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"time"
)

// ManifestEntry represents a single file in the backup manifest.
type ManifestEntry struct {
	Path      string `json:"path"`
	Size      int64  `json:"size"`
	Mode      int64  `json:"mode"`
	Timestamp string `json:"timestamp"`
}

// Manifest represents the backup manifest.
type Manifest struct {
	Timestamp   string         `json:"timestamp"`
	Command     string         `json:"command"`
	Version     string         `json:"version"`
	Files       []ManifestEntry `json:"files"`
	PackageName string         `json:"packageName,omitempty"`
}

// CreateBackup creates a backup of the opencode.json file.
func CreateBackup(configPath, backupDir, command string) (string, error) {
	// Ensure backup directory exists
	if err := os.MkdirAll(backupDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create backup directory: %w", err)
	}

	// Generate backup name with timestamp
	timestamp := time.Now().Format("20060102-150405")
	backupName := fmt.Sprintf("backup-%s.tar.gz", timestamp)
	backupPath := filepath.Join(backupDir, backupName)

	// Open backup file for writing
	f, err := os.Create(backupPath)
	if err != nil {
		return "", fmt.Errorf("failed to create backup file: %w", err)
	}
	defer f.Close()

	gz := gzip.NewWriter(f)
	defer gz.Close()

	tw := tar.NewWriter(gz)
	defer tw.Close()

	// Read the config file
	configData, err := os.ReadFile(configPath)
	if err != nil {
		return "", fmt.Errorf("failed to read config file: %w", err)
	}

	// Create tar header for the config file
	header := &tar.Header{
		Name:    filepath.Base(configPath),
		Size:    int64(len(configData)),
		Mode:    0644,
		ModTime: time.Now(),
	}

	if err := tw.WriteHeader(header); err != nil {
		return "", fmt.Errorf("failed to write tar header: %w", err)
	}

	if _, err := tw.Write(configData); err != nil {
		return "", fmt.Errorf("failed to write config data: %w", err)
	}

	// Create manifest
	manifest := Manifest{
		Timestamp: timestamp,
		Command:    command,
		Version:    "dev",
		Files: []ManifestEntry{
			{
				Path:      configPath,
				Size:      int64(len(configData)),
				Mode:      0644,
				Timestamp: time.Now().Format(time.RFC3339),
			},
		},
	}

	manifestData, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal manifest: %w", err)
	}

	// Write manifest to tar
	header = &tar.Header{
		Name:    "manifest.json",
		Size:    int64(len(manifestData)),
		Mode:    0644,
		ModTime: time.Now(),
	}

	if err := tw.WriteHeader(header); err != nil {
		return "", fmt.Errorf("failed to write manifest header: %w", err)
	}

	if _, err := tw.Write(manifestData); err != nil {
		return "", fmt.Errorf("failed to write manifest data: %w", err)
	}

	// Close tar writer to flush
	if err := tw.Close(); err != nil {
		return "", fmt.Errorf("failed to close tar writer: %w", err)
	}

	// Prune old backups
	if err := pruneBackups(backupDir, 5); err != nil {
		// Log but don't fail the backup
		fmt.Printf("[WARN] Failed to prune old backups: %v\n", err)
	}

	return backupPath, nil
}

// RestoreFromBackup restores the opencode.json from a backup.
func RestoreFromBackup(backupPath, restorePath string) error {
	// Open the backup file
	f, err := os.Open(backupPath)
	if err != nil {
		return fmt.Errorf("failed to open backup file: %w", err)
	}
	defer f.Close()

	gz, err := gzip.NewReader(f)
	if err != nil {
		return fmt.Errorf("failed to create gzip reader: %w", err)
	}
	defer gz.Close()

	tr := tar.NewReader(gz)

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read tar header: %w", err)
		}

		switch header.Name {
		case "manifest.json":
			// Skip manifest during restore
			continue
		default:
			// This is the config file
			data, err := io.ReadAll(tr)
			if err != nil {
				return fmt.Errorf("failed to read file from tar: %w", err)
			}

			// Ensure directory exists
			dir := filepath.Dir(restorePath)
			if err := os.MkdirAll(dir, 0755); err != nil {
				return fmt.Errorf("failed to create restore directory: %w", err)
			}

			// Write the file
			if err := os.WriteFile(restorePath, data, 0644); err != nil {
				return fmt.Errorf("failed to write restored file: %w", err)
			}
		}
	}

	return nil
}

// pruneBackups keeps only the most recent count backups.
func pruneBackups(backupDir string, keep int) error {
	entries, err := os.ReadDir(backupDir)
	if err != nil {
		return err
	}

	// Collect backup files
	var backups []os.FileInfo
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		info, err := entry.Info()
		if err != nil {
			continue
		}
		if !info.Mode().IsRegular() {
			continue
		}
		backups = append(backups, info)
	}

	// Sort by modification time (newest first)
	sort.Slice(backups, func(i, j int) bool {
		return backups[i].ModTime().After(backups[j].ModTime())
	})

	// Delete oldest backups beyond keep count
	for i := keep; i < len(backups); i++ {
		backupPath := filepath.Join(backupDir, backups[i].Name())
		if err := os.Remove(backupPath); err != nil {
			fmt.Printf("[WARN] Failed to delete old backup %s: %v\n", backupPath, err)
		}
	}

	return nil
}
