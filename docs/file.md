# File Package Documentation

## Overview

The file package provides utilities for validating file types based on both file extension and content. It ensures that a file is of the expected type by examining both its name and actual content, preventing issues that arise from incorrect file extensions.

## Key Functions

### MustBeType
```go
err := file.MustBeType(ext, path)
```

Validates that a file at the specified path is of the specified type by checking both the file extension and the actual content.

**Parameters:**
- `ext`: Expected file extension (can be with or without leading dot, e.g., ".csv" or "csv")
- `path`: Path to the file to validate

**Returns:**
- `error`: nil if the file is valid, specific error type if validation fails

**Error Types:**
- `ErrEmptyFile`: File exists but is empty
- `ErrFileDoesNotExist`: File does not exist at the specified path
- `ErrInvalidType`: File exists but is not of the expected type

This function is the primary validation function in the package and delegates to specific type validation functions based on the extension.

## Supported File Types

The file package supports validation for the following file types:

### CSV Files (.csv)
Validates that a file is a properly formatted CSV file.

### JSON Files (.json)
Validates that a file contains valid JSON content.

### XML Files (.xml)
Validates that a file contains valid XML content.

### PDF Files (.pdf)
Validates that a file contains a valid PDF document by checking for the PDF header.

### ZIP Files (.zip)
Validates that a file is a valid ZIP archive.

### Text Files (.txt)
Validates that a file is a text file.

## Usage Examples

### Validating a CSV File

```go
// Validate that a file is a CSV file
err := file.MustBeType(".csv", "/path/to/data.csv")
if err != nil {
	// Handle validation error
	switch {
	case errors.Is(err, file.ErrFileDoesNotExist):
		log.Fatal("File does not exist")
	case errors.Is(err, file.ErrEmptyFile):
		log.Fatal("File is empty")
	case errors.Is(err, file.ErrInvalidType):
		log.Fatal("File is not a valid CSV")
	default:
		log.Fatal("Unexpected error:", err)
	}
}
// File is valid CSV
```

### Validating a JSON File

```go
// Validate that a file is a JSON file
err := file.MustBeType("json", "/path/to/config.json")
if err != nil {
	log.Fatal("Invalid JSON file:", err)
}
```

### Validating Multiple Files

```go
// Validate multiple files of different types
files := map[string]string{
	".csv": "/path/to/data.csv",
	"json": "/path/to/config.json",
	".xml": "/path/to/data.xml",
}

for ext, path := range files {
	if err := file.MustBeType(ext, path); err != nil {
		log.Printf("Validation failed for %s: %v\n", path, err)
		// Handle error as needed
	} else {
		log.Printf("File %s is valid\n", path)
	}
}
```

### Comprehensive Error Handling

```go
func processFile(filePath string) error {
	// First validate it's a CSV file
	if err := file.MustBeType("csv", filePath); err != nil {
		// Handle different error types appropriately
		var msg string
		switch {
		case errors.Is(err, file.ErrFileDoesNotExist):
			msg = "Input file not found"
		case errors.Is(err, file.ErrEmptyFile):
			msg = "Input file is empty"
		case errors.Is(err, file.ErrInvalidType):
			msg = "Input file is not a valid CSV file"
		default:
			msg = fmt.Sprintf("Unexpected file error: %v", err)
		}
		return fmt.Errorf("%s: %s", msg, filePath)
	}
	
	// File is valid, proceed with processing
	// ... processing logic here ...
	
	return nil
}
```

## Implementation Details

### Extension Normalization

The package normalizes file extensions by ensuring they have a leading dot:

```go
// Both ".csv" and "csv" are treated the same
file.MustBeType("csv", "path.csv")    // valid
file.MustBeType(".csv", "path.csv")   // also valid
```

### Content-Based Validation

For more reliable validation, the package doesn't just rely on file extensions but also examines the actual file content:

- **CSV**: Attempts to parse the content as CSV
- **JSON**: Attempts to unmarshal the content as JSON
- **XML**: Attempts to parse the content as XML
- **PDF**: Checks for the PDF header ("%PDF-")
- **ZIP**: Checks for the ZIP header signature
- **TXT**: Performs basic text validation

This dual-check approach prevents security issues and processing errors that could occur if an attacker renamed a malicious file with a different extension.

## Testing

The file package includes comprehensive tests that verify:

- Successful validation of correct file types
- Proper error handling for incorrect file types
- Empty file detection
- Non-existent file handling
- Extension normalization
- Edge cases and boundary conditions

Tests use temporary files with various combinations of correct and incorrect extensions and content types.

Run tests with:
```bash
make test-file
```

## Integration with Path Package

The file package works closely with the path package for comprehensive file validation. While the file package validates file types and content, the path package validates file and directory existence and other path-related conditions.

Common usage pattern:

```go
// First check if file exists
if err := path.MustExists(filePath); err != nil {
	log.Fatal("File does not exist:", err)
}

// Then check if it's the right type
if err := file.MustBeType("csv", filePath); err != nil {
	log.Fatal("Invalid file type:", err)
}

// Now it's safe to process
```

## Error Types

The package defines specific error types for different validation failure scenarios:

```go
// Error types defined in the package
var (
	ErrEmptyFile     = errors.New("file is empty")
	ErrFileDoesNotExist = os.ErrNotExist
	ErrInvalidType     = errors.New("invalid file type")
)
```

These error types can be checked using `errors.Is()` for precise error handling.

## Performance Considerations

- Reading file content for validation has some overhead
- For large files, this validation ensures the file is of the expected type but may take longer
- The package is optimized to fail quickly when a file doesn't have the expected extension
- Content validation only occurs after passing the extension check

## Security Considerations

- Prevents file type spoofing by verifying content, not just extension
- Helps prevent injection attacks by ensuring files are of expected type
- Validates file existence before processing
- Properly handles empty files to prevent processing of incomplete data

## Future Improvements

- Add support for additional file types (e.g., YAML, TOML)
- Implement streaming validation for very large files
- Add size limits and performance optimizations
- Support for custom validation functions
- Integration with MIME type detection
- Enhanced error messages with more context
- Support for directory content validation
- Add support for compressed formats (gz, bz2)