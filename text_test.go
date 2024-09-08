package components

import "testing"

func TestTextWithBoldStarsStr(t *testing.T) {
	table := []struct {
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
			name:     "no stars",
			input:    "Hello world",
			expected: "Hello world",
		},
		{
			name:     "one star",
			input:    "Hello *world",
			expected: "Hello *world",
		},
		{
			name:     "two stars",
			input:    "Hello *world*",
			expected: "Hello <strong>world</strong>",
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			actual := textWithBoldStarsStr(tt.input)
			if actual != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, actual)
			}
		})
	}
}
