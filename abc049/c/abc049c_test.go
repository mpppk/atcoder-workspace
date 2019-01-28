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
		expected string
	}{
		{
			name: "example1",
			input: `
erasedream
					`,
			expected: "YES",
		},
		{
			name: "example2",
			input: `
dreameraser
			`,
			expected: "YES",
		},
		{
			name: "example3",
			input: `
dreamerer
					`,
			expected: "NO",
		},
		{
			name:     "example4",
			input:    strings.Repeat("dreamer", 10000),
			expected: "YES",
		},
	}

	for _, tt := range tests {
		reader := bufio.NewReader(strings.NewReader(utl_TrimSpaceAndNewLineCodeAndTab(tt.input)))
		input, err := utl_NewInputFromReader(reader)
		if err != nil {
			t.Errorf("unexpected error occurred in test %s: %v", tt.name, err)
		}
		actual := solve(input)
		if actual != tt.expected {
			t.Errorf("%s is expected to return %s when input %q is given, but actually return %v",
				tt.name, tt.expected, input.lines, actual)
		}
	}
}
