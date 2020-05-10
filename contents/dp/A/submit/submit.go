package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

func lib_AbsInt64(value int64) int64 {
	return int64(math.Abs(float64(value)))
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
	h := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		h[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	fmt.Println(solve(N, h))
}
func solve(N int64, h []int64) int64 {
	dp := lib_NewInt64SliceWithInitialValue(int(N), math.MaxInt64/2)
	dp[0] = 0
	dp[1] = lib_AbsInt64(h[1] - h[0])
	for n := int64(2); n < N; n++ {
		dp[n] = lib_MustMinInt64(dp[n-1]+lib_AbsInt64(h[n]-h[n-1]), dp[n-2]+lib_AbsInt64(h[n]-h[n-2]))
	}
	return dp[N-1]
}
