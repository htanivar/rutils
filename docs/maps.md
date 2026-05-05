# Maps Package Documentation

## Overview

The maps package provides utilities for validating map data structures in Go. It offers simple functions to check whether maps are empty or not empty, helping to prevent issues that arise from nil or empty maps in application logic.

## Key Functions

The maps package provides two primary validation functions that follow the "must" pattern, returning specific errors when validation conditions are not met.

### MustNotBeEmpty
```go
err := maps.MustNotBeEmpty(m)
```

Validates that a map is not empty (contains at least one key-value pair).

**Parameters:**
- `m`: The map to validate

**Returns:**
- `error`: nil if the map is not empty, `ErrEmptyMap` if the map is empty or nil

This function checks if the provided map has any key-value pairs. It returns an error if the map is either nil or empty (has no entries).

**Behavior:**
- Returns `nil` if the map has one or more key-value pairs
- Returns `ErrEmptyMap` if the map is `nil`
- Returns `ErrEmptyMap` if the map is empty (initialized but has no entries)

### MustBeEmpty
```go
err := maps.MustBeEmpty(m)
```

Validates that a map is empty (contains no key-value pairs).

**Parameters:**
- `m`: The map to validate

**Returns:**
- `error`: nil if the map is empty, `ErrNotEmptyMap` if the map is not empty

This function checks if the provided map has no key-value pairs. It returns an error if the map contains any entries.

**Behavior:**
- Returns `nil` if the map is `nil`
- Returns `nil` if the map is empty (initialized but has no entries)
- Returns `ErrNotEmptyMap` if the map has one or more key-value pairs

## Error Types

The package defines specific error types for different validation failure scenarios:

```go
var (
	ErrEmptyMap    = errors.New("map is empty")
	ErrNotEmptyMap = errors.New("map is not empty")
)
```

These error types can be checked using `errors.Is()` for precise error handling.

## Usage Examples

### Ensuring a Configuration Map is Not Empty

```go
// Validate that a configuration map has been properly loaded
if err := maps.MustNotBeEmpty(config); err != nil {
	log.Fatal("Configuration map is empty or nil")
}
// Configuration is valid, safe to use
```

### Checking API Response Data

```go
// After making an API call that returns map data
if err := maps.MustNotBeEmpty(response.Data); err != nil {
	// Handle empty response data
	if errors.Is(err, maps.ErrEmptyMap) {
		log.Print("API returned empty data set")
		// Consider using default values or retrying
	}
} else {
	// Process the non-empty data
	for key, value := range response.Data {
		// ... process each key-value pair ...
	}
}
```

### Initializing a Map with Defaults

```go
// Function to initialize a map with default values if empty
func initConfig(config map[string]interface{}) map[string]interface{} {
	// If config is nil or empty, create with defaults
	if err := maps.MustNotBeEmpty(config); err != nil {
		if errors.Is(err, maps.ErrEmptyMap) {
			config = map[string]interface{}{
				"timeout": 30,
				"retries": 3,
				"debug":   false,
			}
		}
	}
	return config
}
```

### Ensuring a Cache Map is Empty Before Population

```go
// Clear and validate cache before populating with new data
if err := maps.MustBeEmpty(cache); err != nil {
	if errors.Is(err, maps.ErrNotEmptyMap) {
		log.Print("Cache is not empty, clearing...")
		// Clear existing cache
		for k := range cache {
			delete(cache, k)
		}
	} else {
		log.Printf("Unexpected error checking cache: %v\n", err)
	}
}

// Cache is now empty, safe to populate
// ... populate cache with new data ...
```

### Processing Optional Map Parameters

```go
// Function with an optional map parameter
func processData(data map[string]string, options map[string]interface{}) error {
	// Validate required data map
	if err := maps.MustNotBeEmpty(data); err != nil {
		return fmt.Errorf("required data map is empty: %w", err)
	}
	
	// Handle optional options map
	if err := maps.MustNotBeEmpty(options); err != nil {
		if errors.Is(err, maps.ErrEmptyMap) {
			// Options map is empty or nil, use defaults
			log.Print("No options provided, using defaults")
			options = getDefaultOptions()
		} else {
			return fmt.Errorf("error checking options: %w", err)
		}
	}
	
	// Process data with options
	// ... processing logic ...
	
	return nil
}
```

### Comprehensive Error Handling

```go
func validateUserInput(input map[string]string) error {
	// Check for required fields
	requiredFields := []string{"name", "email", "age"}
	
	// First ensure the input map itself is not empty
	if err := maps.MustNotBeEmpty(input); err != nil {
		return fmt.Errorf("input data is missing or empty")
	}
	
	// Check for each required field
	for _, field := range requiredFields {
		if value, exists := input[field]; !exists {
			return fmt.Errorf("required field '%s' is missing", field)
		} else if value == "" {
			return fmt.Errorf("required field '%s' is empty", field)
		}
	}
	
	return nil
}
```

## Integration with Other Packages

The maps package is designed to work with other packages in the rutils library:

### With Path Package

```go
// Validate both path and configuration map
if err := path.MustExists(configPath); err != nil {
	log.Fatal("Configuration file path invalid: ", err)
}

if err := maps.MustNotBeEmpty(configMap); err != nil {
	log.Fatal("Configuration data is empty: ", err)
}
```

### With String Package

```go
// Validate a map of string values
if err := maps.MustNotBeEmpty(userMap); err != nil {
	log.Fatal("User data missing")
}

// Validate specific string fields within the map
if err := string.MustNotBeEmpty(userMap["name"]); err != nil {
	log.Fatal("User name is empty")
}
```

## Testing

The maps package includes comprehensive tests that verify:

- Both validation functions work correctly with nil maps
- Proper handling of initialized but empty maps
- Correct behavior with maps containing various numbers of entries
- Proper error type return values
- Edge cases and boundary conditions

Run tests with:
```bash
make test-maps
```

## Implementation Details

The implementation is straightforward, leveraging Go's built-in `len()` function to determine the number of key-value pairs in a map:

```go
// Simplified implementation of MustNotBeEmpty
func MustNotBeEmpty(m map[string]interface{}) error {
	if len(m) == 0 {
		return ErrEmptyMap
	}
	return nil
}
```

The function works with any map type due to Go's type system, but the signature uses `map[string]interface{}` for maximum flexibility.

## Performance Considerations

- The `len()` operation on maps is O(1) - very fast regardless of map size
- No iteration over map contents is required
- Minimal memory overhead
- Suitable for high-frequency validation in performance-critical code paths

## Security Considerations

- Helps prevent nil pointer dereference panics by validating maps before use
- Ensures data integrity by validating that required map data is present
- Prevents logic errors that could occur when processing empty maps as if they contained data

## Future Improvements

- Add support for validating map size ranges (minimum and maximum number of entries)
- Implement validation for required keys
- Add support for validating map value types
- Support for deep validation of nested maps
- Integration with JSON schema validation
- Add support for map comparison and diffing
- Implement performance benchmarks
- Support for concurrent map validation
- Add validation for map keys
- Support for custom validation functions