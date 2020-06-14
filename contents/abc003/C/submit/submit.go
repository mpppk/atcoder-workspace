package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var N int
	scanner.Scan()
	N, _ = strconv.Atoi(scanner.Text())
	var K int
	scanner.Scan()
	K, _ = strconv.Atoi(scanner.Text())
	R := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		R[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(solve(N, K, R))
}
func solve(N int, K int, R []int) (rate float64) {
	sort.Ints(R)
	for i := N - K; i < N; i++ {
		rate = (rate + float64(R[i])) / 2.0
	}
	return rate
}
