package main

import (
	"bufio"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name: "example1",
			input: `
		7
		`,
			expected: 2,
		},
		{
			name: "example1",
			input: `
				127
				`,
			expected: 4,
		},
		{
			name: "example1",
			input: `
				3
				`,
			expected: 3,
		},
		{
			name: "example1",
			input: `
		44852
		`,
			expected: 16,
		},
	}

	for _, tt := range tests {
		reader := bufio.NewReader(strings.NewReader(lib_TrimSpaceAndNewLineCodeAndTab(tt.input)))
		input, err := lib_NewInputFromReader(reader)
		if err != nil {
			t.Errorf("unexpected error occurred in test %s: %v", tt.name, err)
		}
		actual := solve(input)
		if actual != tt.expected {
			t.Errorf("%s is expected to return %v when input %q is given, but actually return %v",
				tt.name, tt.expected, input.lines, actual)
		}
	}
}
