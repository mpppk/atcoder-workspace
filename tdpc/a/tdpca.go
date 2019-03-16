package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func solve(input *lib_Input) int {
	pList := input.MustGetIntLine(1)
	sum := lib_MemoizeIntToInt(func(v1, v2 int) int {
		return v1 + v2
	})
	return len(listSum(pList[0], pList[1:], sum))
}

func listSum(currentValue int, otherList []int, sum func(v1, v2 int) int) []int {
	if len(otherList) == 0 {
		return []int{0, currentValue}
	}
	sums := listSum(otherList[0], otherList[1:], sum)
	addedSums := lib_MapInt(sums, func(s int) int {
		return sum(currentValue, s)
	})
	return lib_UniqInt(append(sums, addedSums...))
}

func main() {
	input := lib_MustNewInputFromReader(bufio.NewReader(io.Reader(os.Stdin)))
	fmt.Println(solve(input))
}
