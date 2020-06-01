package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var defaultInitialValue = -999
var initialValueSInt int = defaultInitialValue
var recFuncSInt func(int) int = nil

type SIntList []int

func lib_DirectedAdjacencyList(x []int, y []int, length int) ([][]int, error) {
	if len(x) != len(y) {
		return nil, fmt.Errorf("x and y lengths are different. x:%d y:%d", len(x), len(y))
	}
	ret := make([][]int, length, length)
	for i := 0; i < len(x); i++ {
		ret[x[i]-1] = append(ret[x[i]-1], y[i])
	}
	return ret, nil
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
func lib_MustDirectedAdjacencyList(x []int, y []int, length int) [][]int {
	_v0, _err := lib_DirectedAdjacencyList(x, y, length)
	if _err != nil {
		panic(_err)
	}
	return _v0
}
func lib_MustMaxInt(values ...int) (max int) {
	max, err := lib_MaxInt(values...)
	if err != nil {
		panic(err)
	}
	return max
}
func lib_NewSIntList(length int, initialValue int) SIntList {
	if initialValueSInt != defaultInitialValue {
		panic("SIntList is already used")
	}
	ret := make([]int, length, length)
	for i := 0; i < length; i++ {
		ret[i] = initialValue
	}
	initialValueSInt = initialValue
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
func (si SIntList) Rec(f func(int) int) {
	recFuncSInt = f
}
func (si SIntList) Update(v int) int {
	if recFuncSInt == nil {
		panic("recFuncSInt is nil")
	}
	if si[v] != initialValueSInt {
		return si[v]
	}
	ret := recFuncSInt(v)
	si[v] = ret
	return ret
}
func solve(N int, M int, x []int, y []int) (ret int) {
	dp := lib_NewSIntList(N+5, 0)
	neighbors := lib_MustDirectedAdjacencyList(x, y, N)
	dp.Rec(func(n int) (res int) {
		for _, neighbor := range neighbors[n-1] {
			res = lib_MustMaxInt(res, dp.Update(neighbor)+1)
		}
		return res
	})
	for i := 1; i <= N; i++ {
		ret = lib_MustMaxInt(ret, dp.Update(i))
	}
	return ret
}
