package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/mpppk/atcoder-workspace/lib"
)

func solve(N int, W int, w []int, v []int) int {
	list := lib.NewInt2DList(N+100, W+100, 0)
	for i := 0; i < N; i++ {
		for curW := 0; curW <= W; curW++ {
			list.ChMax(i+1, curW, list[i][curW])
			if newW := curW + w[i]; newW <= W {
				list.ChMax(i+1, newW, list[i][curW]+v[i])
			}
		}
	}
	return lib.MustMaxInt(list[N]...)
}

//func solve(N int, W int, w []int, v []int) int {
//	m := lib.NewInt2DMap(N, W)
//	m.Set(0, 0, 0)
//	for i := 0; i < N; i++ {
//		for curW, beforeSumV := range m[i] {
//			m.ChMax(i+1, curW, beforeSumV)
//			if curW+w[i] <= W {
//				m.ChMax(i+1, curW+w[i], beforeSumV+v[i])
//			}
//		}
//	}
//	return lib.MustMaxInt(m[N].Values()...)
//}

//func solve(N int64, W int64, w []int64, v []int64) int64 {
//	dp := lib.New2DimInt64Slice(int(N)+1, int(W+1))
//	for i := 0; i < int(N); i++ {
//		for j := int64(0); j <= W; j++ {
//			newW := j + w[i]
//			if newW <= W {
//				dp[i+1][newW] = lib.MustMaxInt64(
//					dp[i+1][newW],
//					dp[i][j]+v[i],
//				)
//			}
//			dp[i+1][j] = lib.MustMaxInt64(
//				dp[i+1][j],
//				dp[i][j],
//			)
//		}
//	}
//
//	return lib.MustMaxInt64(dp[N]...)
//}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var N int
	scanner.Scan()
	tempN, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	N = int(tempN)
	var W int
	scanner.Scan()
	tempW, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	W = int(tempW)
	w := make([]int, N)
	v := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		tempw, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		w[i] = int(tempw)
		scanner.Scan()
		tempv, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		v[i] = int(tempv)
	}
	fmt.Println(solve(N, W, w, v))
}
