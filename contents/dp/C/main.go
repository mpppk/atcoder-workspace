package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/mpppk/atcoder-workspace/lib"
)

const (
	B = 1
	C = 2
	A = 0
)

func solve(N int64, a []int64, b []int64, c []int64) int64 {
	dp := lib.New2DimInt64Slice(int(N)+1, 3)
	dp[0][A], dp[0][B], dp[0][C] = 0, 0, 0

	for i := 1; i <= int(N); i++ {
		dp[i][A] = lib.MustMaxInt64(
			dp[i-1][B]+b[i-1],
			dp[i-1][C]+c[i-1],
		)
		dp[i][B] = lib.MustMaxInt64(
			dp[i-1][A]+a[i-1],
			dp[i-1][C]+c[i-1],
		)
		dp[i][C] = lib.MustMaxInt64(
			dp[i-1][A]+a[i-1],
			dp[i-1][B]+b[i-1],
		)
	}

	return lib.MustMaxInt64(dp[N][A], dp[N][B], dp[N][C])
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
	a := make([]int64, N)
	b := make([]int64, N)
	c := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		a[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		b[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		c[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	fmt.Println(solve(N, a, b, c))
}
