package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/mpppk/atcoder-workspace/lib"
)

const (
	A = 0
	B = 1
	C = 2
)

func solve(N int, a []int, b []int, c []int) int {
	m := lib.NewInt2DMap()
	m.Set(0, A, 0)
	m.Set(0, B, 0)
	m.Set(0, C, 0)

	for i := 0; i < N; i++ {
		m.ChMax(i+1, A, m[i][B]+a[i], m[i][C]+a[i])
		m.ChMax(i+1, B, m[i][A]+b[i], m[i][C]+b[i])
		m.ChMax(i+1, C, m[i][A]+c[i], m[i][B]+c[i])
	}
	return lib.MustMaxInt(m[N].Values()...)
}

//func solve(N int, a []int, b []int, c []int) int {
//	dp := lib.New2DimInt64Slice(int(N)+1, 3)
//	dp[0][A], dp[0][B], dp[0][C] = 0, 0, 0
//
//	for i := 1; i <= int(N); i++ {
//		dp[i][A] = lib.MustMaxInt64(
//			dp[i-1][B]+b[i-1],
//			dp[i-1][C]+c[i-1],
//		)
//		dp[i][B] = lib.MustMaxInt64(
//			dp[i-1][A]+a[i-1],
//			dp[i-1][C]+c[i-1],
//		)
//		dp[i][C] = lib.MustMaxInt64(
//			dp[i-1][A]+a[i-1],
//			dp[i-1][B]+b[i-1],
//		)
//	}
//
//	return lib.MustMaxInt64(dp[N][A], dp[N][B], dp[N][C])
//}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var N int
	scanner.Scan()
	N = parseInt(scanner.Text())
	a := make([]int, N)
	b := make([]int, N)
	c := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		a[i] = parseInt(scanner.Text())
		scanner.Scan()
		b[i] = parseInt(scanner.Text())
		scanner.Scan()
		c[i] = parseInt(scanner.Text())
	}
	fmt.Println(solve(N, a, b, c))
}

func parseInt(s string) int {
	ret, _ := strconv.ParseInt(s, 10, 64)
	return int(ret)
}
