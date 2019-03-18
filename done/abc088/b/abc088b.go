package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func solve(input *utl_Input) int {
	nums := input.MustGetIntLine(1)
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	alice := 0
	bob := 0
	for i, num := range nums {
		if i%2 == 0 {
			alice += num
		} else {
			bob += num
		}
	}
	return alice - bob
}

func main() {
	input := utl_NewInput(bufio.NewScanner(os.Stdin))
	fmt.Println(solve(input))
}
