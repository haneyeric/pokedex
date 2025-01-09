package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello   world ",
			expected: []string{"hello", "world"},
		},
		{
			input: "CAPS lock	",
			expected: []string{"caps", "lock"},
		},
		{
			input: "	123 j k l",
			expected: []string{"123", "j", "k", "l"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Mismatch length actual: %v, expected: %v", actual, c.expected)
		}

		for i := range actual {
			word := actual[i]
			expected := c.expected[i]
			if word != expected {
				t.Errorf("Words don't match, word: %s, expected: %s", word, expected)
			}
		}
	}
}
