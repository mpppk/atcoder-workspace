package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	x := make([]int, 5)
	for i := 0; i < 5; i++ {
		scanner.Scan()
		x[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(solve(x))
}
func solve(x []int) int {
	for i, xi := range x {
		if xi == 0 {
			return i + 1
		}
	}
	panic("")
}
