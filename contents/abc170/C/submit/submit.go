package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func lib_CountInt(values []int) map[int]int {
	m := map[int]int{}
	for _, value := range values {
		m[value]++
	}
	return m
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var X int
	scanner.Scan()
	X, _ = strconv.Atoi(scanner.Text())
	var N int
	scanner.Scan()
	N, _ = strconv.Atoi(scanner.Text())
	p := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		p[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(solve(X, N, p))
}
func solve(X int, N int, p []int) int {
	counts := lib_CountInt(p)
	for i := 0; ; i++ {
		if _, ok := counts[X-i]; !ok {
			return X - i
		}
		if _, ok := counts[X+i]; !ok {
			return X + i
		}
	}
}
