package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func lib_PowInt64(x, y int64) int64 {
	return int64(math.Pow(float64(x), float64(y)))
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var X int64
	scanner.Scan()
	X, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(X)
}
func solve(X int64) {
	for A := int64(0); A < 120; A++ {
		for B := int64(-120); B < 120; B++ {
			if lib_PowInt64(A, 5)-lib_PowInt64(B, 5) == X {
				fmt.Println(A, B)
				return
			}
		}
	}
}
