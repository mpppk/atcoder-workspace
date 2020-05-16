package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

type IntMap map[ // IntMap は、map[Int][Int]に便利メソッドを追加します.
int]int

func lib_AbsInt(value int) int {
	return int(math.Abs(float64(value)))
}
func lib_MinInt(values ...int) (min int, err error) {
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
func lib_NewIntMap() IntMap {
	return map[int]int{}
}
func (m IntMap) ChMin(key, value int, values ...int) (replaced bool, valueAlreadyExist bool) {
	min, _ := lib_MinInt(append(values, value)...)
	if v, ok := m[key]; ok {
		if v > min {
			m[key] = min
			return true, true
		} else {
			return false, true
		}
	}
	m[key] = min
	return true, false
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
	h := make([]int, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		t, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		h[i] = int(t)
	}
	fmt.Println(solve(int(N), int(K), h))
}
func solve(N int, K int, h []int) int {
	m := lib_NewIntMap()
	m[0] = 0
	for i := 0; i < N; i++ {
		for j := 1; j <= K; j++ {
			if i+j < len(h) {
				m.ChMin(i+j, m[i]+lib_AbsInt(h[i+j]-h[i]))
			}
		}
	}
	return m[N-1]
}
