# rutils

A collection of utility functions and tools written in Go.

## Features

- String manipulation utilities
- Mathematical operations
- Easy to use and extend

## Installation

```bash
go get github.com/yourusername/rutils
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/yourusername/rutils/pkg/math"
    "github.com/yourusername/rutils/internal/utils"
)

func main() {
    // Using math utilities
    sum := math.Add(5, 3)
    product := math.Multiply(5, 3)
    fmt.Printf("Sum: %d, Product: %d\n", sum, product)
    
    // Using string utilities
    reversed := utils.Reverse("hello")
    fmt.Printf("Reversed: %s\n", reversed)
}
```

## Project Structure

```
rutils/
├── cmd/
│   └── main.go          # Example main program
├── internal/
│   └── utils/
│       └── utils.go     # Internal utilities
├── pkg/
│   └── math/
│       └── math.go      # Public math utilities
├── go.mod               # Go module definition
└── README.md            # This file
```

## Development

To build and run the example:

```bash
go run cmd/main.go
```

To build a binary:

```bash
go build -o rutils cmd/main.go
```

## Testing

Run all tests:

```bash
go test ./...
```

Or use the provided Makefile:

```bash
make test           # Run all tests
make test-string    # Run only string package tests
make test-math      # Run only math package tests
```

To run tests for a specific package:

```bash
go test ./string
go test ./math
```

## License

MIT
