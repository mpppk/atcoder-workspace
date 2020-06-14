package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int, Q int, A []int, B []int, C []int, D []int) string {
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
	var Q int
	scanner.Scan()
	Q, _ = strconv.Atoi(scanner.Text())
	A := make([]int, N)
	B := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		A[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		B[i], _ = strconv.Atoi(scanner.Text())
	}
	C := make([]int, Q)
	D := make([]int, Q)
	for i := 0; i < Q; i++ {
		scanner.Scan()
		C[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		D[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(solve(N, Q, A, B, C, D))
}
