package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const YES = "Yes"
const NO = "No"

func solve(A [][]int64, N int64, b []int64) string {
	return ""
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	A := make([][]int64, 3)
	for i := int64(0); i < 3; i++ {
		A[i] = make([]int64, 3)
	}
	for i := int64(0); i < 3; i++ {
		for j := int64(0); j < 3; j++ {
			scanner.Scan()
			A[i][j], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		}
	}
	var N int64
	scanner.Scan()
	N, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	b := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		b[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	fmt.Println(solve(A, N, b))
}
