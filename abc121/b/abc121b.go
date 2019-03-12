package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func solve(input *lib_Input) int {
	c := input.MustGetIntValue(0, 2)
	bList := input.MustGetIntLine(1)
	aLists := input.MustGetIntLinesFrom(2)

	cnt := 0
	for _, aList := range aLists {
		sum := 0
		for i, a := range aList {
			sum += a * bList[i]
		}
		if sum+c > 0 {
			cnt++
		}
	}
	return cnt
}

func main() {
	input := lib_MustNewInputFromReader(bufio.NewReader(io.Reader(os.Stdin)))
	fmt.Println(solve(input))
}
