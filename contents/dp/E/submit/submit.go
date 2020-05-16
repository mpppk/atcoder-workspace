package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Int2DList [][]int
type IntList []int

func (a Int2DList) ChMin(i, j int, value int) bool {
	curV := a[i][j]
	if curV > value {
		a[i][j] = value
		return true
	}
	return false
}
func lib_NewInt2DList(length1, length2 int, initialValue int) Int2DList {
	ret := make([][]int, length1, length1)
	for i := 0; i < length1; i++ {
		ret[i] = lib_NewIntList(length2, initialValue)
	}
	return ret
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
	var W int
	scanner.Scan()
	W, _ = strconv.Atoi(scanner.Text())
	w := make([]int, N)
	v := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		w[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		v[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(solve(N, W, w, v))
}
func solve(N int, W int, w []int, v []int) int {
	maxV := N * 1000
	list := lib_NewInt2DList(N+10, maxV+10, math.MaxInt64/2)
	list[0][0] = 0
	for i := 0; i < N; i++ {
		for curV := 0; curV <= maxV; curV++ {
			list.ChMin(i+1, curV, list[i][curV])
			if newW := list[i][curV] + w[i]; newW <= W {
				list.ChMin(i+1, curV+v[i], newW)
			}
		}
	}
	retV := 0
	for curV, w := range list[N] {
		if curV > retV && w <= W {
			retV = curV
		}
	}
	return retV
}
