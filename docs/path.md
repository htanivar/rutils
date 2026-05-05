# Path Package Documentation

## Overview

The path package provides utilities for validating file and directory paths. It offers a set of "must" functions that check various conditions about paths and return specific errors when conditions are not met. These functions are designed to be used in validation workflows to ensure paths meet expected criteria before file operations are performed.

## Key Functions

The path package provides several validation functions that follow the pattern `MustX` where X is the condition being checked. All functions return an error if the condition is not met, or nil if the condition is satisfied.

### MustExists
```go
err := path.MustExists(p)
```

Validates that a path exists (file or directory).

**Parameters:**
- `p`: Path to check

**Returns:**
- `error`: nil if path exists, `ErrNotExist` if path does not exist

This is the most basic check to ensure a file or directory exists before attempting operations on it.

### MustNotExist
```go
err := path.MustNotExist(p)
```

Validates that a path does not exist.

**Parameters:**
- `p`: Path to check

**Returns:**
- `error`: nil if path does not exist, `ErrExist` if path exists

Useful when creating new files or directories to ensure you don't overwrite existing data.

### MustBeEmpty
```go
err := path.MustBeEmpty(p)
```

Validates that a directory is empty. If the path points to a file, an error is returned.

**Parameters:**
- `p`: Path to directory to check

**Returns:**
- `error`: nil if directory exists and is empty, specific error if not

**Error Types:**
- `ErrNotExist`: Path does not exist
- `ErrNotADirectory`: Path exists but is not a directory
- `ErrNotEmpty`: Directory exists but is not empty

### MustNotBeEmpty
```go
err := path.MustNotBeEmpty(p)
```

Validates that a directory is not empty. If the path points to a file, an error is returned.

**Parameters:**
- `p`: Path to directory to check

**Returns:**
- `error`: nil if directory exists and is not empty, specific error if not

**Error Types:**
- `ErrNotExist`: Path does not exist
- `ErrNotADirectory`: Path exists but is not a directory
- `ErrEmpty`: Directory exists but is empty

### MustBeFile
```go
err := path.MustBeFile(p)
```

Validates that a path exists and is a regular file (not a directory).

**Parameters:**
- `p`: Path to check

**Returns:**
- `error`: nil if path is a file, specific error if not

**Error Types:**
- `ErrNotExist`: Path does not exist
- `ErrNotAFile`: Path exists but is not a regular file (e.g., it's a directory)

### MustBeDir
```go
err := path.MustBeDir(p)
```

Validates that a path exists and is a directory.

**Parameters:**
- `p`: Path to check

**Returns:**
- `error`: nil if path is a directory, specific error if not

**Error Types:**
- `ErrNotExist`: Path does not exist
- `ErrNotADirectory`: Path exists but is not a directory

## Error Types

The package defines specific error types for different validation failure scenarios:

```go
var (
	ErrNotExist       = errors.New("path does not exist")
	ErrExist          = errors.New("path exists")
	ErrNotADirectory  = errors.New("path is not a directory")
	ErrNotAFile       = errors.New("path is not a file")
	ErrEmpty          = errors.New("directory is empty")
	ErrNotEmpty       = errors.New("directory is not empty")
)
```

These error types can be checked using `errors.Is()` for precise error handling.

## Usage Examples

### Checking if a File Exists

```go
// Check if a configuration file exists
if err := path.MustExists("/etc/myapp/config.json"); err != nil {
	log.Fatalf("Configuration file missing: %v", err)
}
// File exists, safe to read
```

### Creating a New Directory

```go
// Ensure directory doesn't already exist before creating
if err := path.MustNotExist("/tmp/mydata"); err != nil {
	if errors.Is(err, path.ErrExist) {
		log.Fatal("Directory already exists, refusing to overwrite")
	} else {
		log.Fatal("Error checking directory: ", err)
	}
}

// Directory doesn't exist, safe to create
if err := os.Mkdir("/tmp/mydata", 0755); err != nil {
	log.Fatal("Failed to create directory: ", err)
}
```

### Processing a Directory of Files

```go
// Ensure directory exists and is not empty before processing
if err := path.MustExists(inputDir); err != nil {
	log.Fatal("Input directory does not exist: ", err)
}

if err := path.MustNotBeEmpty(inputDir); err != nil {
	switch {
	case errors.Is(err, path.ErrEmpty):
		log.Print("Input directory is empty, nothing to process")
	case errors.Is(err, path.ErrNotADirectory):
		log.Fatal("Input path is not a directory")
	default:
		log.Fatal("Error checking input directory: ", err)
	}
	return // Nothing to process
}

// Directory exists and has files, safe to process
files, err := os.ReadDir(inputDir)
// ... process files ...
```

### Ensuring a Directory is Empty

```go
// Clear output directory by ensuring it exists and is empty
if err := path.MustExists(outputDir); err != nil {
	// Directory doesn't exist, create it
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatal("Failed to create output directory: ", err)
	}
} else {
	// Directory exists, ensure it's empty
	if err := path.MustBeEmpty(outputDir); err != nil {
		if errors.Is(err, path.ErrNotEmpty) {
			log.Print("Output directory is not empty, clearing contents...")
			// Clear directory contents
			if err := os.RemoveAll(outputDir); err != nil {
				log.Fatal("Failed to clear output directory: ", err)
			}
			// Recreate the directory
			if err := os.Mkdir(outputDir, 0755); err != nil {
				log.Fatal("Failed to recreate output directory: ", err)
			}
		} else {
			log.Fatal("Error checking output directory: ", err)
		}
	}
}
```

### Checking File vs Directory

```go
// Determine if a path is a file or directory
if err := path.MustBeFile(path); err == nil {
	// Path is a file
	fmt.Printf("%s is a file\n", path)
} else if errors.Is(err, path.ErrNotAFile) {
	// Path exists but is not a file (likely a directory)
	if err := path.MustBeDir(path); err == nil {
		fmt.Printf("%s is a directory\n", path)
	} else {
		fmt.Printf("%s is neither a file nor directory (special file?)\n", path)
	}
} else if errors.Is(err, path.ErrNotExist) {
	fmt.Printf("%s does not exist\n", path)
} else {
	log.Printf("Unexpected error checking path %s: %v\n", path, err)
}
```

## Testing

The path package includes comprehensive tests that verify:

- Each validation function works correctly
- Proper error handling for edge cases
- Behavior with temporary files and directories
- Error type specificity

Tests use Go's `testing` package with parallel execution and temporary directories to ensure isolation.

Run tests with:
```bash
make test-path
```

## Integration with Other Packages

The path package is designed to work with other packages in the rutils library:

### With File Package

```go
// First check if file exists, then check if it's the right type
if err := path.MustExists(filePath); err != nil {
	log.Fatal("File does not exist: ", err)
}

if err := file.MustBeType("csv", filePath); err != nil {
	log.Fatal("Invalid file type: ", err)
}
```

### With Maps Package

```go
// Ensure configuration directory exists and configuration map is not empty
if err := path.MustExists(configDir); err != nil {
	log.Fatal("Configuration directory missing: ", err)
}

if err := maps.MustNotBeEmpty(configMap); err != nil {
	log.Fatal("Configuration is empty")
}
```

## Error Handling Patterns

The path package follows Go's error handling patterns, allowing for precise error checking:

```go
err := path.MustExists(p)
if err != nil {
	if errors.Is(err, path.ErrNotExist) {
		// Handle missing path specifically
	} else {
		// Handle other errors
	}
}
```

This allows callers to respond appropriately to different types of validation failures.

## Performance Considerations

- All functions use `os.Stat()` or `os.Lstat()` for path information
- Operations are generally fast as they don't read file contents
- For directories, `MustBeEmpty` and `MustNotBeEmpty` read the directory contents to check if it's empty, which has O(n) complexity where n is the number of entries

## Security Considerations

- Validates paths before file operations, preventing issues with missing files
- Helps prevent directory traversal issues by validating expected path types
- Ensures operations are performed on the expected type of path (file vs directory)

## Future Improvements

- Add support for symbolic link validation
- Implement path permission checks
- Add file size validation functions
- Support for path pattern matching
- Integration with file system events
- Add recursive directory validation
- Support for path normalization and cleanup
- Enhanced error messages with more context about the path state