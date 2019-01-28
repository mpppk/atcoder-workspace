package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(input *utl_Input) int {
	ss := input.MustGetFirstValue(0)
	cnt := 0
	for _, s := range ss {
		if s == '1' {
			cnt++
		}
	}

	return cnt
}

func main() {
	input := utl_NewInput(bufio.NewScanner(os.Stdin))
	fmt.Println(solve(input))
}
