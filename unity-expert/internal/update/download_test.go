package update

import (
	"archive/zip"
	"os"
	"path/filepath"
	"testing"
)

func TestExtractBinary(t *testing.T) {
	t.Run("extracts binary from zip", func(t *testing.T) {
		// Create a temp directory
		tempDir, err := os.MkdirTemp("", "update-test-*")
		if err != nil {
			t.Fatalf("Failed to create temp dir: %v", err)
		}
		defer os.RemoveAll(tempDir)

		// Create a zip file with a binary
		zipPath := filepath.Join(tempDir, "test.zip")
		zipWriter, err := os.Create(zipPath)
		if err != nil {
			t.Fatalf("Failed to create zip file: %v", err)
		}

		zipArchive := zip.NewWriter(zipWriter)
		binaryContent := []byte("#!/bin/bash\necho hello")
		header := &zip.FileHeader{
			Name:   "unity-agent-expert",
			Method: zip.Deflate,
		}
		writer, err := zipArchive.CreateHeader(header)
		if err != nil {
			t.Fatalf("Failed to create zip entry: %v", err)
		}
		writer.Write(binaryContent)
		zipArchive.Close()
		zipWriter.Close()

		// Extract the binary
		destDir := filepath.Join(tempDir, "dest")
		extractedPath, err := ExtractBinary(zipPath, destDir, "unity-agent-expert")
		if err != nil {
			t.Fatalf("ExtractBinary() error = %v", err)
		}

		// Verify the file exists and has content
		data, err := os.ReadFile(extractedPath)
		if err != nil {
			t.Fatalf("Failed to read extracted file: %v", err)
		}
		if string(data) != string(binaryContent) {
			t.Errorf("Extracted content = %q, want %q", string(data), string(binaryContent))
		}
	})

	t.Run("handles empty zip", func(t *testing.T) {
		tempDir, err := os.MkdirTemp("", "update-test-*")
		if err != nil {
			t.Fatalf("Failed to create temp dir: %v", err)
		}
		defer os.RemoveAll(tempDir)

		zipPath := filepath.Join(tempDir, "empty.zip")
		zipWriter, err := os.Create(zipPath)
		if err != nil {
			t.Fatalf("Failed to create zip file: %v", err)
		}
		zip.NewWriter(zipWriter).Close()
		zipWriter.Close()

		destDir := filepath.Join(tempDir, "dest")
		_, err = ExtractBinary(zipPath, destDir, "unity-agent-expert")
		if err == nil {
			t.Error("ExtractBinary() expected error for empty zip, got nil")
		}
	})
}

func TestReplaceBinary(t *testing.T) {
	t.Run("successful binary replacement", func(t *testing.T) {
		tempDir, err := os.MkdirTemp("", "update-test-*")
		if err != nil {
			t.Fatalf("Failed to create temp dir: %v", err)
		}
		defer os.RemoveAll(tempDir)

		// Create a fake current binary
		currentPath := filepath.Join(tempDir, "unity-agent-expert")
		if err := os.WriteFile(currentPath, []byte("old binary"), 0644); err != nil {
			t.Fatalf("Failed to create current binary: %v", err)
		}

		// Create a fake new binary
		newPath := filepath.Join(tempDir, "new-binary")
		if err := os.WriteFile(newPath, []byte("new binary"), 0644); err != nil {
			t.Fatalf("Failed to create new binary: %v", err)
		}

		// Replace
		if err := ReplaceBinary(currentPath, newPath); err != nil {
			t.Fatalf("ReplaceBinary() error = %v", err)
		}

		// Verify current binary has new content
		data, err := os.ReadFile(currentPath)
		if err != nil {
			t.Fatalf("Failed to read current binary: %v", err)
		}
		if string(data) != "new binary" {
			t.Errorf("Current binary content = %q, want %q", string(data), "new binary")
		}

		// Verify .old backup exists with old content
		oldPath := currentPath + ".old"
		oldData, err := os.ReadFile(oldPath)
		if err != nil {
			t.Fatalf("Failed to read .old backup: %v", err)
		}
		if string(oldData) != "old binary" {
			t.Errorf(".old backup content = %q, want %q", string(oldData), "old binary")
		}

		// Verify new binary no longer exists at original location
		if _, err := os.Stat(newPath); !os.IsNotExist(err) {
			t.Error("New binary should have been moved")
		}
	})

	t.Run("rollback on failure", func(t *testing.T) {
		tempDir, err := os.MkdirTemp("", "update-test-*")
		if err != nil {
			t.Fatalf("Failed to create temp dir: %v", err)
		}
		defer os.RemoveAll(tempDir)

		// Create a fake current binary
		currentPath := filepath.Join(tempDir, "unity-agent-expert")
		if err := os.WriteFile(currentPath, []byte("old binary"), 0644); err != nil {
			t.Fatalf("Failed to create current binary: %v", err)
		}

		// Create a fake new binary that we'll corrupt
		newPath := filepath.Join(tempDir, "new-binary")
		if err := os.WriteFile(newPath, []byte("new binary"), 0644); err != nil {
			t.Fatalf("Failed to create new binary: %v", err)
		}

		// Make current binary read-only to cause rename failure on some systems
		// Actually, we'll simulate failure by removing newPath first
		os.Remove(newPath)

		err = ReplaceBinary(currentPath, newPath)
		if err == nil {
			t.Error("ReplaceBinary() expected error when new binary doesn't exist")
		}

		// Verify current binary is still the old one (rollback worked)
		data, err := os.ReadFile(currentPath)
		if err != nil {
			t.Fatalf("Failed to read current binary after failed replace: %v", err)
		}
		if string(data) != "old binary" {
			t.Errorf("Current binary should still be old after failed replace, got %q", string(data))
		}
	})
}

func TestGetCurrentBinaryPath(t *testing.T) {
	path, err := GetCurrentBinaryPath()
	if err != nil {
		t.Errorf("GetCurrentBinaryPath() error = %v", err)
	}
	if path == "" {
		t.Error("GetCurrentBinaryPath() returned empty path")
	}
}

func TestGetTempFilePath(t *testing.T) {
	binaryPath := filepath.Join("C:", "Users", "Test", "bin", "unity-agent-expert")
	suffix := ".zip"

	tempPath := GetTempFilePath(binaryPath, suffix)

	if tempPath == "" {
		t.Error("GetTempFilePath() returned empty path")
	}

	expectedDir := filepath.Join("C:", "Users", "Test", "bin")
	if filepath.Dir(tempPath) != expectedDir {
		t.Errorf("GetTempFilePath() dir = %q, want %q", filepath.Dir(tempPath), expectedDir)
	}
}
