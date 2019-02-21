package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func solve(input *lib_Input) string {
	ticketStrs := input.MustGetFirstValue(0)
	ticketNumbers := lib_MustStringToIntSlice(ticketStrs)
	for _, bits := range lib_BitEnumeration(uint(len(ticketNumbers) - 1)) {
		sum := ticketNumbers[0]
		for i, bit := range bits {
			if bit {
				sum += ticketNumbers[i+1]
			} else {
				sum -= ticketNumbers[i+1]
			}
		}

		if sum == 7 {
			output := fmt.Sprintf("%d", ticketNumbers[0])
			for i, bit := range bits {
				if bit {
					output += "+"
				} else {
					output += "-"
				}
				output += fmt.Sprintf("%d", ticketNumbers[i+1])
			}
			return output + "=7"
		}
	}
	return "no answer"
}

func main() {
	input := lib_MustNewInputFromReader(bufio.NewReader(io.Reader(os.Stdin)))
	fmt.Println(solve(input))
}
