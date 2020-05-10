package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

func lib_CopyInt64(values []int64) []int64 {
	dst := make([]int64, len(values))
	copy(dst, values)
	return dst
}
func lib_MinInt64(values ...int64) (min int64, err error) {
	if len(values) == 0 {
		return 0, errors.New("empty slice is given")
	}
	min = values[0]
	for _, value := range values {
		if min > value {
			min = value
		}
	}
	return
}
func lib_MustMinInt64(values ...int64) (min int64) {
	min, err := lib_MinInt64(values...)
	if err != nil {
		panic(err)
	}
	return min
}
func lib_New2DimInt64SliceWithInitialValue(row, col int, initialValue int64) [][]int64 {
	ret := make([][]int64, row)
	values := lib_NewInt64SliceWithInitialValue(col, initialValue)
	for r := 0; r < row; r++ {
		ret[r] = lib_CopyInt64(values)
	}
	return ret
}
func lib_NewInt64SliceWithInitialValue(length int, initialValue int64) []int64 {
	ret := make([]int64, length)
	for i := 0; i < length; i++ {
		ret[i] = initialValue
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
	maxV := N * int64(1000)
	dp := lib_New2DimInt64SliceWithInitialValue(int(N)+1, int(maxV+1), math.MaxInt64/2)
	dp[0][0] = 0
	for i := 0; i < int(N); i++ {
		for j := int64(0); j <= maxV; j++ {
			if newV := j + v[i]; newV <= maxV {
				newW := dp[i][j] + w[i]
				dp[i+1][newV] = lib_MustMinInt64(dp[i+1][newV], newW)
			}
			dp[i+1][j] = lib_MustMinInt64(dp[i+1][j], dp[i][j])
		}
	}
	resV := int64(1)
	for i := int64(1); i <= maxV; i++ {
		if dp[N][i] <= W && resV < i {
			resV = i
		}
	}
	return resV
}
