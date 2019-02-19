package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

func solve(input *generic_Input) float64 {
	points := input.MustGetIntLines()
	points = points[1:]

	maxLength := 0.0
	for _, c := range generic_MustIntSliceCombination(points, 2) {
		a := c[0]
		b := c[1]

		l := math.Pow(float64(a[0]-b[0]), 2) + math.Pow(float64(a[1]-b[1]), 2)
		l2 := math.Sqrt(l)

		if maxLength < l2 {
			maxLength = l2
		}
	}

	return maxLength
}

func main() {
	input := generic_MustNewInputFromReader(bufio.NewReader(io.Reader(os.Stdin)))
	fmt.Println(solve(input))
}
