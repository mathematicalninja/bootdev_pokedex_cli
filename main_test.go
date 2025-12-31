package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "over-spacing",
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			name:     "Capitalisation",
			input:    "Pikachu Bulbasaur DITTO",
			expected: []string{"pikachu", "bulbasaur", "ditto"},
		},
		{
			name:     "empty string",
			input:    "",
			expected: []string{},
		},
		{
			name:     "single word",
			input:    "mew",
			expected: []string{"mew"},
		},
		{
			name:     "dasterdly tabs",
			input:    "mew\t mewtwo",
			expected: []string{"mew", "mewtwo"},
		},
		// add more cases here
	}
	for _, c := range cases {
		actual := CleanInput(c.input)
		if len(c.expected) != len(actual) {
			// fail
			t.Errorf("test %s, slice's of different length. Expected: %v, Actual: %v", c.name, c.expected, actual)
		}
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		for i, expectedWord := range c.expected {
			word := actual[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("test %s, expected: %s, actual: %s", c.name, expectedWord, word)
			}
		}
	}
}
