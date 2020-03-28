package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func solve(input *lib_Input) int {
	nums := input.MustGetIntLine(0)
	n := nums[0]
	m := nums[1]
	nNum := 0
	if n > 1 {
		nNum = lib_MustCombination(n, 2)
	}
	mNum := 0
	if m > 1 {
		mNum = lib_MustCombination(m, 2)
	}

	return nNum + mNum
}

func main() {
	input := lib_MustNewInputFromReader(bufio.NewReader(io.Reader(os.Stdin)))
	fmt.Println(solve(input))
}
