package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/mpppk/atcoder-workspace/lib"
)

const MOD = 10007

func solve(n int) int16 {
	dp := lib.NewSInt16List(n+5, 0)
	dp[2], dp[3] = 1, 1
	for i := 4; i < n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % MOD
		dp[i] += dp[i-3] % MOD
		dp[i] %= MOD
	}
	return dp[n-1]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var n int
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	fmt.Println(solve(n))
}
