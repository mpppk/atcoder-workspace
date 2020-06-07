package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

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
func lib_MeanFloat64(values []float64) float64 {
	sum := lib_SumFloat64(values)
	return sum / float64(len(values))
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
func lib_MustMinInt(values ...int) (min int) {
	min, err := lib_MinInt(values...)
	if err != nil {
		panic(err)
	}
	return min
}
func lib_SumFloat64(values []float64) float64 {
	var sum float64 = 0
	for _, value := range values {
		sum += value
	}
	return sum
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
	var A int
	scanner.Scan()
	A, _ = strconv.Atoi(scanner.Text())
	var B int
	scanner.Scan()
	B, _ = strconv.Atoi(scanner.Text())
	S := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		S[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(solve(N, A, B, S))
}
func solve(N int, A int, B int, S []int) string {
	min, max := lib_MustMinInt(S...), lib_MustMaxInt(S...)
	if min == max {
		return "-1"
	}
	P := float64(B) / float64(max-min)
	PS := make([]float64, 0, len(S))
	for _, s := range S {
		PS = append(PS, P*float64(s))
	}
	mean := lib_MeanFloat64(PS)
	diff := mean - float64(A)
	return fmt.Sprintf("%v %v", P, -diff)
}
