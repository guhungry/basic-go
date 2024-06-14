package test

import (
	"testing"
)

func TestTransformText(t *testing.T) {
	cases := []struct {
		input       string
		expected    string
		description string
	}{
		{
			input:       "",
			expected:    "",
			description: "case empty string",
		},
		{
			input:       "a",
			expected:    "a",
			description: "return self",
		},
		{
			input:       "aB",
			expected:    "ab",
			description: "return lowercase",
		},
		{
			input:       " ac ",
			expected:    "ac",
			description: "return trim whitespaces",
		},
		{
			input:       "a d",
			expected:    "ad",
			description: "return remove space between",
		},
		{
			input:       "1",
			expected:    "one",
			description: "return replace number",
		},
		{
			input:       " 1 one4",
			expected:    "oneonefour",
			description: "return replace number",
		},
	}

	for _, test := range cases {
		actual := TransformText(test.input)
		if actual != test.expected {
			t.Errorf(`TransformText("%v") got = "%v", expected = "%v", %s`, test.input, actual, test.expected, test.description)
		}
	}
}
