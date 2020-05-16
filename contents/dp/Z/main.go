package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int, C int, h []int) string {
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
	var C int
	scanner.Scan()
	C, _ = strconv.Atoi(scanner.Text())
	h := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		h[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(solve(N, C, h))
}
