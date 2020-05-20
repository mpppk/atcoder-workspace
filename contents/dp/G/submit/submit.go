package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type IntList []int

func (a IntList) ChMax(i int, value int) bool {
	curV := a[i]
	if curV < value {
		a[i] = value
		return true
	}
	return false
}
func lib_MaxInt(values ...int) (max int, err error) {
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
func lib_MustMaxInt(values ...int) (max int) {
	max, err := lib_MaxInt(values...)
	if err != nil {
		panic(err)
	}
	return max
}
func lib_NewIntList(length int, initialValue int) IntList {
	ret := make([]int, length, length)
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
func rec(dp IntList, neighbors [][]int, n int) int {
	if dp[n] != -1 {
		return dp[n]
	}
	dp[n] = 0
	for _, neighbor := range neighbors[n] {
		dp.ChMax(n, rec(dp, neighbors, neighbor)+1)
	}
	return dp[n]
}
func solve(N int, M int, x []int, y []int) int {
	neighbors := make([][]int, N+5, N+5)
	dp := lib_NewIntList(N+5, -1)
	for i := 0; i < M; i++ {
		curX, curY := x[i], y[i]
		neighbors[curX] = append(neighbors[curX], curY)
	}
	ret := 0
	for i := 1; i <= N; i++ {
		ret = lib_MustMaxInt(ret, rec(dp, neighbors, i))
	}
	return ret
}
