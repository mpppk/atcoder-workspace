package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(input *utl_Input) int {
	maxNum := input.MustGetFirstIntValue(0)
	minSum := input.MustGetIntValue(0, 1)
	maxSum := input.MustGetIntValue(0, 2)

	sum := 0
	for n := 0; n <= maxNum; n++ {
		ns := utl_GetEachDigitSumInt(n)
		if ns <= maxSum && ns >= minSum {
			sum += n
		}
	}
	return sum
}

func main() {
	input := utl_NewInput(bufio.NewScanner(os.Stdin))
	fmt.Println(solve(input))
}
