package maputil

import (
	"errors"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	t.Parallel()

	t.Run("string interface map", func(t *testing.T) {
		tests := []struct {
			name    string
			in      map[string]any
			wantErr error
		}{
			{
				name:    "nil map is empty",
				in:      nil,
				wantErr: nil,
			},
			{
				name:    "empty map is empty",
				in:      map[string]any{},
				wantErr: nil,
			},
			{
				name:    "non-empty map returns ErrMapNotEmpty",
				in:      map[string]any{"k": 1},
				wantErr: ErrMapNotEmpty,
			},
		}

		for _, tc := range tests {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				if err := IsEmpty(tc.in); !errors.Is(err, tc.wantErr) {
					t.Fatalf("IsEmpty(%v) err=%v want=%v", tc.in, err, tc.wantErr)
				}
			})
		}
	})

	t.Run("int string map", func(t *testing.T) {
		m := map[int]string{1: "one"}
		if err := IsEmpty(m); !errors.Is(err, ErrMapNotEmpty) {
			t.Fatalf("expected ErrMapNotEmpty, got %v", err)
		}

		emptyM := map[int]string{}
		if err := IsEmpty(emptyM); err != nil {
			t.Fatalf("expected nil error, got %v", err)
		}
	})
}

func TestIsNotEmpty(t *testing.T) {
	t.Parallel()

	t.Run("string interface map", func(t *testing.T) {
		tests := []struct {
			name    string
			in      map[string]any
			wantErr error
		}{
			{
				name:    "nil map returns ErrEmptyMap",
				in:      nil,
				wantErr: ErrEmptyMap,
			},
			{
				name:    "empty map returns ErrEmptyMap",
				in:      map[string]any{},
				wantErr: ErrEmptyMap,
			},
			{
				name:    "non-empty map is ok",
				in:      map[string]any{"k": 1},
				wantErr: nil,
			},
		}

		for _, tc := range tests {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				if err := IsNotEmpty(tc.in); !errors.Is(err, tc.wantErr) {
					t.Fatalf("IsNotEmpty(%v) err=%v want=%v", tc.in, err, tc.wantErr)
				}
			})
		}
	})
}

func TestMustBeEmpty(t *testing.T) {
	t.Run("should not panic on empty", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("MustBeEmpty panicked unexpectedly: %v", r)
			}
		}()
		MustBeEmpty(map[string]int{})
	})

	t.Run("should panic on non-empty", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Error("expected MustBeEmpty to panic")
			}
			err, ok := r.(error)
			if !ok || !errors.Is(err, ErrMapNotEmpty) {
				t.Errorf("expected panic with ErrMapNotEmpty, got %v", r)
			}
		}()
		MustBeEmpty(map[string]int{"k": 1})
	})
}

func TestMustNotBeEmpty(t *testing.T) {
	t.Run("should not panic on non-empty", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("MustNotBeEmpty panicked unexpectedly: %v", r)
			}
		}()
		MustNotBeEmpty(map[string]int{"k": 1})
	})

	t.Run("should panic on empty", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Error("expected MustNotBeEmpty to panic")
			}
			err, ok := r.(error)
			if !ok || !errors.Is(err, ErrEmptyMap) {
				t.Errorf("expected panic with ErrEmptyMap, got %v", r)
			}
		}()
		MustNotBeEmpty(map[string]int{})
	})
}
