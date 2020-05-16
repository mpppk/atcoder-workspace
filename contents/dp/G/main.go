package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int, M int, x []int, y []int) string {
	return ""
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
	var M int
	scanner.Scan()
	M, _ = strconv.Atoi(scanner.Text())
	x := make([]int, M)
	y := make([]int, M)
	for i := 0; i < M; i++ {
		scanner.Scan()
		x[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		y[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(solve(N, M, x, y))
}
