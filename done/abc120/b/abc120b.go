package main

import (
	"bufio"
	"fmt"
	"github.com/mpppk/atcoder/done/abc120/b"
	"io"
	"os"
)

func solve(input *b.lib_Input) int {
	line := input.MustGetIntLine(0)
	a, b, k := line[0], line[1], line[2]
	min := a
	if a > b {
		min = b
	}

	var answers []int
	for i := 1; i <= min; i++ {
		if a%i == 0 && b%i == 0 {
			answers = append(answers, i)
		}
	}
	return answers[len(answers)-k]
}

func main() {
	input := b.lib_MustNewInputFromReader(bufio.NewReader(io.Reader(os.Stdin)))
	fmt.Println(solve(input))
}
