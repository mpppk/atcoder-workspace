package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(M []int, D []int) string {
	return ""
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	M := make([]int, 2)
	D := make([]int, 2)
	for i := 0; i < 2; i++ {
		scanner.Scan()
		M[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		D[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(solve(M, D))
}
