package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, p []float64) string {
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
	p := make([]float64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		p[i], _ = strconv.ParseFloat(scanner.Text(), 64)
	}
	fmt.Println(solve(N, p))
}
