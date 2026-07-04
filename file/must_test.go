package file_test

import (
	"archive/zip"
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/htanivar/rutils/file"
	"github.com/htanivar/rutils/pathutil"
)

func createValidZIP(t *testing.T) []byte {
	t.Helper()
	var buf bytes.Buffer
	w := zip.NewWriter(&buf)
	f, err := w.Create("test.txt")
	if err != nil {
		t.Fatal(err)
	}
	_, _ = f.Write([]byte("test content"))
	_ = w.Close()
	return buf.Bytes()
}

func TestValidateType(t *testing.T) {
	tmp := t.TempDir()

	write := func(name string, b []byte) string {
		t.Helper()
		p := filepath.Join(tmp, name)
		if err := os.WriteFile(p, b, 0o644); err != nil {
			t.Fatalf("write %s: %v", name, err)
		}
		return p
	}

	// Good samples
	csvPath := write("ok.csv", []byte("a,b\n1,2\n3,4\n"))
	jsonPath := write("ok.json", []byte(`{"a":1}`))
	xmlPath := write("ok.xml", []byte(`<?xml version="1.0" encoding="UTF-8"?><root></root>`))
	pdfPath := write("ok.pdf", []byte("%PDF-1.4\n%...\n"))
	zipPath := write("ok.zip", createValidZIP(t))
	txtPath := write("ok.txt", []byte("hello\nworld\n"))

	// Bad samples / mismatches
	jsonButCsvName := write("mismatch.csv", []byte(`{"a":1}`))
	pdfButTxtName := write("mismatch.txt", []byte("%PDF-1.7\n..."))
	emptyFile := write("empty.csv", []byte(""))

	// Non-existent path
	nonExist := filepath.Join(tmp, "nope.csv")

	tests := []struct {
		name      string
		ext       string
		filePath  string
		wantErr   bool
		expectErr error
	}{
		// Happy paths
		{name: "csv ok", ext: ".csv", filePath: csvPath, wantErr: false},
		{name: "csv ok (no dot ext)", ext: "csv", filePath: csvPath, wantErr: false},
		{name: "json ok", ext: ".json", filePath: jsonPath, wantErr: false},
		{name: "xml ok", ext: ".xml", filePath: xmlPath, wantErr: false},
		{name: "pdf ok", ext: ".pdf", filePath: pdfPath, wantErr: false},
		{name: "zip ok", ext: ".zip", filePath: zipPath, wantErr: false},
		{name: "txt ok", ext: ".txt", filePath: txtPath, wantErr: false},

		// Invalid type
		{name: "csv ext but json content", ext: ".csv", filePath: jsonButCsvName, wantErr: true, expectErr: file.ErrInvalidType},
		{name: "txt ext but pdf content", ext: ".txt", filePath: pdfButTxtName, wantErr: true, expectErr: file.ErrInvalidType},
		{name: "empty file is invalid", ext: ".csv", filePath: emptyFile, wantErr: true, expectErr: file.ErrInvalidType},

		// Unknown ext => fallback to extension match only
		{name: "unknown ext matches filename", ext: ".foo", filePath: write("x.foo", []byte("whatever")), wantErr: false},
		{name: "unknown ext mismatches filename", ext: ".foo", filePath: write("x.bar", []byte("whatever")), wantErr: true, expectErr: file.ErrInvalidType},

		// Missing file => should return Exists error (ErrNotExist)
		{name: "file does not exist", ext: ".csv", filePath: nonExist, wantErr: true, expectErr: pathutil.ErrNotExist},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := file.ValidateType(tt.ext, tt.filePath)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				if tt.expectErr != nil && !errors.Is(err, tt.expectErr) {
					t.Fatalf("expected error %v, got %v", tt.expectErr, err)
				}
			} else {
				if err != nil {
					t.Fatalf("expected no error, got %v", err)
				}
			}
		})
	}
}

func TestMustBeType(t *testing.T) {
	tmp := t.TempDir()
	jsonPath := filepath.Join(tmp, "ok.json")
	if err := os.WriteFile(jsonPath, []byte(`{"a":1}`), 0o644); err != nil {
		t.Fatal(err)
	}

	t.Run("should not panic when type matches", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("MustBeType panicked unexpectedly: %v", r)
			}
		}()
		file.MustBeType(".json", jsonPath)
	})

	t.Run("should panic when type does not match", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Error("expected MustBeType to panic")
			}
			err, ok := r.(error)
			if !ok || !errors.Is(err, file.ErrInvalidType) {
				t.Errorf("expected panic with ErrInvalidType, got %v", r)
			}
		}()
		file.MustBeType(".csv", jsonPath)
	})
}
