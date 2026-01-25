package file

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/htanivar/rutils/path"
)

var (
	ErrInvalidType = errors.New("file type does not match expected extension")
)

// MustBeType validates that the file at filePath actually matches the expected extension type
// Supported extensions: .csv, .json, .xml, .pdf, .zip, .txt
func MustBeType(ext, filePath string) error {
	// Normalize extension (ensure it starts with .)
	if !strings.HasPrefix(ext, ".") {
		ext = "." + ext
	}
	ext = strings.ToLower(ext)

	// Check if file exists
	if err := path.MustExists(filePath); err != nil {
		return err
	}

	// Open file for reading
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	// Read first few KB for magic number/signature detection
	const sampleSize = 8192
	buf := make([]byte, sampleSize)
	n, _ := f.Read(buf)
	buf = buf[:n]

	if n == 0 {
		return ErrInvalidType
	}

	// Validate based on extension
	switch ext {
	case ".csv":
		return validateCSV(filePath)
	case ".json":
		return validateJSON(buf)
	case ".xml":
		return validateXML(buf)
	case ".pdf":
		return validatePDF(buf)
	case ".zip":
		return validateZIP(buf)
	case ".txt":
		return validateText(buf)
	default:
		// For unknown extensions, just check file extension matches
		actualExt := strings.ToLower(filepath.Ext(filePath))
		if actualExt != ext {
			return ErrInvalidType
		}
		return nil
	}
}

func validateCSV(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	// Read sample for quick checks
	const sampleSize = 64 * 1024
	buf := make([]byte, sampleSize)
	n, _ := f.Read(buf)
	buf = buf[:n]

	if n == 0 {
		return ErrInvalidType
	}

	// Reject binary content (NUL bytes)
	if bytes.IndexByte(buf, 0x00) != -1 {
		return ErrInvalidType
	}

	// Reject other known formats by magic numbers/signatures
	trim := bytes.TrimSpace(buf)
	if len(trim) == 0 {
		return ErrInvalidType
	}

	if bytes.HasPrefix(trim, []byte("%PDF-")) {
		return ErrInvalidType
	}
	if bytes.HasPrefix(trim, []byte("PK\x03\x04")) {
		return ErrInvalidType
	}
	if trim[0] == '{' || trim[0] == '[' {
		return ErrInvalidType
	}
	if trim[0] == '<' {
		return ErrInvalidType
	}

	// Reset to beginning for full CSV parsing
	if _, err := f.Seek(0, 0); err != nil {
		return err
	}

	r := csv.NewReader(bufio.NewReader(f))
	r.FieldsPerRecord = -1 // Allow variable fields initially
	r.LazyQuotes = true
	r.TrimLeadingSpace = true

	// Read all records to ensure entire file is parseable
	var records [][]string
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			// CSV parsing error means it's not valid CSV
			return ErrInvalidType
		}

		// Skip completely empty lines
		if len(rec) == 1 && strings.TrimSpace(rec[0]) == "" {
			continue
		}

		records = append(records, rec)
	}

	// Must have at least one non-empty record
	if len(records) == 0 {
		return ErrInvalidType
	}

	// Validate column consistency
	// All records should have the same number of fields
	expectedCols := len(records[0])
	if expectedCols == 0 {
		return ErrInvalidType
	}

	for _, rec := range records {
		if len(rec) != expectedCols {
			// Allow some tolerance for trailing empty fields
			// but reject if difference is significant
			diff := expectedCols - len(rec)
			if diff < 0 {
				diff = -diff
			}
			if diff > 0 {
				return ErrInvalidType
			}
		}
	}

	return nil
}

func validateJSON(buf []byte) error {
	trim := bytes.TrimSpace(buf)
	if len(trim) == 0 {
		return ErrInvalidType
	}
	// JSON must start with { or [
	if trim[0] != '{' && trim[0] != '[' {
		return ErrInvalidType
	}
	return nil
}

func validateXML(buf []byte) error {
	trim := bytes.TrimSpace(buf)
	if len(trim) == 0 {
		return ErrInvalidType
	}
	// XML must start with <
	if trim[0] != '<' {
		return ErrInvalidType
	}
	return nil
}

func validatePDF(buf []byte) error {
	// PDF magic number: %PDF-
	if !bytes.HasPrefix(buf, []byte("%PDF-")) {
		return ErrInvalidType
	}
	return nil
}

func validateZIP(buf []byte) error {
	// ZIP magic number: PK\x03\x04
	if len(buf) < 4 {
		return ErrInvalidType
	}
	if !bytes.HasPrefix(buf, []byte("PK\x03\x04")) {
		return ErrInvalidType
	}
	return nil
}

func validateText(buf []byte) error {
	// Text files should not have NUL bytes
	if bytes.IndexByte(buf, 0x00) != -1 {
		return ErrInvalidType
	}

	// Reject known binary/structured formats
	trim := bytes.TrimSpace(buf)
	if len(trim) == 0 {
		return nil // Empty is OK for text
	}

	// Reject PDF
	if bytes.HasPrefix(trim, []byte("%PDF-")) {
		return ErrInvalidType
	}

	// Reject ZIP
	if bytes.HasPrefix(trim, []byte("PK\x03\x04")) {
		return ErrInvalidType
	}

	// Reject JSON
	if trim[0] == '{' || trim[0] == '[' {
		return ErrInvalidType
	}

	// Reject XML
	if trim[0] == '<' {
		return ErrInvalidType
	}

	return nil
}
