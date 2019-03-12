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
		expected int64
	}{
		{
			name: "example1",
			input: `
2 5
4 9
2 4
					`,
			expected: 12,
		},
		{
			name: "example1",
			input: `
4 30
6 18
2 5
3 10
7 9
							`,
			expected: 130,
		},
		{
			name: "example1",
			input: `
1 100000
1000000000 100000
					`,
			expected: 100000000000000,
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
