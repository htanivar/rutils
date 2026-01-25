package path

import (
	"os"
	"testing"
)

func TestMustExists(t *testing.T) {
	// Setup: Create temporary file and directory
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	defer tmpfile.Close()

	tmpdir, err := os.MkdirTemp("", "testdir")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)

	// Test cases
	tests := []struct {
		name      string
		path      string
		wantErr   bool
		expectErr error
	}{
		{
			name:      "file exists",
			path:      tmpfile.Name(),
			wantErr:   false,
			expectErr: nil,
		},
		{
			name:      "file does not exist",
			path:      "/tmp/this_file_should_not_exist_1234567890",
			wantErr:   true,
			expectErr: ErrNotExist,
		},
		{
			name:      "directory exists",
			path:      tmpdir,
			wantErr:   false,
			expectErr: nil,
		},
		{
			name:      "directory does not exists",
			path:      tmpdir + "xxx",
			wantErr:   true,
			expectErr: ErrNotExist,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := MustExists(tt.path)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
				if tt.expectErr != nil && err != tt.expectErr {
					t.Errorf("Expected error %v, got %v", tt.expectErr, err)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}
			}
		})
	}
}

func TestMustNotExist(t *testing.T) {
	// Setup: Create temporary file and directory
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	defer tmpfile.Close()

	tmpdir, err := os.MkdirTemp("", "testdir")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)

	// Test cases
	tests := []struct {
		name      string
		path      string
		wantErr   bool
		expectErr error
	}{
		{
			name:      "file exists",
			path:      tmpfile.Name(),
			wantErr:   true,
			expectErr: ErrExists,
		},
		{
			name:      "file does not exist",
			path:      "/tmp/this_file_should_not_exist_1234567890",
			wantErr:   false,
			expectErr: nil,
		},
		{
			name:      "directory exists",
			path:      tmpdir,
			wantErr:   true,
			expectErr: ErrExists,
		},
		{
			name:      "directory does not exist",
			path:      tmpdir + "xxx",
			wantErr:   false,
			expectErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := MustNotExist(tt.path)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
				if tt.expectErr != nil && err != tt.expectErr {
					t.Errorf("Expected error %v, got %v", tt.expectErr, err)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}
			}
		})
	}
}

func TestMustBeEmpty(t *testing.T) {
	// Setup: Create temporary empty directory
	emptyDir, err := os.MkdirTemp("", "emptydir")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(emptyDir)

	// Setup: Create temporary non-empty directory
	nonEmptyDir, err := os.MkdirTemp("", "nonemptydir")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(nonEmptyDir)

	// Add a file to make it non-empty
	tmpfile, err := os.CreateTemp(nonEmptyDir, "file")
	if err != nil {
		t.Fatal(err)
	}
	tmpfile.Close()

	// Setup: Create a regular file (not a directory)
	regularFile, err := os.CreateTemp("", "regularfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(regularFile.Name())
	defer regularFile.Close()

	// Test cases
	tests := []struct {
		name      string
		path      string
		wantErr   bool
		expectErr error
	}{
		{
			name:      "empty directory",
			path:      emptyDir,
			wantErr:   false,
			expectErr: nil,
		},
		{
			name:      "non-empty directory",
			path:      nonEmptyDir,
			wantErr:   true,
			expectErr: ErrNotEmpty,
		},
		{
			name:      "path does not exist",
			path:      "/tmp/this_dir_should_not_exist_1234567890",
			wantErr:   true,
			expectErr: ErrNotExist,
		},
		{
			name:      "path is a file not directory",
			path:      regularFile.Name(),
			wantErr:   true,
			expectErr: ErrNotDirectory,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := MustBeEmpty(tt.path)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
				if tt.expectErr != nil && err != tt.expectErr {
					t.Errorf("Expected error %v, got %v", tt.expectErr, err)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}
			}
		})
	}
}

func TestMustNotBeEmpty(t *testing.T) {
	// Setup: empty directory
	emptyDir, err := os.MkdirTemp("", "emptydir")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(emptyDir)

	// Setup: non-empty directory
	nonEmptyDir, err := os.MkdirTemp("", "nonemptydir")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(nonEmptyDir)

	// Add a file to make it non-empty
	f, err := os.CreateTemp(nonEmptyDir, "file")
	if err != nil {
		t.Fatal(err)
	}
	_ = f.Close()

	// Setup: regular file (not a directory)
	regularFile, err := os.CreateTemp("", "regularfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(regularFile.Name())
	defer regularFile.Close()

	tests := []struct {
		name      string
		path      string
		wantErr   bool
		expectErr error
	}{
		{
			name:      "non-empty directory",
			path:      nonEmptyDir,
			wantErr:   false,
			expectErr: nil,
		},
		{
			name:      "empty directory",
			path:      emptyDir,
			wantErr:   true,
			expectErr: ErrEmpty,
		},
		{
			name:      "path does not exist",
			path:      "/tmp/this_dir_should_not_exist_1234567890",
			wantErr:   true,
			expectErr: ErrNotExist,
		},
		{
			name:      "path is a file not directory",
			path:      regularFile.Name(),
			wantErr:   true,
			expectErr: ErrNotDirectory,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := MustNotBeEmpty(tt.path)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("Expected error, got nil")
				}
				if tt.expectErr != nil && err != tt.expectErr {
					t.Fatalf("Expected error %v, got %v", tt.expectErr, err)
				}
			} else {
				if err != nil {
					t.Fatalf("Expected no error, got %v", err)
				}
			}
		})
	}
}
