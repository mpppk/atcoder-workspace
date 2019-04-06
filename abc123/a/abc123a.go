package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

func solve(input *lib_Input) string {
	antennas := input.MustGetColIntLine(0)
	k := antennas[len(antennas)-1]
	antennas = antennas[:len(antennas)-1]
	for i := 0; i < len(antennas); i++ {
		for j := 0; j < len(antennas); j++ {
			if int(math.Abs(float64(antennas[i]-antennas[j]))) > k {
				return ":("
			}
		}
	}
	return "Yay!"
}

func main() {
	input := lib_MustNewInputFromReader(bufio.NewReader(io.Reader(os.Stdin)))
	fmt.Println(solve(input))
}
