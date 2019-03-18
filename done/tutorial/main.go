package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(input *utl_Input) string {
	line1 := input.MustGetIntLine(0)
	line2 := input.MustGetIntLine(1)
	str := input.MustGetFirstValue(2)
	sum := utl_SumInt(append(line1, line2...))
	return fmt.Sprintf("%v %v", sum, str)
}

func main() {
	input := utl_NewInput(bufio.NewScanner(os.Stdin))
	fmt.Println(solve(input))
}
