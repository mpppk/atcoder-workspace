package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func lib_IntCombination(values []int, r int) (combinations [][]int, err error) {
	if r == 1 {
		for _, value := range values {
			combinations = append(combinations, []int{value})
		}
		return
	}
	for i := range values {
		newValues := values
		for j := 0; j <= i; j++ {
			newValues = newValues[1:]
		}
		partialCombinations, err := lib_IntCombination(newValues, r-1)
		if err != nil {
			return nil, err
		}
		for _, pc := range partialCombinations {
			newC := append(pc, values[i])
			combinations = append(combinations, newC)
		}
	}
	return
}
func lib_IntRange(start, end, step int) ([]int, error) {
	if end < start {
		return nil, fmt.Errorf("end(%v) is bigger than start(%v)", end, start)
	}
	s := make([]int, 0, int(1+(end-start)/step))
	for start < end {
		s = append(s, start)
		start += step
	}
	return s, nil
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
func lib_MustIntCombination(values []int, r int) (combinations [][]int) {
	combinations, err := lib_IntCombination(values, r)
	if err != nil {
		panic(err)
	}
	return combinations
}
func lib_MustIntRange(start, end, step int) []int {
	_v0, _err := lib_IntRange(start, end, step)
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
	A := make([][]int, N)
	for i := 0; i < N; i++ {
		A[i] = make([]int, M)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			scanner.Scan()
			A[i][j], _ = strconv.Atoi(scanner.Text())
		}
	}
	fmt.Println(solve(N, M, A))
}
func solve(N int, M int, A [][]int) int {
	max := 0
	for _, pair := range lib_MustIntCombination(lib_MustIntRange(1, M+1, 1), 2) {
		sum := 0
		for i := 0; i < N; i++ {
			t1, t2 := pair[0]-1, pair[1]-1
			sum += lib_MustMaxInt(A[i][t1], A[i][t2])
		}
		max = lib_MustMaxInt(max, sum)
	}
	return max
}
