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
2
3 1 2
6 1 1
					`,
			expected: "Yes",
		},
		{
			name: "example2",
			input: `
1
2 100 100
					`,
			expected: "No",
		},
		{
			name: "example3",
			input: `
2
5 1 1
100 1 1
					`,
			expected: "No",
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
