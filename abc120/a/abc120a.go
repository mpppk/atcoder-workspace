package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func solve(input *lib_Input) int {
	line := input.MustGetIntLine(0)
	a, b, c := line[0], line[1], line[2]
	d := b / a
	if d > c {
		return c
	}
	return d
}

func main() {
	input := lib_MustNewInputFromReader(bufio.NewReader(io.Reader(os.Stdin)))
	fmt.Println(solve(input))
}
