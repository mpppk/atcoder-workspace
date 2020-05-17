package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int, M int, A []int, B []int, C []int, D []int, E []int, F []int) string {
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
	A := make([]int, N)
	B := make([]int, N)
	C := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		A[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		B[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		C[i], _ = strconv.Atoi(scanner.Text())
	}
	D := make([]int, M)
	E := make([]int, M)
	F := make([]int, M)
	for i := 0; i < M; i++ {
		scanner.Scan()
		D[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		E[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		F[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(solve(N, M, A, B, C, D, E, F))
}
