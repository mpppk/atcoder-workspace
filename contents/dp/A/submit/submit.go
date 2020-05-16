package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Int64Map map[ // Int64Map は、map[Int64][Int64]に便利メソッドを追加します.
int64]int64

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
func lib_NewInt64Map() Int64Map {
	return map[int64]int64{}
}
func (m Int64Map) ChMin(key, value int64, values ...int64) (replaced bool, valueAlreadyExist bool) {
	min, _ := lib_MinInt64(append(values, value)...)
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
func (m Int64Map) MustGet(key int64) int64 {
	v, ok := m[key]
	if !ok {
		panic(fmt.Sprintf("ivnalid key is specfied in Int64Map: %v", key))
	}
	return v
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
	m := lib_NewInt64Map()
	m[0], m[1] = 0, lib_AbsInt64(h[1]-h[0])
	for i := int64(2); i < N; i++ {
		m.ChMin(i, m[i-1]+lib_AbsInt64(h[i-1]-h[i]), m[i-2]+lib_AbsInt64(h[i-2]-h[i]))
	}
	return m.MustGet(N - 1)
}
