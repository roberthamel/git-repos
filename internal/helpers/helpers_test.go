package helpers

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileExists(t *testing.T) {
	tempFile, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	if !FileExists(tempFile.Name()) {
		t.Errorf("FileExists(%s) = false; want true", tempFile.Name())
	}

	if FileExists("nonexistentfile.xyz") {
		t.Error(`FileExists("nonexistentfile.xyz") = true; want false`)
	}
}

func TestIsFileEmpty(t *testing.T) {
	tempFile, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	isEmpty, err := IsFileEmpty(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}
	if !isEmpty {
		t.Errorf("IsFileEmpty(%s) = false; want true", tempFile.Name())
	}

	if _, err := IsFileEmpty("nonexistentfile.xyz"); err == nil {
		t.Error(`IsFileEmpty("nonexistentfile.xyz") = nil; want error`)
	}
}

func TestWriteFile(t *testing.T) {
	tempFile, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	err = WriteFile(tempFile.Name(), "test content")
	if err != nil {
		t.Fatal(err)
	}

	content, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}
	if string(content) != "test content" {
		t.Errorf("WriteFile(%s) = %s; want 'test content'", tempFile.Name(), string(content))
	}
}

func TestTruncateFile(t *testing.T) {
	tempFile, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	err = WriteFile(tempFile.Name(), "test content")
	if err != nil {
		t.Fatal(err)
	}

	err = TruncateFile(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	content, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}
	if string(content) != "" {
		t.Errorf("TruncateFile(%s) = %s; want ''", tempFile.Name(), string(content))
	}
}

func TestReadLastLineOfFile(t *testing.T) {
	// Create a temporary file
	tempFile, err := os.CreateTemp("", "test")
	if err != nil {
			t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	// Write some lines to the file
	lines := []string{"First line", "Second line", "Last line"}
	for _, line := range lines {
			if _, err := tempFile.WriteString(line + "\n"); err != nil {
					t.Fatal(err)
			}
	}
	tempFile.Close()

	// Test the function
	lastLine, err := ReadLastLineOfFile(tempFile.Name())
	if err != nil {
			t.Fatal(err)
	}

	// Check the result
	assert.Equal(t, "Last line", lastLine)
}
