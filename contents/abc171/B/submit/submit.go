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
	p := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		p[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(solve(N, K, p))
}
func solve(N int, K int, p []int) (sum int) {
	sort.Ints(p)
	for i := 0; i < K; i++ {
		sum += p[i]
	}
	return sum
}
