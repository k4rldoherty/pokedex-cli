package repl

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	// Start by creating a slice of test structs
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "     world  ",
			expected: []string{"world"},
		},
		{
			input:    "  hello   world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  hello   world CAPS ",
			expected: []string{"hello", "world", "caps"},
		},
	}

	// Loop through the cases and do the testing
	for _, c := range cases {
		actual := CleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("\nTest Failed.\nExpected: %v\nGot %v", c.expected, actual)
		}
		for i := range actual {
			if c.expected[i] != actual[i] {
				t.Errorf("\nTest Failed.\nExpected: %v\nGot %v", c.expected[i], actual[i])
			}
		}
	}
}
