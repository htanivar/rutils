package strutil

import "errors"

var (
	ErrEmptyString = errors.New("string is empty")
)

// NotEmpty returns nil if the string is not empty, otherwise returns ErrEmptyString
func NotEmpty(s string) error {
	if s == "" {
		return ErrEmptyString
	}
	return nil
}

// MustNotBeEmpty panics if the string is empty.
func MustNotBeEmpty(s string) {
	if err := NotEmpty(s); err != nil {
		panic(err)
	}
}
