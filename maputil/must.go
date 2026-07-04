package maputil

import "errors"

var (
	ErrEmptyMap    = errors.New("map is empty")
	ErrMapNotEmpty = errors.New("map is not empty")
)

// IsEmpty returns nil if the map is empty, otherwise returns ErrMapNotEmpty.
func IsEmpty[K comparable, V any](m map[K]V) error {
	if len(m) != 0 {
		return ErrMapNotEmpty
	}
	return nil
}

// IsNotEmpty returns nil if the map is not empty, otherwise returns ErrEmptyMap.
func IsNotEmpty[K comparable, V any](m map[K]V) error {
	if len(m) == 0 {
		return ErrEmptyMap
	}
	return nil
}

// MustBeEmpty panics if the map is not empty.
func MustBeEmpty[K comparable, V any](m map[K]V) {
	if err := IsEmpty(m); err != nil {
		panic(err)
	}
}

// MustNotBeEmpty panics if the map is empty.
func MustNotBeEmpty[K comparable, V any](m map[K]V) {
	if err := IsNotEmpty(m); err != nil {
		panic(err)
	}
}
