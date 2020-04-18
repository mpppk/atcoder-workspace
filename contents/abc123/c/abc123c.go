package main

import (
	"bufio"
	"fmt"
	"github.com/mpppk/atcoder/done/abc123/c"
	"io"
	"math"
	"os"
)

func solve(input *c.lib_Input) int64 {
	inputs := input.MustGetColInt64Line(0)
	N := inputs[0]
	transportations := inputs[1:]

	minTime := transportations[0]

	for _, transportation := range transportations {
		if minTime > transportation {
			minTime = transportation
		}
	}

	ret := int64(math.Ceil(float64(N) / float64(minTime)))
	return ret + 4
}

func main() {
	input := c.lib_MustNewInputFromReader(bufio.NewReader(io.Reader(os.Stdin)))
	fmt.Println(solve(input))
}
