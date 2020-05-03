package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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
	H := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		H[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	A := make([]int64, M)
	B := make([]int64, M)
	for i := int64(0); i < M; i++ {
		scanner.Scan()
		A[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		B[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	fmt.Println(solve(N, M, H, A, B))
}
func solve(N int64, M int64, H []int64, A []int64, B []int64) int {
	m := map[int64][]int64{}
	for i := int64(1); i <= N; i++ {
		m[i] = []int64{}
	}
	for i, a := range A {
		b := B[i]
		m[a] = append(m[a], b)
		m[b] = append(m[b], a)
	}
	good := 0
	for curN, neighbors := range m {
		curH := H[curN-1]
		isGood := true
		for _, neighbor := range neighbors {
			if H[neighbor-1] >= curH {
				isGood = false
				break
			}
		}
		if isGood {
			good++
		}
	}
	return good
}
