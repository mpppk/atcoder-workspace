package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int, w []int, s []int, v []int) string {
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
	w := make([]int, N)
	s := make([]int, N)
	v := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		w[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		s[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		v[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(solve(N, w, s, v))
}
