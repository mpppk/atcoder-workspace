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
4
10
8
8
6
					`,
			expected: 3,
		},
		{
			name: "example2",
			input: `
3
15
15
15
			`,
			expected: 1,
		},
		{
			name: "example3",
			input: `
7
50
30
50
100
50
80
30
					`,
			expected: 4,
		},
	}

	for _, tt := range tests {
		scanner := bufio.NewScanner(strings.NewReader(utl_TrimSpaceAndNewLineCodeAndTab(tt.input)))
		input := utl_NewInput(scanner)
		actual := solve(input)
		if actual != tt.expected {
			t.Errorf("%s is expected to return %d when input %q is given, but actually return %d",
				tt.name, tt.expected, input.lines, actual)
		}
	}
}
