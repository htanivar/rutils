package csvutil

import (
	"os"
	"path/filepath"
	"testing"
)

func writeTempFile(t *testing.T, name, content string) string {
	t.Helper()

	dir := t.TempDir()
	path := filepath.Join(dir, name)

	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	return path
}

func TestCSVToRecords(t *testing.T) {
	tests := []struct {
		name        string
		filename    string
		content     string
		wantLen     int
		wantRecord  map[string]string
		expectError bool
	}{
		{
			name:     "valid csv",
			filename: "test.csv",
			content: `name,age,city
Alice,30,Delhi`,
			wantLen: 1,
			wantRecord: map[string]string{
				"name": "Alice",
				"age":  "30",
				"city": "Delhi",
			},
		},
		{
			name:     "missing column value",
			filename: "test.csv",
			content: `name,age,city
Alice,30`,
			wantLen: 1,
			wantRecord: map[string]string{
				"name": "Alice",
				"age":  "30",
				"city": "",
			},
		},
		{
			name:     "only header",
			filename: "test.csv",
			content:  `name,age,city`,
			wantLen:  0,
		},
		{
			name:        "non csv file (txt ext)",
			filename:    "test.txt",
			content:     `dummy`,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := writeTempFile(t, tt.filename, tt.content)

			records, err := CSVToRecords(path)

			if tt.expectError {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if len(records) != tt.wantLen {
				t.Fatalf("expected %d records, got %d", tt.wantLen, len(records))
			}

			if tt.wantLen > 0 {
				for k, v := range tt.wantRecord {
					if records[0][k] != v {
						t.Fatalf("key %q: expected %q, got %q", k, v, records[0][k])
					}
				}
			}
		})
	}
}
