package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(input *utl_Input) string {
	line1, err := input.GetIntLine(0)
	utl_PanicIfErrorExist(err)
	line2, err := input.GetIntLine(1)
	utl_PanicIfErrorExist(err)
	line3, err := input.GetLine(2)
	utl_PanicIfErrorExist(err)
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
