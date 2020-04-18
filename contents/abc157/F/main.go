package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, K int64, x []int64, y []int64, c []int64) string {
	return ""
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var N int64
	scanner.Scan()
	N, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var K int64
	scanner.Scan()
	K, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	x := make([]int64, N)
	y := make([]int64, N)
	c := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		x[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		y[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		c[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	fmt.Println(solve(N, K, x, y, c))
}
