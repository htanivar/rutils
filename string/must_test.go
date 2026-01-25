package string

import "testing"

func TestMustNotBeEmpty(t *testing.T) {
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
		{
			name:      "string with single character",
			input:     "a",
			wantErr:   false,
			expectErr: nil,
		},
		{
			name:      "string with newline",
			input:     "\n",
			wantErr:   false,
			expectErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := MustNotBeEmpty(tt.input)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
				if tt.expectErr != nil && err != tt.expectErr {
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
