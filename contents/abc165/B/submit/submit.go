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
	var X int64
	scanner.Scan()
	X, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Println(solve(X))
}
func solve(X int64) int {
	money := int64(100)
	year := 0
	for {
		if money >= X {
			return year
		}
		money += int64(float64(money) * 0.01)
		year++
	}
}
