package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Int2DList [][]int
type IntList []int

func (a Int2DList) ChMax(i, j int, value int) bool {
	curV := a[i][j]
	if curV < value {
		a[i][j] = value
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
func solve(N int, W int, w []int, v []int) int {
	list := lib_NewInt2DList(N+100, W+100, 0)
	for i := 0; i < N; i++ {
		for curW := 0; curW <= W; curW++ {
			list.ChMax(i+1, curW, list[i][curW])
			if newW := curW + w[i]; newW <= W {
				list.ChMax(i+1, newW, list[i][curW]+v[i])
			}
		}
	}
	return lib_MustMaxInt(list[N]...)
}
