package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/mpppk/atcoder-workspace/lib"
)

func solve(N int64, K int64, h []int64) int64 {
	dummyH := make([]int64, K)
	h = append(h, dummyH...)

	dp := lib.NewInt64SliceWithInitialValue(int(N+K), math.MaxInt64/2)
	dp[0] = 0
	for i := 0; i < int(N); i++ {
		for j := 1; j <= int(K); j++ {
			dp[i+j] = lib.MustMinInt64(
				dp[i+j],
				dp[i]+lib.AbsInt64(h[i]-h[i+j]),
			)
		}
	}

	return dp[N-1]
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
	h := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		h[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	fmt.Println(solve(N, K, h))
}
