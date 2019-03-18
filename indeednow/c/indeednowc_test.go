package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []int
	}{
		{
			name: "example1",
			input: `
3 6
1 2 3 3
3 3 3 6
4 4 4 8
3 4 3
4 4 4
100 100 1
2 3 4
0 0 0
100 100 100
		`,
			expected: []int{6, 8, 0, 3, 0, 8},
		},
	}

	for _, tt := range tests {
		reader := bufio.NewReader(strings.NewReader(lib_TrimSpaceAndNewLineCodeAndTab(tt.input)))
		input, err := lib_NewInputFromReader(reader)
		if err != nil {
			t.Errorf("unexpected error occurred in test %s: %v", tt.name, err)
		}
		actual := solve(input)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("%s is expected to return %v when input %q is given, but actually return %v",
				tt.name, tt.expected, input.lines, actual)
		}
	}
}
