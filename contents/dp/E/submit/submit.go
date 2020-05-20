package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var IntMapCapForInt2DMap = 0

type Int2DMap map[ // IntMap は、map[Int][Int]に便利メソッドを追加します.
int]IntMap

type IntMap map[ // IntMap は、map[Int][Int]に便利メソッドを追加します.
int]int

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
func lib_MustMaxInt(values ...int) (max int) {
	max, err := lib_MaxInt(values...)
	if err != nil {
		panic(err)
	}
	return max
}
func lib_NewInt2DMap(cap, cap2 int) Int2DMap {
	IntMapCapForInt2DMap = cap2
	return make(map[int]IntMap, cap)
}
func lib_NewIntMap(cap int) IntMap {
	return make(map[int]int, cap)
}
func (m Int2DMap) ChMin(key1, key2, value int, values ...int) (replaced bool, valueAlreadyExist bool) {
	min, _ := lib_MinInt(append(values, value)...)
	if v, ok := m.Get(key1, key2); ok {
		if v > min {
			m.Set(key1, key2, min)
			return true, true
		} else {
			return false, true
		}
	}
	m.Set(key1, key2, min)
	return true, false
}
func (m Int2DMap) Get(key1, key2 int) (int, bool) {
	v1, ok := m[key1]
	if !ok {
		return 0, false
	}
	v2, ok := v1[key2]
	if !ok {
		return 0, false
	}
	return v2, true
}
func (m IntMap) Keys() (keys []int) {
	keys = make([]int, 0, len(m))
	for k, _ := range m {
		keys = append(keys, k)
	}
	return
}
func (m Int2DMap) Set(key1, key2, value int) (isNewValue bool) {
	v1, ok := m[key1]
	if !ok {
		m[key1] = lib_NewIntMap(IntMapCapForInt2DMap)
		v1 = m[key1]
	}
	_, ok = v1[key2]
	v1[key2] = value
	return !ok
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
	m := lib_NewInt2DMap(N, 1000*N)
	m.Set(0, 0, 0)
	for i, curV := range v {
		curW := w[i]
		for bV, bW := range m[i] {
			m.ChMin(i+1, bV, bW)
			if newW := bW + curW; newW <= W {
				m.ChMin(i+1, bV+curV, newW)
			}
		}
	}
	return lib_MustMaxInt(m[N].Keys()...)
}
