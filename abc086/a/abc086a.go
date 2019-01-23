package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(input *utl_Input) string {
	a := input.MustGetFirstIntValue(0)
	b := input.MustGetLastIntValue(0)
	c := a * b
	if c%2 == 0 {
		return "Even"
	}
	return "Odd"
}

func main() {
	input := utl_NewInput(bufio.NewScanner(os.Stdin))
	fmt.Println(solve(input))
}
