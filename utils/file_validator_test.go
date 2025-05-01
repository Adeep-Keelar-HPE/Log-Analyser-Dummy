package utils

import (
	"os"
	"testing"
)

// Test IsValidFile function.
func TestIsValidFile(t *testing.T) {
	// -------------------------------------------------------
	// Test One -- Valid File Existence and Check.
	tmpFile, err := os.CreateTemp("", "testfile.log")
	if err != nil {
		t.Fatalf("Failed to create the temporary file: %v", err)
	}
	defer os.Remove((tmpFile.Name())) // Clean up the temp file.

	if !IsValidFile(tmpFile.Name()) {
		t.Errorf("Expected True for the valid file, Got False...")
	}
	// -------------------------------------------------------

	// -------------------------------------------------------
	// Test Two -- Invalid File Existence and Check.
	if IsValidFile("NonExistentFile.log") {
		t.Errorf("Expected false for the non-existent file, got true...")
	}
	// -------------------------------------------------------
	// -------------------------------------------------------
	// Test Three -- Directory Check.
	tmpDir, err := os.MkdirTemp("", "testdir.log")
	if err != nil {
		t.Fatalf("Failed to create the temporary directory: %v", err)
	}
	defer os.Remove(tmpDir) // Clean up the temp directory.
	if IsValidFile(tmpDir) {
		t.Errorf("Expected false for the directory, got true...")
	}
	// -------------------------------------------------------
}
