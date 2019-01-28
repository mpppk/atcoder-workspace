package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(input *utl_Input) int {
	cnt := 0
	nums := input.MustGetInt64Line(1)
	for {
		oddNums := utl_FilterInt64(nums, func(n int64) bool {
			return n%2 == 1
		})
		if len(oddNums) > 0 {
			return cnt
		}
		nums = utl_MapInt64(nums, func(n int64) int64 {
			return n / 2
		})
		cnt++
	}
}

func main() {
	input := utl_NewInput(bufio.NewScanner(os.Stdin))
	fmt.Println(solve(input))
}
