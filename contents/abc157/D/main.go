package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, M int64, K int64, A []int64, B []int64, C []int64, D []int64) string {
	return ""
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
	var M int64
	scanner.Scan()
	M, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var K int64
	scanner.Scan()
	K, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	A := make([]int64, M)
	B := make([]int64, M)
	for i := int64(0); i < M; i++ {
		scanner.Scan()
		A[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		B[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	C := make([]int64, K)
	D := make([]int64, K)
	for i := int64(0); i < K; i++ {
		scanner.Scan()
		C[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		D[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	fmt.Println(solve(N, M, K, A, B, C, D))
}
