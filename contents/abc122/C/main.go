package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int, Q int, S string, l []int, r []int) string {
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
	var Q int
	scanner.Scan()
	Q, _ = strconv.Atoi(scanner.Text())
	var S string
	scanner.Scan()
	S = scanner.Text()
	l := make([]int, Q)
	r := make([]int, Q)
	for i := 0; i < Q; i++ {
		scanner.Scan()
		l[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		r[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(solve(N, Q, S, l, r))
}
