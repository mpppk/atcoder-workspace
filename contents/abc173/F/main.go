package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int, u []int, v []int) string {
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
	u := make([]int, N-1)
	v := make([]int, N-1)
	for i := 0; i < N-1; i++ {
		scanner.Scan()
		u[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		v[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(solve(N, u, v))
}
