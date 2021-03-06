package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/mpppk/atcoder-workspace/lib"
)

func solve(N int, h []int) int {
	m := lib.NewIntMap()
	m[0], m[1] = 0, lib.AbsInt(h[1]-h[0])
	for i := 2; i < N; i++ {
		m.ChMin(i, m[i-1]+lib.AbsInt(h[i-1]-h[i]), m[i-2]+lib.AbsInt(h[i-2]-h[i]))
	}
	return m.MustGet(N - 1)
}

//func solve(N int64, h []int64) int64 {
//	dp := lib.NewInt64SliceWithInitialValue(int(N), math.MaxInt64/2)
//	dp[0] = 0
//	dp[1] = lib.AbsInt64(h[1] - h[0])
//
//	for n := int64(2); n < N; n++ {
//		dp[n] = lib.MustMinInt64(
//			dp[n-1]+lib.AbsInt64(h[n]-h[n-1]),
//			dp[n-2]+lib.AbsInt64(h[n]-h[n-2]),
//		)
//	}
//	return dp[N-1]
//}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var N int64
	scanner.Scan()
	N, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	h := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		h[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	fmt.Println(solve(int(N), int(h)))
}
