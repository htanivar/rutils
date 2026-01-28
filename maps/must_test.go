package maps

import "testing"

func TestMustBeEmpty(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		in      map[string]interface{}
		wantErr error
	}{
		{
			name:    "nil map is empty",
			in:      nil,
			wantErr: nil,
		},
		{
			name:    "empty map is empty",
			in:      map[string]interface{}{},
			wantErr: nil,
		},
		{
			name:    "non-empty map returns ErrMapNotEmpty",
			in:      map[string]interface{}{"k": 1},
			wantErr: ErrMapNotEmpty,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if err := MustBeEmpty(tc.in); err != tc.wantErr {
				t.Fatalf("MustBeEmpty(%v) err=%v want=%v", tc.in, err, tc.wantErr)
			}
		})
	}
}

func TestMustNotBeEmpty(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		in      map[string]interface{}
		wantErr error
	}{
		{
			name:    "nil map returns ErrEmptyMap",
			in:      nil,
			wantErr: ErrEmptyMap,
		},
		{
			name:    "empty map returns ErrEmptyMap",
			in:      map[string]interface{}{},
			wantErr: ErrEmptyMap,
		},
		{
			name:    "non-empty map is ok",
			in:      map[string]interface{}{"k": 1},
			wantErr: nil,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if err := MustNotBeEmpty(tc.in); err != tc.wantErr {
				t.Fatalf("MustNotBeEmpty(%v) err=%v want=%v", tc.in, err, tc.wantErr)
			}
		})
	}
}
