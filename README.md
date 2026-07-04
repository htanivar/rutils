# rutils

`rutils` is a modern, high-efficiency utility library for Go projects. It is built to leverage Go generics, modern context-based patterns, and idiomatic Go programming paradigms (such as distinguishing between error-returning validation functions and panicking `Must` helper functions).

## Features

- **`email`**: SMTP email client featuring HTML body alternatives, file attachments, connection reuse for batch sends, and first-class **`context.Context`** cancellation/timeout support.
- **`file`**: Structured file type and content validation (CSV, JSON, XML, PDF, ZIP, TXT) using robust parsing and signature checking.
- **`pathutil`**: Idiomatic file and directory existence/emptiness checks.
- **`maputil`**: Type-safe generic map validation using Go generics (`[K comparable, V any]`).
- **`strutil`**: String validity verification and rune-wise reversing.
- **`csvutil`**: Parsing and conversion utilities for reading CSV documents into maps.

---

## Installation

```bash
go get github.com/htanivar/rutils
```

---

## Usage

All utility packages implement two patterns:
1. **Standard functions** (e.g., `Exists`, `ValidateType`, `IsEmpty`) which return standard Go `error` values.
2. **Panicking "Must" wrappers** (e.g., `MustExist`, `MustBeType`, `MustBeEmpty`) which automatically panic if validation fails, ideal for inline initialization.

### 1. Email Notifications
Provides a thread-safe SMTP client with context support for connection timeouts.

```go
import (
	"context"
	"log"
	"time"

	"github.com/htanivar/rutils/email"
)

client := email.NewClient(&email.Config{
	Host:     "smtp.gmail.com",
	Port:     587,
	Username: "user@gmail.com",
	Password: "password",
	From:     "sender@gmail.com",
})

ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

err := client.SendWithContext(ctx, &email.Message{
	To:       []string{"recipient@example.com"},
	Subject:  "Greetings",
	Body:     "Plain text body alternative",
	HTMLBody: "<h1>Hello!</h1><p>HTML body</p>",
})
if err != nil {
	log.Fatal(err)
}
```

### 2. File Validation
Ensures files match their declared extensions by validating actual content signatures, parsing structure, or reading magic headers.

```go
import (
	"log"

	"github.com/htanivar/rutils/file"
)

// Returns an error if the file isn't a valid CSV
if err := file.ValidateType("csv", "report.csv"); err != nil {
	log.Fatal(err)
}

// Panics inline if not a valid JSON structure
file.MustBeType(".json", "config.json")
```

### 3. Path Utilities
Checks paths for existence, folder emptiness, and non-emptiness.

```go
import (
	"github.com/htanivar/rutils/pathutil"
)

// Non-panicking
if err := pathutil.Exists("/data/logs"); err != nil {
	log.Printf("directory missing: %v", err)
}

// Panicking helper
pathutil.MustExist("/etc/config.json")
pathutil.MustBeEmpty("/tmp/empty_dir")
```

### 4. Map Validation (Generics)
Generics-based type-safe validation for Go maps.

```go
import (
	"github.com/htanivar/rutils/maputil"
)

m := map[int]string{1: "active"}

// Check if a map is not empty
if err := maputil.IsNotEmpty(m); err != nil {
	log.Fatal(err)
}

// Panics inline if empty
maputil.MustNotBeEmpty(m)
```

### 5. String Utilities
```go
import (
	"fmt"

	"github.com/htanivar/rutils/strutil"
)

// Check non-emptiness
if err := strutil.NotEmpty("hello"); err != nil {
	fmt.Println("string is empty")
}

// Reverse unicode string
reversed := strutil.Reverse("Hello, 世界") // "界世 ,olleH"
```

### 6. CSV parsing
```go
import (
	"log"

	"github.com/htanivar/rutils/csvutil"
)

records, err := csvutil.CSVToRecords("contacts.csv")
if err != nil {
	log.Fatal(err)
}
// records is []map[string]string
```

---

## Package Directory Structure

- **`email`**: SMTP email delivery (batch sending & context support).
- **`file`**: File extension & MIME content structure verification.
- **`pathutil`**: Directory/file status checks.
- **`maputil`**: Generic-safe map emptiness checks.
- **`strutil`**: Basic string verification & rune reversal.
- **`csvutil`**: Tabular CSV parsing.

---

## Testing

Run all package tests:
```bash
make test
```

Or run package-specific tests:
```bash
make test-email
make test-file
make test-path
make test-maps
make test-string
make test-task
```
