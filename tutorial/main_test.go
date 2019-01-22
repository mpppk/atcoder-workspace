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
1
2 3
test
			`,
			expected: "6 test",
		},
	}

	for _, tt := range tests {
		scanner := bufio.NewScanner(strings.NewReader(utl_TrimSpaceAndNewLineCodeAndTab(tt.input)))
		input := utl_NewInput(scanner)
		actual := solve(input)
		if actual != tt.expected {
			t.Fatalf("%s is expected to return %q when input %q is given, but actually return %q",
				tt.name, tt.expected, tt.input, actual)
		}
	}
}
