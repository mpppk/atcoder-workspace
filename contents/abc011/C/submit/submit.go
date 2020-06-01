package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const (
	NO  = "NO"
	YES = "YES"
)

func lib_FilterInt(values []int, f func(v int) bool) (newValues []int) {
	for _, value := range values {
		if f(value) {
			newValues = append(newValues, value)
		}
	}
	return
}
func lib_HasInt(values []int, v int) bool {
	for _, value := range values {
		if value == v {
			return true
		}
	}
	return false
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var N int
	scanner.Scan()
	N, _ = strconv.Atoi(scanner.Text())
	NG := make([]int, 3)
	for i := 0; i < 3; i++ {
		scanner.Scan()
		NG[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(solve(N, NG))
}
func solve(N int, NG []int) string {
	sort.Ints(NG)
	if lib_HasInt(NG, N) {
		return NO
	}
	NG2 := lib_FilterInt(NG, func(v int) bool {
		return v < N
	})
	if len(NG2) == 3 && NG[0]+2 == NG[1]+1 && NG[1]+1 == NG[2] {
		return NO
	}
	cnt := 0
	for i := len(NG2) - 1; i >= 0; i-- {
		if (NG2[i]-N%3-cnt)%3 == 0 {
			cnt++
		}
	}
	if (300 - N) < cnt {
		return NO
	}
	return YES
}
