package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

const (
	A = 0
	B = 1
	C = 2
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
func solve(N int64, a []int64, b []int64, c []int64) int64 {
	dp := lib_New2DimInt64Slice(int(N)+1, 3)
	dp[0][A], dp[0][B], dp[0][C] = 0, 0, 0
	for i := 1; i <= int(N); i++ {
		dp[i][A] = lib_MustMaxInt64(dp[i-1][B]+b[i-1], dp[i-1][C]+c[i-1])
		dp[i][B] = lib_MustMaxInt64(dp[i-1][A]+a[i-1], dp[i-1][C]+c[i-1])
		dp[i][C] = lib_MustMaxInt64(dp[i-1][A]+a[i-1], dp[i-1][B]+b[i-1])
	}
	return lib_MustMaxInt64(dp[N][A], dp[N][B], dp[N][C])
}
