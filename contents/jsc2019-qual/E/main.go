package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int, H int, W int, R []int, C []int, A []int) string {
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
	var H int
	scanner.Scan()
	H, _ = strconv.Atoi(scanner.Text())
	var W int
	scanner.Scan()
	W, _ = strconv.Atoi(scanner.Text())
	R := make([]int, N)
	C := make([]int, N)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		R[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		C[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		A[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(solve(N, H, W, R, C, A))
}
