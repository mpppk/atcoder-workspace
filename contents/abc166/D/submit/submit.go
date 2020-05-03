package main

import (
	"bufio"
	"fmt"
	"math"
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
	solve(X)
}
func pow5(v int) int64 {
	return int64(math.Pow(float64(v), 5))
}
func pow9(v int) int64 {
	return int64(math.Pow(float64(v), 5))
}
func solve(X int64) {
	for i := 0; ; i++ {
		if pow5(i+1)-pow5(i) > pow9(10) {
			fmt.Println(i)
		}
	}
}
