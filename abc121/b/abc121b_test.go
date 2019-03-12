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
2 3 -10
1 2 3
3 2 1
1 2 2
					`,
			expected: 1,
		},
		{
			name: "example1",
			input: `
5 2 -4
-2 5
100 41
100 40
-3 0
-6 -2
18 -13
					`,
			expected: 2,
		},
		{
			name: "example1",
			input: `
3 3 0
100 -100 0
0 100 100
100 100 100
-100 100 100
					`,
			expected: 0,
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
