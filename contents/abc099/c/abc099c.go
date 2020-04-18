package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var cache = map[int]int{}

func solve(input *lib_Input) int {
	N := input.MustGetFirstIntValue(0)
	return calcMinDrawerNum(N)
}

func calcMinDrawerNum(balance int) int {
	if cachedResult, ok := cache[balance]; ok {
		return cachedResult
	}
	drawerAmounts := []int{1, 6, 36, 216, 1296, 7776, 46656, 9, 81, 729, 6561, 59049}
	var drawerNums []int
	for _, drawerAmount := range drawerAmounts {
		newBalance := balance - drawerAmount
		if newBalance == 0 {
			return 1
		}
		if newBalance < 0 {
			continue
		}

		drawerNums = append(drawerNums, calcMinDrawerNum(newBalance))
	}
	result := lib_MustMinInt(drawerNums) + 1
	cache[balance] = result
	return result
}

func main() {
	input := lib_MustNewInputFromReader(bufio.NewReader(io.Reader(os.Stdin)))
	fmt.Println(solve(input))
}
