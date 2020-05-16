package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/mpppk/atcoder-workspace/lib"
)

func solve(N int, K int, h []int) int {
	m := lib.NewIntMap()
	m[0] = 0
	for i := 0; i < N; i++ {
		for j := 1; j <= K; j++ {
			if i+j < len(h) {
				m.ChMin(i+j, m[i]+lib.AbsInt(h[i+j]-h[i]))
			}
		}
	}
	return m[N-1]
}

//func solve(N int64, K int64, h []int64) int64 {
//	dummyH := make([]int64, K)
//	h = append(h, dummyH...)
//
//	dp := lib.NewInt64SliceWithInitialValue(int(N+K), math.MaxInt64/2)
//	dp[0] = 0
//	for i := 0; i < int(N); i++ {
//		for j := 1; j <= int(K); j++ {
//			dp[i+j] = lib.MustMinInt64(
//				dp[i+j],
//				dp[i]+lib.AbsInt64(h[i]-h[i+j]),
//			)
//		}
//	}
//
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
	var K int64
	scanner.Scan()
	K, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	h := make([]int, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		t, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		h[i] = int(t)
	}
	fmt.Println(solve(int(N), int(K), h))
}
