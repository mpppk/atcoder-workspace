package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(input *utl_Input) int {
	fiveHundredYenNum := input.MustGetFirstIntValue(0)
	hundredYenNum := input.MustGetLastIntValue(1)
	fiftyYenNum := input.MustGetLastIntValue(2)
	expectedSum := input.MustGetLastIntValue(3)

	cnt := 0
	for a := 0; a <= fiveHundredYenNum; a++ {
		for b := 0; b <= hundredYenNum; b++ {
			for c := 0; c <= fiftyYenNum; c++ {
				sum := 500*a + 100*b + 50*c
				if sum == expectedSum {
					cnt++
				}
			}
		}
	}
	return cnt
}

func main() {
	input := utl_NewInput(bufio.NewScanner(os.Stdin))
	fmt.Println(solve(input))
}
