# File Utilities (file) Documentation

## Overview

The `file` package provides utilities for validating file types based on both extension and actual content signatures/magic bytes. This ensures that files are of the expected format (CSV, JSON, XML, PDF, ZIP, TXT) and prevents errors due to mismatched or spoofed file extensions.

---

## API Reference

### Functions

#### ValidateType
Validates that the file exists, has the specified extension, and matches the expected content structure (magic bytes or format parsers).
```go
func ValidateType(ext, filePath string) error
```

#### MustBeType
Validates that the file matches the expected extension and content signature, panicking if it does not.
```go
func MustBeType(ext, filePath string)
```

---

## Error Types
```go
var ErrInvalidType = errors.New("file type does not match expected extension")
```
*Note: If the file does not exist, `ValidateType` returns `pathutil.ErrNotExist`.*

---

## Supported Formats & Validation Rules

- **`.csv`**: Checks for non-binary content (no NUL bytes), rejects headers of other formats (PDF, ZIP, JSON, XML), and successfully parses the content as comma-separated values using `encoding/csv`.
- **`.json`**: Verifies the file starts with `{` or `[` and that its first token is valid according to `encoding/json`.
- **`.xml`**: Verifies the file starts with `<` and that its first token can be parsed by `encoding/xml`.
- **`.pdf`**: Checks for the `%PDF-` signature prefix.
- **`.zip`**: Attempts to open and read the file's central directory headers using Go's standard `archive/zip` library.
- **`.txt`**: Verifies there are no NUL (binary) bytes and that the content does not match any of the structured formats above (PDF, ZIP, JSON, XML).

---

## Code Examples

### File Type Validation
```go
package main

import (
	"log"

	"github.com/htanivar/rutils/file"
)

func main() {
	// Standard validation
	if err := file.ValidateType("json", "config.json"); err != nil {
		log.Fatalf("invalid configuration format: %v", err)
	}

	// Panicking validation
	file.MustBeType(".zip", "archive.zip")
}
```