package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected float64
	}{
		{
			name: "example1",
			input: `
3
1 1
2 4
4 3
					`,
			expected: 3.605551,
		},
		{
			name: "example2",
			input: `
10
1 8
4 0
3 7
2 4
5 9
9 1
6 2
0 2
8 6
7 8
					`,
			expected: 10.630146,
		},
		{
			name: "example3",
			input: `
4
0 0
0 100
100 0
100 100
					`,
			expected: 141.421356,
		},
		{
			name: "example4",
			input: `
5
3 0
1 0
0 0
4 0
2 0
					`,
			expected: 4.000000,
		},
		{
			name: "example4",
			input: `
4
2 2
0 0
1 1
3 3
					`,
			expected: 4.242641,
		},
	}

	for _, tt := range tests {
		reader := bufio.NewReader(strings.NewReader(generic_TrimSpaceAndNewLineCodeAndTab(tt.input)))
		input, err := generic_NewInputFromReader(reader)
		if err != nil {
			t.Errorf("unexpected error occurred in test %s: %v", tt.name, err)
		}
		actual := solve(input)
		diff := math.Abs(actual - tt.expected)
		if diff >= math.Pow(10, -3) {
			t.Errorf("%s is expected to return %v when input %q is given, but actually return %v, diff: %v",
				tt.name, tt.expected, input.lines, actual, diff)
		}
	}
}
