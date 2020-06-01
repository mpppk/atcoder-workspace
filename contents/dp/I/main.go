package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/mpppk/atcoder-workspace/lib"
)

func solve(N int, p []float64) float64 {
	dp := lib.NewFloat642DList(N+5, N+5, 0)
	dp[0][0] = 1

	for i := 0; i < N; i++ {
		for j := 0; j <= i; j++ {
			dp[i+1][j+1] += dp[i][j] * p[i]
			dp[i+1][j] += dp[i][j] * (1 - p[i])
		}
	}

	ret := 0.0

	for j := N; j > N/2; j-- {
		ret += dp[N][j]
	}

	return ret
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
	p := make([]float64, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		p[i], _ = strconv.ParseFloat(scanner.Text(), 64)
	}
	fmt.Println(solve(N, p))
}
