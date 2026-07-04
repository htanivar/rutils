# Path Utilities (pathutil) Documentation

## Overview

The `pathutil` package provides utilities for validating file and directory paths. In alignment with Go standards:
- Standard functions (e.g. `Exists`, `BeEmpty`) return Go `error` values.
- Wrapper functions prefixed with `Must` (e.g. `MustExist`, `MustBeEmpty`) **panic** if validation fails.

---

## API Reference

### Functions

#### Exists
Returns `nil` if the path exists, otherwise `ErrNotExist` or other filesystem errors.
```go
func Exists(path string) error
```

#### NotExist
Returns `nil` if the path does not exist, otherwise `ErrExists` if it does, or filesystem errors.
```go
func NotExist(path string) error
```

#### BeEmpty
Returns `nil` if the directory exists and is empty. Returns errors if not.
```go
func BeEmpty(path string) error
```

#### NotBeEmpty
Returns `nil` if the directory exists and contains one or more items. Returns errors if empty.
```go
func NotBeEmpty(path string) error
```

---

### Panicking "Must" Wrappers

#### MustExist
Panics if the path does not exist.
```go
func MustExist(path string)
```

#### MustNotExist
Panics if the path exists.
```go
func MustNotExist(path string)
```

#### MustBeEmpty
Panics if the directory is not empty.
```go
func MustBeEmpty(path string)
```

#### MustNotBeEmpty
Panics if the directory is empty.
```go
func MustNotBeEmpty(path string)
```

---

## Error Types
```go
var (
	ErrNotExist     = errors.New("file does not exist")
	ErrExists       = errors.New("file already exists")
	ErrNotDirectory = errors.New("path is not a directory")
	ErrNotEmpty     = errors.New("directory is not empty")
	ErrEmpty        = errors.New("directory is empty")
)
```

---

## Code Examples

### Path Checks
```go
package main

import (
	"log"

	"github.com/htanivar/rutils/pathutil"
)

func main() {
	configPath := "/etc/myapp/config.json"

	// 1. Error-returning validation
	if err := pathutil.Exists(configPath); err != nil {
		log.Fatalf("missing config: %v", err)
	}

	// 2. Panicking assertion
	pathutil.MustExist(configPath)

	// 3. Directory emptiness
	pathutil.MustBeEmpty("/tmp/scratch_dir")
}
```