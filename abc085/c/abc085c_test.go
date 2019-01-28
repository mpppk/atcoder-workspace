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
9 45000
					`,
			expected: "4 0 5",
		},
		{
			name: "example2",
			input: `
20 196000
			`,
			expected: "-1 -1 -1",
		},
		{
			name: "example3",
			input: `
1000 1234000
					`,
			expected: "14 27 959",
		},
		{
			name: "example4",
			input: `
2000 20000000
					`,
			expected: "2000 0 0",
		},
	}

	for _, tt := range tests {
		scanner := bufio.NewScanner(strings.NewReader(utl_TrimSpaceAndNewLineCodeAndTab(tt.input)))
		input := utl_NewInput(scanner)
		actual := solve(input)
		if actual != tt.expected {
			t.Errorf("%s is expected to return %s when input %q is given, but actually return %v",
				tt.name, tt.expected, input.lines, actual)
		}
	}
}
