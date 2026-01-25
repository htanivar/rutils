package path

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

// MustExists returns nil if the file at the given path exists, otherwise returns ErrNotExist
func MustExists(path string) error {
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

// MustNotExist returns nil if the file at the given path does not exist, otherwise returns ErrExists
func MustNotExist(path string) error {
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

// MustBeEmpty returns nil if the directory at the given path is empty, otherwise returns an error
func MustBeEmpty(path string) error {
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

// MustNotBeEmpty returns nil if the directory at the given path is not empty, otherwise returns an error
func MustNotBeEmpty(path string) error {
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
