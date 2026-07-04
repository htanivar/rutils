package pathutil

import (
	"errors"
	"os"
)

var (
	ErrNotExist     = errors.New("file does not exist")
	ErrExists       = errors.New("file already exists")
	ErrNotDirectory = errors.New("path is not a directory")
	ErrNotEmpty     = errors.New("directory is not empty")
	ErrEmpty        = errors.New("directory is empty")
)

// Exists returns nil if the file at the given path exists, otherwise returns ErrNotExist
func Exists(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return ErrNotExist
		}
		// For other errors (like permission issues), return them as is
		return err
	}
	return nil
}

// NotExist returns nil if the file at the given path does not exist, otherwise returns ErrExists
func NotExist(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // File doesn't exist - this is what we want
		}
		// For other errors (like permission issues), return them as is
		return err
	}
	return ErrExists // File exists - this is an error for this function
}

// BeEmpty returns nil if the directory at the given path is empty, otherwise returns an error
func BeEmpty(path string) error {
	// First check if path exists
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return ErrNotExist
		}
		return err
	}

	// Check if it's a directory
	if !info.IsDir() {
		return ErrNotDirectory
	}

	// Read directory contents
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	// Check if empty
	if len(entries) > 0 {
		return ErrNotEmpty
	}

	return nil
}

// NotBeEmpty returns nil if the directory at the given path is not empty, otherwise returns an error
func NotBeEmpty(path string) error {
	// First check if path exists
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return ErrNotExist
		}
		return err
	}

	// Check if it's a directory
	if !info.IsDir() {
		return ErrNotDirectory
	}

	// Read directory contents
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	// Check if empty
	if len(entries) == 0 {
		return ErrEmpty
	}

	return nil
}

// MustExist panics if the file at the given path does not exist.
func MustExist(path string) {
	if err := Exists(path); err != nil {
		panic(err)
	}
}

// MustNotExist panics if the file at the given path exists.
func MustNotExist(path string) {
	if err := NotExist(path); err != nil {
		panic(err)
	}
}

// MustBeEmpty panics if the directory at the given path is not empty.
func MustBeEmpty(path string) {
	if err := BeEmpty(path); err != nil {
		panic(err)
	}
}

// MustNotBeEmpty panics if the directory at the given path is empty.
func MustNotBeEmpty(path string) {
	if err := NotBeEmpty(path); err != nil {
		panic(err)
	}
}
