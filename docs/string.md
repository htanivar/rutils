# String Package Documentation

## Overview

The string package provides utilities for string validation and manipulation in Go. It includes functions for validating that strings are not empty and for reversing strings.

## Key Components

### String Validation

The package provides a validation function to ensure strings meet certain criteria.

#### MustNotBeEmpty
```go
err := string.MustNotBeEmpty(s)
```

Validates that a string is not empty. This function checks if the string has any characters, returning an error if the string is empty.

**Parameters:**
- `s`: The string to validate

**Returns:**
- `error`: nil if the string is not empty, `ErrEmptyString` if the string is empty

**Behavior:**
- Returns `nil` if the string has one or more characters
- Returns `ErrEmptyString` if the string is empty (zero length)
- Note: Strings containing only whitespace are considered non-empty

This function is useful for validating user input, configuration values, and other string data where an empty string would be invalid or indicate a problem.

### String Manipulation

The package provides a function for reversing the characters in a string.

#### Reverse
```go
reversed := string.Reverse(s)
```

Reverses the characters in a string.

**Parameters:**
- `s`: The string to reverse

**Returns:**
- `string`: A new string with the characters in reverse order

**Behavior:**
- Returns a new string with characters in reverse order
- Returns empty string if input is empty
- Handles Unicode characters correctly, including multi-byte UTF-8 sequences
- Preserves all characters, including whitespace and special characters

This function correctly handles Unicode by treating the string as a sequence of runes rather than bytes, ensuring that multi-byte characters (like emojis or non-ASCII characters) are properly reversed.

## Error Types

The package defines a specific error type for validation failures:

```go
var ErrEmptyString = errors.New("string is empty")
```

This error type can be checked using `errors.Is()` for precise error handling.

## Usage Examples

### Validating User Input

```go
// Validate required user input fields
func validateUser(name, email, phone string) error {
	// Check that required fields are not empty
	if err := string.MustNotBeEmpty(name); err != nil {
		return fmt.Errorf("name is required: %w", err)
	}
	
	if err := string.MustNotBeEmpty(email); err != nil {
		return fmt.Errorf("email is required: %w", err)
	}
	
	if err := string.MustNotBeEmpty(phone); err != nil {
		return fmt.Errorf("phone is required: %w", err)
	}
	
	return nil
}
```

### Processing Configuration Values

```go
// Validate configuration values before use
func loadConfig(config map[string]string) error {
	requiredKeys := []string{"api_key", "base_url", "timeout"}
	
	for _, key := range requiredKeys {
		value, exists := config[key]
		if !exists {
			return fmt.Errorf("missing required config key: %s", key)
		}
		
		// Validate that the value is not empty
		if err := string.MustNotBeEmpty(value); err != nil {
			return fmt.Errorf("config value for '%s' is empty", key)
		}
	}
	
	return nil
}
```

### Reversing Strings for Processing

```go
// Function to check if a string is a palindrome
func isPalindrome(s string) bool {
	// Convert to lowercase and remove spaces for comparison
	processed := strings.ToLower(strings.ReplaceAll(s, " ", ""))
	
	// A string is a palindrome if it reads the same forwards and backwards
	return processed == string.Reverse(processed)
}

// Usage
fmt.Println(isPalindrome("racecar")) // true
fmt.Println(isPalindrome("hello"))    // false
fmt.Println(isPalindrome("A man a plan a canal Panama")) // true
```

### Reversing Strings with Unicode

```go
// The Reverse function correctly handles Unicode characters
func demonstrateUnicodeReversal() {
	// Regular ASCII string
	ascii := "hello"
	fmt.Printf("'%s' reversed is '%s'\n", ascii, string.Reverse(ascii))
	// Output: 'hello' reversed is 'olleh'
	
	// String with Unicode characters
	unicode := "Hello, 世界!"
	fmt.Printf("'%s' reversed is '%s'\n", unicode, string.Reverse(unicode))
	// Output: 'Hello, 世界!' reversed is '!界世 ,olleH'
	
	// String with emojis
	emoji := "Hello 👋🌍!"
	fmt.Printf("'%s' reversed is '%s'\n", emoji, string.Reverse(emoji))
	// Output: 'Hello 👋🌍!' reversed is '!🌍👋 olleH'
}
```

### Processing Strings in a Slice

```go
// Process a slice of strings, validating and reversing each one
func processStrings(stringsList []string) ([]string, error) {
	if err := maps.MustNotBeEmpty(stringsList); err != nil {
		return nil, fmt.Errorf("string list is empty")
	}
	
	result := make([]string, 0, len(stringsList))
	
	for i, s := range stringsList {
		// Validate that the string is not empty
		if err := string.MustNotBeEmpty(s); err != nil {
			return nil, fmt.Errorf("string at index %d is empty", i)
		}
		
		// Reverse the string and add to result
		reversed := string.Reverse(s)
		result = append(result, reversed)
	}
	
	return result, nil
}
```

### Chaining Validation and Manipulation

```go
// Function that validates and processes a string in one operation
func validateAndReverse(input string) (string, error) {
	// First validate the input
	if err := string.MustNotBeEmpty(input); err != nil {
		return "", fmt.Errorf("invalid input: %w", err)
	}
	
	// Then process it
	result := string.Reverse(input)
	
	return result, nil
}

// Usage
if reversed, err := validateAndReverse("hello"); err != nil {
	log.Fatal(err)
} else {
	fmt.Printf("Reversed: %s\n", reversed)
}
```

## Testing

The string package includes comprehensive tests that verify:

- `MustNotBeEmpty` function works correctly with empty strings
- `MustNotBeEmpty` function works correctly with non-empty strings (including whitespace)
- `Reverse` function correctly reverses ASCII strings
- `Reverse` function correctly handles Unicode strings and multi-byte characters
- `Reverse` function handles edge cases like empty strings and single characters
- Proper error type return values

The tests include various string inputs:
- Empty string
- Single character
- Basic ASCII strings
- Strings with spaces and special characters
- Unicode strings with non-ASCII characters
- Strings with emojis
- Palindromes

Run tests with:
```bash
make test-string
```

## Implementation Details

### MustNotBeEmpty Implementation

The `MustNotBeEmpty` function is straightforward:

```go
func MustNotBeEmpty(s string) error {
	if len(s) == 0 {
		return ErrEmptyString
	}
	return nil
}
```

It uses Go's built-in `len()` function to check the string length, which returns the number of bytes in the string. For UTF-8 encoded strings, this means the length in bytes rather than runes, but for the purpose of checking if a string is empty, this is sufficient.

### Reverse Implementation

The `Reverse` function handles Unicode correctly by converting the string to a slice of runes:

```go
func Reverse(s string) string {
	// Convert string to rune slice to handle Unicode correctly
	runes := []rune(s)
	
	// Reverse the rune slice
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	
	// Convert back to string
	return string(runes)
}
```

This approach ensures that multi-byte UTF-8 characters (which may be 2-4 bytes each) are treated as single units (runes) during reversal, preventing corruption of Unicode characters.

## Performance Considerations

### MustNotBeEmpty
- O(1) complexity - only checks string length
- Very fast regardless of string size
- Minimal memory overhead

### Reverse
- O(n) complexity where n is the number of runes in the string
- Creates a new rune slice of the same size as the input string
- Memory usage is approximately 2x the input size (original string + rune slice)
- Performance is acceptable for typical string sizes
- For very large strings, consider streaming processing

## Security Considerations

- `MustNotBeEmpty` helps prevent issues that could arise from empty string values
  - Prevents using empty API keys, passwords, or tokens
  - Ensures required input fields are provided
  - Helps avoid logic errors when empty strings have special meaning
- The functions do not perform any input sanitization
  - They are validation and manipulation functions, not security filters
  - Consider additional validation for security-critical applications

## Future Improvements

- Add additional string validation functions (e.g., MustBeEmail, MustBeURL, MustMatchRegex)
- Implement additional string manipulation functions (e.g., CamelCase, KebabCase, SnakeCase)
- Add string sanitization functions for security purposes
- Implement string comparison functions with normalization
- Add support for string templates
- Implement string compression/decompression
- Add support for string encryption/decryption
- Implement performance benchmarks
- Add support for streaming string processing for very large strings
- Implement case conversion functions (ToCamelCase, ToSnakeCase, etc.)
- Add support for string truncation with ellipsis
- Implement string distance algorithms (Levenshtein, etc.)