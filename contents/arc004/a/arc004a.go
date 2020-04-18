package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

func solve(input *lib_Input) float64 {
	coordinates := input.MustGetIntLinesFrom(1)
	com := lib_MustIntSliceCombination(coordinates, 2)
	return lib_MustMaxFloat64ByIntSlice2(com, func(c [][]int) float64 {
		x := c[0]
		y := c[1]
		l := math.Pow(float64(x[0]-y[0]), 2) + math.Pow(float64(x[1]-y[1]), 2)
		return math.Sqrt(l)
	})
}

func main() {
	input := lib_MustNewInputFromReader(bufio.NewReader(io.Reader(os.Stdin)))
	fmt.Println(solve(input))
}
