package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/mpppk/atcoder-workspace/lib"
)

func solve(input *lib.Input) int {
	line := input.MustGetIntLine(0)
	N := line[0]
	K := line[1]
	m := map[int]bool{}
	for i := 1; i <= N; i++ {
		m[i] = false
	}

	for i := 1; i <= K*2; i = i + 2 {
		//d := input.MustGetIntLine(i)
		a := input.MustGetIntLine(i + 1)

		for _, ai := range a {
			m[ai] = true
		}
	}

	cnt := 0
	for _, b := range m {
		if !b {
			cnt++
		}
	}
	return cnt
}

func main() {
	input := lib.MustNewInputFromReader(bufio.NewReader(io.Reader(os.Stdin)))
	fmt.Println(solve(input))
}
