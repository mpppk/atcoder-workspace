package a

import (
	"bufio"
	"github.com/mpppk/atcoder/done/abc159tmp/a"
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
			input: `
2 1
`,
			expected: 1,
		},
		{
			input: `
4 3
`,
			expected: 9,
		},
		{
			input: `
1 1
`,
			expected: 0,
		},
		{
			input: `
13 3
`,
			expected: 81,
		},
		{
			input: `
0 3
`,
			expected: 3,
		},
		{
			input: `
100 100
`,
			expected: 3,
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
