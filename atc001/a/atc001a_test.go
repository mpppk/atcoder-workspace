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
		4 5
		s####
		....#
		#####
		#...g					`,
			expected: "No",
		},
		{
			name: "example2",
			input: `
		4 4
		...s
		....
		....
		.g..
		`,
			expected: "Yes",
		},
		{
			name: "example3",
			input: `
10 10
s.........
#########.
#.......#.
#..####.#.
##....#.#.
#####.#.#.
g.#.#.#.#.
#.#.#.#.#.
###.#.#.#.
#.....#...
`,
			expected: "No",
		},
		{
			name: "example4",
			input: `
		10 10
		s.........
		#########.
		#.......#.
		#..####.#.
		##....#.#.
		#####.#.#.
		g.#.#.#.#.
		#.#.#.#.#.
		#.#.#.#.#.
		#.....#...
		`,
			expected: "Yes",
		},
		{
			name: "example5",
			input: `
		1 10
		s..####..g
		`,
			expected: "No",
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
