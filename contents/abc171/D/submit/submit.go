package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func lib_CountIntAsList(values []int, max int) []int {
	m := make([]int, max, max)
	for _, value := range values {
		m[value]++
	}
	return m
}
func lib_Pow10Int(y int) int {
	return int(math.Pow10(int(y)))
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
	A := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		A[i], _ = strconv.Atoi(scanner.Text())
	}
	var Q int
	scanner.Scan()
	Q, _ = strconv.Atoi(scanner.Text())
	B := make([]int, Q)
	C := make([]int, Q)
	for i := 0; i < Q; i++ {
		scanner.Scan()
		B[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		C[i], _ = strconv.Atoi(scanner.Text())
	}
	solve(N, A, Q, B, C)
}
func solve(N int, A []int, Q int, B []int, C []int) {
	counter := lib_CountIntAsList(A, lib_Pow10Int(5)+1)
	cnt := 0
	for i, c := range counter {
		cnt += i * c
	}
	for i := 0; i < Q; i++ {
		bCnt := counter[B[i]]
		cnt -= bCnt * B[i]
		cnt += bCnt * C[i]
		counter[B[i]] = 0
		counter[C[i]] += bCnt
		fmt.Println(cnt)
	}
}
