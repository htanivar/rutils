# Task Package Documentation

## Overview

The task package provides utilities for data conversion tasks, currently focused on converting CSV files to JSON records. This functionality enables easy transformation of comma-separated value data into structured JSON format for further processing or API consumption.

## Key Function

### CSVToRecords
```go
records, err := task.CSVToRecords(path)
```

Converts a CSV file to a slice of map[string]string records, where each map represents a row in the CSV file with column headers as keys.

**Parameters:**
- `path`: Path to the CSV file to convert

**Returns:**
- `[]map[string]string`: Slice of maps, each representing a row with header keys
- `error`: nil if successful, specific error if conversion fails

**Error Conditions:**
- File does not exist
- File is not a CSV file (based on extension and content validation)
- CSV parsing error (malformed CSV content)

The function reads the CSV file, validates it as a CSV file using the file package's validation functions, and then parses it into a structured format where:
- The first row is treated as column headers
- Each subsequent row becomes a map with header names as keys
- Missing values in rows are represented as empty strings

## Return Value Structure

The function returns a slice of `map[string]string` where:
- Each map represents one row from the CSV file
- Map keys are the column headers from the first row
- Map values are the corresponding cell values from that row
- If a cell is empty in the CSV, the map value will be an empty string

For example, given a CSV with:
```
name,age,city
Alice,30,New York
Bob,25,Seattle
```

The result would be:
```go
[]map[string]string{
	{"name": "Alice", "age": "30", "city": "New York"},
	{"name": "Bob", "age": "25", "city": "Seattle"},
}
```

## Usage Examples

### Converting a CSV File to JSON

```go
// Convert CSV file to records
records, err := task.CSVToRecords("/path/to/data.csv")
if err != nil {
	log.Fatal("Failed to convert CSV: ", err)
}

// Convert records to JSON
jsonData, err := json.Marshal(records)
if err != nil {
	log.Fatal("Failed to marshal to JSON: ", err)
}

// Write JSON to file or send via HTTP
err = os.WriteFile("/path/to/data.json", jsonData, 0644)
if err != nil {
	log.Fatal("Failed to write JSON file: ", err)
}
```

### Processing CSV Data Directly

```go
// Convert and process CSV data without JSON marshaling
records, err := task.CSVToRecords("users.csv")
if err != nil {
	log.Fatal("Error reading CSV: ", err)
}

// Process each record
for i, record := range records {
	name := record["name"]
	age := record["age"]
	city := record["city"]
	
	fmt.Printf("User %d: %s is %s years old and lives in %s\n", 
		i+1, name, age, city)
}
```

### Handling CSV with Missing Values

```go
// CSV with missing values in some rows
records, err := task.CSVToRecords("incomplete.csv")
if err != nil {
	log.Fatal("Error: ", err)
}

for _, record := range records {
	// Handle potentially empty values
	name := record["name"]
	if name == "" {
		name = "Unknown" // Default value
	}
	
	age := record["age"]
	if age == "" {
		age = "0" // Default value
	}
	
	city := record["city"] // May be empty
	
	fmt.Printf("%s, %s years old, from %s\n", name, age, city)
}
```

### Validating CSV Before Conversion

```go
// Check file existence and type before conversion
if err := path.MustExists("data.csv"); err != nil {
	log.Fatal("CSV file not found: ", err)
}

if err := file.MustBeType("csv", "data.csv"); err != nil {
	log.Fatal("Invalid CSV file: ", err)
}

// Now safe to convert
records, err := task.CSVToRecords("data.csv")
if err != nil {
	log.Fatal("Conversion failed: ", err)
}

fmt.Printf("Successfully converted %d records\n", len(records))
```

### Error Handling Patterns

```go
func processCSV(filePath string) error {
	records, err := task.CSVToRecords(filePath)
	if err != nil {
		// Handle different error types appropriately
		if os.IsNotExist(err) {
			return fmt.Errorf("CSV file not found: %s", filePath)
		} else if strings.Contains(err.Error(), "invalid type") {
			return fmt.Errorf("file is not a valid CSV: %s", filePath)
		} else {
			return fmt.Errorf("failed to parse CSV %s: %v", filePath, err)
		}
	}
	
	// Process the records
	for i, record := range records {
		// ... process each record ...
	}
	
	return nil
}
```

## Integration with Other Packages

The task package integrates with other packages in the rutils library for comprehensive data processing:

### With File and Path Packages

```go
// Comprehensive validation before CSV processing
func safeCSVProcessing(filePath string) error {
	// Validate path and file type
	if err := path.MustExists(filePath); err != nil {
		return fmt.Errorf("file not found: %w", err)
	}
	
	if err := file.MustBeType("csv", filePath); err != nil {
		return fmt.Errorf("invalid file type: %w", err)
	}
	
	// Convert CSV to records
	records, err := task.CSVToRecords(filePath)
	if err != nil {
		return fmt.Errorf("failed to convert CSV: %w", err)
	}
	
	// Process records
	// ... processing logic ...
	
	return nil
}
```

### With JSON Package

```go
// Convert CSV to JSON and write to file
func csvToJSON(csvPath, jsonPath string) error {
	// Convert CSV to records
	records, err := task.CSVToRecords(csvPath)
	if err != nil {
		return fmt.Errorf("failed to read CSV: %w", err)
	}
	
	// Marshal to JSON
	jsonData, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}
	
	// Write JSON file
	if err := os.WriteFile(jsonPath, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write JSON file: %w", err)
	}
	
	return nil
}
```

## Testing

The task package includes comprehensive tests that verify:

- Successful conversion of valid CSV files
- Proper handling of CSV files with missing values
- Error handling for non-existent files
- Error handling for non-CSV files
- Behavior with empty CSV files (only headers)
- Edge cases like single row CSVs, single column CSVs

The tests use table-driven testing patterns with various CSV content scenarios:
- Valid CSV with multiple rows and columns
- CSV with missing values in some cells
- CSV with only headers (no data rows)
- Empty CSV file
- Non-existent file paths
- Non-CSV file types

Run tests with:
```bash
make test-task
```

## Implementation Details

### CSV Validation

The `CSVToRecords` function leverages the file package's `MustBeType` function to validate that the input file is a CSV file:

```go
// Validation using file package
if err := file.MustBeType("csv", path); err != nil {
	return nil, err
}
```

This ensures the file has a .csv extension and contains valid CSV content.

### CSV Parsing

The function uses Go's standard `encoding/csv` package to parse the CSV content:

```go
// Open and read CSV file
file, err := os.Open(path)
if err != nil {
	return nil, err
}
defer file.Close()

// Create CSV reader
reader := csv.NewReader(file)

// Read all records
records, err := reader.ReadAll()
if err != nil {
	return nil, err
}
```

### Data Transformation

After reading the CSV data, the function transforms it into the desired slice of maps format:

```go
// Extract headers (first row)
if len(records) == 0 {
	return []map[string]string{}, nil
}
headers := records[0]

// Transform remaining rows to maps
result := make([]map[string]string, 0, len(records)-1)
for _, row := range records[1:] {
	record := make(map[string]string)
	for i, value := range row {
		// Use header as key, row value as value
		if i < len(headers) {
			record[headers[i]] = value
		} else {
			// Handle rows with more columns than headers
			record[fmt.Sprintf("column_%d", i)] = value
		}
	}
	result = append(result, record)
}
```

## Performance Considerations

- The function reads the entire CSV file into memory at once
- Memory usage is proportional to the size of the CSV file
- For very large CSV files, consider streaming processing instead
- The transformation process is O(n*m) where n is the number of rows and m is the number of columns
- Includes file validation overhead from the file package

## Security Considerations

- Validates file type before processing to prevent processing of unexpected file types
- Uses standard CSV parsing which handles escaping and quoting correctly
- Be cautious when processing untrusted CSV files as they could contain malicious content
- Consider validating the structure of the CSV data for security-critical applications
- The function does not perform any input sanitization beyond CSV parsing

## Future Improvements

- Add support for other data formats (TSV, Excel, etc.)
- Implement streaming conversion for very large files
- Add support for CSV to other output formats (XML, YAML)
- Implement field type detection and conversion (string to int, float, bool)
- Add support for custom delimiters and separators
- Implement CSV validation against a schema
- Add support for processing CSV from io.Reader instead of just file paths
- Implement progress reporting for large file conversions
- Add support for CSV to database table operations
- Implement column selection and filtering
- Add support for header normalization (case conversion, special character handling)
- Implement encoding detection for non-UTF-8 CSV files