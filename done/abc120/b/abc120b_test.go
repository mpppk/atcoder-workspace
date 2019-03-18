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
8 12 2
					`,
			expected: 2,
		},
		{
			name: "example1",
			input: `
12 8 2
					`,
			expected: 2,
		},
		{
			name: "example1",
			input: `
100 50 4
					`,
			expected: 5,
		},
		{
			name: "example1",
			input: `
50 100 4
					`,
			expected: 5,
		},
		{
			name: "example1",
			input: `
2 4 1
					`,
			expected: 1,
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
