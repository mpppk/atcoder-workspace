package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const YES = "Yes"
const NO = "No"

func solve(N int, M int, K int) string {
	calcK := func(n, m int) int {
		return n*M + m*N - 2*n*m
	}

	for n := 0; n <= N; n++ {
		for m := 0; m <= M; m++ {
			if calcK(n, m) == K {
				return YES
			}
		}
	}
	return NO
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
	var K int
	scanner.Scan()
	K, _ = strconv.Atoi(scanner.Text())
	fmt.Println(solve(N, M, K))
}
