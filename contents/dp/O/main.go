package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MOD = 1000000007

func solve(N int, a [][]int) string {
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
	a := make([][]int, N)
	for i := 0; i < N; i++ {
		a[i] = make([]int, N)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			scanner.Scan()
			a[i][j], _ = strconv.Atoi(scanner.Text())
		}
	}
	fmt.Println(solve(N, a))
}
