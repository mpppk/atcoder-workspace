package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func lib_MaxInt64(values ...int64) (max int64, err error) {
	if len(values) == 0 {
		return 0, errors.New("empty slice is given")
	}
	max = values[0]
	for _, value := range values {
		if max < value {
			max = value
		}
	}
	return
}
func lib_MustMaxInt64(values ...int64) (max int64) {
	max, err := lib_MaxInt64(values...)
	if err != nil {
		panic(err)
	}
	return max
}
func lib_New2DimInt64Slice(row, col int) [][]int64 {
	ret := make([][]int64, row)
	for r := 0; r < row; r++ {
		ret[r] = make([]int64, col, col)
	}
	return ret
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
	var W int64
	scanner.Scan()
	W, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	w := make([]int64, N)
	v := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		w[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		v[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	fmt.Println(solve(N, W, w, v))
}
func solve(N int64, W int64, w []int64, v []int64) int64 {
	dp := lib_New2DimInt64Slice(int(N)+1, int(W+1))
	for i := 0; i < int(N); i++ {
		for j := int64(0); j <= W; j++ {
			newW := j + w[i]
			if newW <= W {
				dp[i+1][newW] = lib_MustMaxInt64(dp[i+1][newW], dp[i][j]+v[i])
			}
			dp[i+1][j] = lib_MustMaxInt64(dp[i+1][j], dp[i][j])
		}
	}
	return lib_MustMaxInt64(dp[N]...)
}
