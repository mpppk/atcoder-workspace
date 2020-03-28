package a

import (
	"bufio"
	"github.com/mpppk/atcoder/done/abc123/a"
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
29
20
7
35
120
					`,
			expected: 215,
		},
		{
			name: "example1",
			input: `
101
86
119
108
57
					`,
			expected: 481,
		},
	}

	for _, tt := range tests {
		reader := bufio.NewReader(strings.NewReader(lib_TrimSpaceAndNewLineCodeAndTab(tt.input)))
		input, err := lib_NewInputFromReader(reader)
		if err != nil {
			t.Errorf("unexpected error occurred in test %s: %v", tt.name, err)
		}
		actual := a.solve(input)
		if actual != tt.expected {
			t.Errorf("%s is expected to return %v when input %q is given, but actually return %v",
				tt.name, tt.expected, input.lines, actual)
		}
	}
}
