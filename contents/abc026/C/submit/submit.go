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
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var N int
	scanner.Scan()
	N, _ = strconv.Atoi(scanner.Text())
	B := make([]int, N-2+1)
	for i := 0; i < N-2+1; i++ {
		scanner.Scan()
		B[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(solve(N, B))
}
func rec(adj map[int][]int, subs []int) (salary int) {
	var salaries []int
	if len(subs) == 0 {
		return 1
	}
	for _, sub := range subs {
		salaries = append(salaries, rec(adj, adj[sub]))
	}
	return lib_MustMaxInt(salaries...) + lib_MustMinInt(salaries...) + 1
}
func solve(N int, B []int) int {
	adj := make(map[int][]int, N)
	for i, boss := range B {
		adj[boss] = append(adj[boss], i+2)
	}
	return rec(adj, adj[1])
}
