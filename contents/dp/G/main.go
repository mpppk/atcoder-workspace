package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/mpppk/atcoder-workspace/lib"
)

func solve(N int, M int, x []int, y []int) (ret int) {
	dp := lib.NewSIntList(N+5, 0)
	neighbors := lib.MustDirectedAdjacencyList(x, y, N)

	dp.Rec(func(n int) (res int) {
		for _, neighbor := range neighbors[n-1] {
			res = lib.MustMaxInt(res, dp.Update(neighbor)+1)
		}
		return res
	})
	for i := 1; i <= N; i++ {
		ret = lib.MustMaxInt(ret, dp.Update(i))
	}
	return ret
}

//func solve(N int, M int, x []int, y []int) int {
//	neighbors := make([][]int, N+5, N+5)
//	dp := lib.NewIntList(N+5, -1)
//
//	for i := 0; i < M; i++ {
//		curX, curY := x[i], y[i]
//		neighbors[curX] = append(neighbors[curX], curY)
//	}
//	ret := 0
//	for i := 1; i <= N; i++ {
//		ret = lib.MustMaxInt(ret, rec(dp, neighbors, i))
//	}
//	return ret
//}
//
//func rec(dp lib.IntList, neighbors [][]int, n int) int {
//	if dp[n] != -1 {
//		return dp[n]
//	}
//	dp[n] = 0
//	for _, neighbor := range neighbors[n] {
//		dp.ChMax(n, rec(dp, neighbors, neighbor)+1)
//	}
//	return dp[n]
//}

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
