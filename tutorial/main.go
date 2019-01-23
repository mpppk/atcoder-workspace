package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(input *utl_Input) string {
	line1 := input.MustGetIntLine(0)
	line2 := input.MustGetIntLine(1)
	line3 := input.MustGetLine(2)
	str := line3[0]
	concatLine := append(line1, line2...)
	sum := 0
	for _, v := range concatLine {
		sum += v
	}

	return fmt.Sprintf("%v %v", sum, str)
}

func main() {
	input := utl_NewInput(bufio.NewScanner(os.Stdin))
	fmt.Println(solve(input))
}
