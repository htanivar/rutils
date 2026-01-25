package string

import "errors"

var (
	ErrEmptyString = errors.New("string is empty")
)

// MustNotBeEmpty returns nil if the string is not empty, otherwise returns ErrEmptyString
func MustNotBeEmpty(s string) error {
	if s == "" {
		return ErrEmptyString
	}
	return nil
}
