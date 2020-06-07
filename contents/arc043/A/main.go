package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int, A int, B int, S []int) string {
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
	var A int
	scanner.Scan()
	A, _ = strconv.Atoi(scanner.Text())
	var B int
	scanner.Scan()
	B, _ = strconv.Atoi(scanner.Text())
	S := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		S[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(solve(N, A, B, S))
}
