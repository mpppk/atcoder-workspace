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
	var K int64
	scanner.Scan()
	K, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	h := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		h[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	fmt.Println(solve(N, K, h))
}
func solve(N int64, K int64, h []int64) int64 {
	dummyH := make([]int64, K)
	h = append(h, dummyH...)
	dp := lib_NewInt64SliceWithInitialValue(int(N+K), math.MaxInt64/2)
	dp[0] = 0
	for i := 0; i < int(N); i++ {
		for j := 1; j <= int(K); j++ {
			dp[i+j] = lib_MustMinInt64(dp[i+j], dp[i]+lib_AbsInt64(h[i]-h[i+j]))
		}
	}
	return dp[N-1]
}
