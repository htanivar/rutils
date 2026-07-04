package strutil

import (
	"testing"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "single character",
			input:    "a",
			expected: "a",
		},
		{
			name:     "basic string",
			input:    "hello",
			expected: "olleh",
		},
		{
			name:     "string with spaces",
			input:    "hello world",
			expected: "dlrow olleh",
		},
		{
			name:     "string with special characters",
			input:    "Hello, 世界!",
			expected: "!界世 ,olleH",
		},
		{
			name:     "palindrome",
			input:    "racecar",
			expected: "racecar",
		},
		{
			name:     "numbers",
			input:    "12345",
			expected: "54321",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Reverse(tc.input)
			if result != tc.expected {
				t.Errorf("Reverse(%q) = %q, want %q", tc.input, result, tc.expected)
			}
		})
	}
}

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse("Hello, Ravi!")
	}
}
