# String Utilities (strutil) Documentation

## Overview

The `strutil` package provides validation and manipulation utilities for string data. It includes support for checking non-emptiness (both returning errors and panicking) and reversing UTF-8 Unicode strings.

---

## API Reference

### Functions

#### NotEmpty
Validates that a string is not empty. Returns `ErrEmptyString` if it is.
```go
func NotEmpty(s string) error
```

#### MustNotBeEmpty
Validates that a string is not empty, panicking if it is.
```go
func MustNotBeEmpty(s string)
```

#### Reverse
Reverses a string rune-wise, fully supporting multi-byte UTF-8 Unicode characters (such as emojis and non-ASCII alphabets).
```go
func Reverse(s string) string
```

---

## Error Types
```go
var ErrEmptyString = errors.New("string is empty")
```

---

## Code Examples

### Validating & Reversing Strings
```go
package main

import (
	"fmt"
	"log"

	"github.com/htanivar/rutils/strutil"
)

func main() {
	input := "Hello, 世界!"

	// Non-panicking validation
	if err := strutil.NotEmpty(input); err != nil {
		log.Fatalf("input is empty: %v", err)
	}

	// Panicking validation
	strutil.MustNotBeEmpty(input)

	// Unicode-safe reversal
	reversed := strutil.Reverse(input)
	fmt.Printf("Original: %s\nReversed: %s\n", input, reversed)
	// Output:
	// Original: Hello, 世界!
	// Reversed: !界世 ,olleH
}
```