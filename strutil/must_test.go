package strutil

import (
	"errors"
	"testing"
)

func TestNotEmpty(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantErr   bool
		expectErr error
	}{
		{
			name:      "non-empty string",
			input:     "hello",
			wantErr:   false,
			expectErr: nil,
		},
		{
			name:      "empty string",
			input:     "",
			wantErr:   true,
			expectErr: ErrEmptyString,
		},
		{
			name:      "string with spaces",
			input:     "   ",
			wantErr:   false,
			expectErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NotEmpty(tt.input)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
				if tt.expectErr != nil && !errors.Is(err, tt.expectErr) {
					t.Errorf("Expected error %v, got %v", tt.expectErr, err)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}
			}
		})
	}
}

func TestMustNotBeEmpty(t *testing.T) {
	t.Run("should not panic on non-empty", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("MustNotBeEmpty panicked unexpectedly: %v", r)
			}
		}()
		MustNotBeEmpty("hello")
	})

	t.Run("should panic on empty", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Error("expected MustNotBeEmpty to panic")
			}
			err, ok := r.(error)
			if !ok || !errors.Is(err, ErrEmptyString) {
				t.Errorf("expected panic with ErrEmptyString, got %v", r)
			}
		}()
		MustNotBeEmpty("")
	})
}
