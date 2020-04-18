package main

import (
	"bufio"
	"fmt"
	"github.com/mpppk/atcoder/done/abc123/b"
	"io"
	"math"
	"os"
)

func solve(input *b.lib_Input) int {
	dishes := input.MustGetColIntLine(0)
	dishOrders := b.lib_IntPermutation(dishes, 5)
	minTime := math.MaxInt64 / 2
	for _, dishOrder := range dishOrders {
		time := 0
		for _, dish := range dishOrder {
			if time%10 != 0 {
				time += 10 - (time % 10)
			}
			time += dish
		}
		if minTime > time {
			minTime = time
		}
	}
	return minTime
}

func main() {
	input := b.lib_MustNewInputFromReader(bufio.NewReader(io.Reader(os.Stdin)))
	fmt.Println(solve(input))
}
