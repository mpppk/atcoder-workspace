package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int, M int, Q int, L []int, R []int, p []int, q []int) string {
	return ""
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
	var Q int
	scanner.Scan()
	Q, _ = strconv.Atoi(scanner.Text())
	L := make([]int, M)
	R := make([]int, M)
	for i := 0; i < M; i++ {
		scanner.Scan()
		L[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		R[i], _ = strconv.Atoi(scanner.Text())
	}
	p := make([]int, Q)
	q := make([]int, Q)
	for i := 0; i < Q; i++ {
		scanner.Scan()
		p[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		q[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(solve(N, M, Q, L, R, p, q))
}
