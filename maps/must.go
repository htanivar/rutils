package maps

import "errors"

var (
	ErrEmptyMap    = errors.New("Map is empty")
	ErrMapNotEmpty = errors.New("Map is NOT empty")
)

func MustBeEmpty(m map[string]interface{}) error {
	if len(m) != 0 {
		return ErrMapNotEmpty
	}
	return nil
}

func MustNotBeEmpty(m map[string]interface{}) error {
	if len(m) == 0 {
		return ErrEmptyMap
	}
	return nil
}
