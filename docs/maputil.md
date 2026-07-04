# Map Utilities (maputil) Documentation

## Overview

The `maputil` package provides type-safe utilities for validating Go map data structures. By leveraging Go generics, it supports any map key and value types without relying on empty interface assertions (`interface{}`/`any`).

---

## API Reference

### Functions

#### IsEmpty
Checks if a map is empty. Returns `ErrMapNotEmpty` if it has elements.
```go
func IsEmpty[K comparable, V any](m map[K]V) error
```

#### IsNotEmpty
Checks if a map is not empty. Returns `ErrEmptyMap` if the map is empty or nil.
```go
func IsNotEmpty[K comparable, V any](m map[K]V) error
```

#### MustBeEmpty
Validates that a map is empty, panicking if it contains elements.
```go
func MustBeEmpty[K comparable, V any](m map[K]V)
```

#### MustNotBeEmpty
Validates that a map is not empty, panicking if it is empty or nil.
```go
func MustNotBeEmpty[K comparable, V any](m map[K]V)
```

---

## Error Types
```go
var (
	ErrEmptyMap    = errors.New("map is empty")
	ErrMapNotEmpty = errors.New("map is not empty")
)
```

---

## Code Examples

### Generic Map Validation
```go
package main

import (
	"fmt"
	"log"

	"github.com/htanivar/rutils/maputil"
)

func main() {
	// Works automatically on different map types!
	config := map[string]string{
		"env": "production",
	}

	userSessions := map[int]bool{
		12345: true,
	}

	// 1. Check non-emptiness using standard error-returning pattern
	if err := maputil.IsNotEmpty(config); err != nil {
		log.Fatalf("config is missing values: %v", err)
	}

	// 2. Check panicking Must wrapper
	maputil.MustNotBeEmpty(userSessions)

	// 3. Check emptiness validation
	emptyCache := map[string][]byte{}
	if err := maputil.IsEmpty(emptyCache); err == nil {
		fmt.Println("Cache is empty.")
	}
}
```