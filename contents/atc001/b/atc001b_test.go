package main

import (
	"bufio"
	"github.com/mpppk/atcoder/done"
	"reflect"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name: "example1",
			input: `
8 9
0 1 2
0 3 2
1 1 3
1 1 4
0 2 4
1 4 1
0 4 2
0 0 0
1 0 0
`,
			expected: []string{"Yes", "No", "Yes", "Yes"},
		},
	}

	for _, tt := range tests {
		reader := bufio.NewReader(strings.NewReader(done.lib_TrimSpaceAndNewLineCodeAndTab(tt.input)))
		input, err := done.lib_NewInputFromReader(reader)
		if err != nil {
			t.Errorf("unexpected error occurred in test %s: %v", tt.name, err)
		}
		actual := done.solve(input)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("%s is expected to return %v when input %q is given, but actually return %v",
				tt.name, tt.expected, input.lines, actual)
		}
	}
}
