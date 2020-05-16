package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MOD = 1000000007

func solve(H int, W int, N int, r []int, c []int) string {
	return ""
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var H int
	scanner.Scan()
	H, _ = strconv.Atoi(scanner.Text())
	var W int
	scanner.Scan()
	W, _ = strconv.Atoi(scanner.Text())
	var N int
	scanner.Scan()
	N, _ = strconv.Atoi(scanner.Text())
	r := make([]int, N)
	c := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		r[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		c[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(solve(H, W, N, r, c))
}
