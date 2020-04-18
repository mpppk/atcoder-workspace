package c

import (
	"bufio"
	"github.com/mpppk/atcoder/done/abc123/c"
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
5
3
2
4
3
5
					`,
			expected: 7,
		},
		{
			name: "example1",
			input: `
10
123
123
123
123
123
					`,
			expected: 5,
		},
		{
			name: "example1",
			input: `
10000000007
2
3
5
7
11
					`,
			expected: 5000000008,
		},
	}

	for _, tt := range tests {
		reader := bufio.NewReader(strings.NewReader(lib_TrimSpaceAndNewLineCodeAndTab(tt.input)))
		input, err := lib_NewInputFromReader(reader)
		if err != nil {
			t.Errorf("unexpected error occurred in test %s: %v", tt.name, err)
		}
		actual := c.solve(input)
		if actual != tt.expected {
			t.Errorf("%s is expected to return %v when input %q is given, but actually return %v",
				tt.name, tt.expected, input.lines, actual)
		}
	}
}
