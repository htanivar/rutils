# rutils

rutils is a Go utility library that provides various helper functions for common tasks such as email notifications, file validation, path checking, map validation, string utilities, and data conversion. The library follows a consistent pattern of "must" functions that return errors when validation conditions are not met.

## Features

- **Email Notifications**: Send emails with optional attachments and HTML body
- **File Validation**: Validate file types based on extension and content (CSV, JSON, XML, PDF, ZIP, TXT)
- **Path Utilities**: Check if files and directories exist, are empty, or meet other conditions
- **Map Validation**: Check if maps are empty or not empty
- **String Utilities**: Check if strings are empty and reverse strings
- **CSV to JSON Conversion**: Convert CSV files to JSON records

## Installation

```bash
go get github.com/htanivar/rutils
```

## Usage

Each package in rutils provides "must" functions that return errors when conditions are not met. These functions are designed to be used in validation scenarios where you want to ensure certain conditions are met before proceeding.

### Email Notifications

```go
import "github.com/htanivar/rutils/email"

client := email.NewClient("smtp.gmail.com", 587, "user@gmail.com", "password")

msg := &email.Message{
	To:      []string{"recipient@example.com"},
	Subject: "Test Email",
	Body:    "This is the email body",
}

if err := client.Send(msg); err != nil {
	log.Fatal(err)
}
```

### File Validation

```go
import "github.com/htanivar/rutils/file"

// Validate that a file is a CSV file
if err := file.MustBeType(".csv", "/path/to/file.csv"); err != nil {
	log.Fatal(err)
}
```

### Path Utilities

```go
import "github.com/htanivar/rutils/path"

// Check if a file exists
if err := path.MustExists("/path/to/file"); err != nil {
	log.Fatal(err)
}

// Check if a directory is empty
if err := path.MustBeEmpty("/path/to/dir"); err != nil {
	log.Fatal(err)
}
```

### Map Validation

```go
import "github.com/htanivar/rutils/maps"

m := map[string]interface{}{"key": "value"}

// Check if a map is not empty
if err := maps.MustNotBeEmpty(m); err != nil {
	log.Fatal(err)
}
```

### String Utilities

```go
import "github.com/htanivar/rutils/string"

// Check if a string is not empty
s := "hello"
if err := string.MustNotBeEmpty(s); err != nil {
	log.Fatal(err)
}

// Reverse a string
reversed := string.Reverse(s) // "olleh"
```

### CSV to JSON Conversion

```go
import "github.com/htanivar/rutils/task"

// Convert CSV file to JSON records
records, err := task.CSVToRecords("/path/to/file.csv")
if err != nil {
	log.Fatal(err)
}
// records is a slice of map[string]string
```

## Package Structure

- **email**: Email notification functionality
- **file**: File type validation
- **path**: Path and directory validation
- **maps**: Map validation
- **string**: String validation and manipulation
- **task**: Data conversion utilities

## Error Handling

All "must" functions return specific error types that can be checked:

```go
import "errors"

if errors.Is(err, file.ErrInvalidType) {
	// Handle invalid file type
} else if errors.Is(err, path.ErrNotExist) {
	// Handle file not existing
}
```

## Testing

Run all tests:

```bash
make test
```

Run tests for a specific package:

```bash
make test-email
make test-file
make test-path
make test-maps
make test-string
make test-task
```
