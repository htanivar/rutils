#!/bin/bash

# Initialize a Go module called rutils in the current directory
# Assumes you are already in the 'rutils' folder

set -e

MODULE_NAME="rutils"
CURRENT_DIR=$(pwd)

echo "Initializing Go module: $MODULE_NAME in current directory"

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed or not in PATH"
    echo "Please install Go and try again"
    exit 1
fi

# Check if we're in the right directory by checking the last part of the path
# This is optional, but can be helpful
BASENAME=$(basename "$CURRENT_DIR")
if [ "$BASENAME" != "rutils" ]; then
    echo "Warning: Current directory name is '$BASENAME', not 'rutils'"
    echo "Continuing anyway..."
fi

# Check if go.mod already exists
if [ -f "go.mod" ]; then
    echo "Error: go.mod already exists in the current directory"
    echo "Aborting to avoid overwriting existing module"
    exit 1
fi

# Initialize Go module in current directory
echo "Running: go mod init $MODULE_NAME"
if ! go mod init "$MODULE_NAME"; then
    echo "Error: Failed to initialize Go module"
    exit 1
fi

# Create basic directory structure
echo "Creating directory structure..."
mkdir -p cmd internal pkg

# Create a main.go file
echo "Creating cmd/main.go..."
cat > cmd/main.go << 'EOF'
package main

import "fmt"

func main() {
    fmt.Println("rutils module initialized")
}
EOF

# Create a sample internal package
echo "Creating internal/utils/utils.go..."
mkdir -p internal/utils
cat > internal/utils/utils.go << 'EOF'
package utils

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}
EOF

# Create a sample pkg package
echo "Creating pkg/math/math.go..."
mkdir -p pkg/math
cat > pkg/math/math.go << 'EOF'
package math

// Add returns the sum of two integers
func Add(a, b int) int {
    return a + b
}

// Multiply returns the product of two integers
func Multiply(a, b int) int {
    return a * b
}
EOF

# Create or update README.md
echo "Updating README.md..."
cat > README.md << 'EOF'
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

## License

MIT
EOF

# Create or update .gitignore for Go projects
echo "Updating .gitignore for Go..."
cat > .gitignore << 'EOF'
# Go
*.exe
*.exe~
*.dll
*.so
*.dylib
*.test
*.out
*.prof

# Binaries for programs and plugins
bin/
dist/

# Dependency directories
vendor/

# Go workspace file
go.work

# IDE
.vscode/
.idea/
*.swp
*.swo

# OS generated files
.DS_Store
.DS_Store?
._*
.Spotlight-V100
.Trashes
ehthumbs.db
Thumbs.db

# Aider
.aider*
.aider.chat.history.md
.aider.input.history
.aider.tags.cache.v4/
command_log*

# Logs
logs
*.log

# Temporary files
tmp/
temp/

# Secrets
secret*
secret.sh
.env
.env.local
.env.development.local
.env.test.local
.env.production.local

# Build artifacts
coverage.txt
profile.out

# Misc
*--system_prompt=*
EOF

echo ""
echo "Go module '$MODULE_NAME' has been successfully initialized in $CURRENT_DIR"
echo ""
echo "Directory structure:"
find . -type f -name "*.go" | sort
echo ""
echo "Files created:"
echo "  go.mod"
echo "  cmd/main.go"
echo "  internal/utils/utils.go"
echo "  pkg/math/math.go"
echo "  README.md"
echo "  .gitignore"
echo ""
echo "To run the main program:"
echo "  go run cmd/main.go"
echo ""
echo "To build:"
echo "  go build -o rutils cmd/main.go"
echo ""
echo "Next steps:"
echo "  1. Update the module path in go.mod if needed"
echo "  2. Review and customize README.md"
echo "  3. Run 'git init' to initialize a Git repository"
echo "  4. Add files with 'git add .'"
echo "  5. Commit with 'git commit -m \"Initial commit\"'"
