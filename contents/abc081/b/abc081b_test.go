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
3
8 12 40
			`,
			expected: 2,
		},
		{
			name: "example2",
			input: `
4
5 6 8 10
			`,
			expected: 0,
		},
	}

	for _, tt := range tests {
		scanner := bufio.NewScanner(strings.NewReader(utl_TrimSpaceAndNewLineCodeAndTab(tt.input)))
		input := utl_NewInput(scanner)
		actual := solve(input)
		if actual != tt.expected {
			t.Fatalf("%s is expected to return %d when input %q is given, but actually return %d",
				tt.name, tt.expected, tt.input, actual)
		}
	}
}
