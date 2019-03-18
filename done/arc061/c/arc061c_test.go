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
0
					`,
			expected: 0,
		},
		{
			name: "example1",
			input: `
1
					`,
			expected: 1,
		},
		{
			name: "example1",
			input: `
51
					`,
			expected: 57,
		},
		{
			name: "example1",
			input: `
125
					`,
			expected: 176,
		},
		{
			name: "example2",
			input: `
9999999999
					`,
			expected: 12656242944,
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
			t.Errorf("%s is expected to return %v when input %q is given, but actually return %v",
				tt.name, tt.expected, input.lines, actual)
		}
	}
}
