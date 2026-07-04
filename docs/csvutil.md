# CSV Utilities (csvutil) Documentation

## Overview

The `csvutil` package provides utility functions to parse and convert tabular CSV files into slices of map records (`[]map[string]string`), using the headers of the CSV file as keys.

---

## API Reference

### Functions

#### CSVToRecords
Reads a CSV file, performs extension and content type validation, and parses its rows into a slice of maps.
```go
func CSVToRecords(path string) ([]map[string]string, error) {
```

**Parameters:**
- `path`: Absolute or relative path to the CSV file.

**Returns:**
- `[]map[string]string`: Slice of maps, where each map corresponds to a row in the CSV file, and the keys are the column headers.
- `error`: Returns `file.ErrInvalidType` if the file doesn't match the `.csv` type, `pathutil.ErrNotExist` if the file is missing, or standard CSV parsing errors.

---

## Code Examples

### Parsing CSV into JSON Records
```go
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/htanivar/rutils/csvutil"
)

func main() {
	records, err := csvutil.CSVToRecords("contacts.csv")
	if err != nil {
		log.Fatalf("failed to parse CSV: %v", err)
	}

	for i, record := range records {
		fmt.Printf("Row %d: Name=%q, Email=%q\n", i+1, record["Name"], record["Email"])
	}

	// Easily marshal to JSON
	jsonData, _ := json.MarshalIndent(records, "", "  ")
	fmt.Println(string(jsonData))
}
```