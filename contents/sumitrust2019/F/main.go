package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(T []int, A []int, B []int) string {
	return ""
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	T := make([]int, 2)
	for i := 0; i < 2; i++ {
		scanner.Scan()
		T[i], _ = strconv.Atoi(scanner.Text())
	}
	A := make([]int, 2)
	for i := 0; i < 2; i++ {
		scanner.Scan()
		A[i], _ = strconv.Atoi(scanner.Text())
	}
	B := make([]int, 2)
	for i := 0; i < 2; i++ {
		scanner.Scan()
		B[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(solve(T, A, B))
}
