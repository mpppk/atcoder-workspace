package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	NO  = "NO"
	YES = "YES"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var N int
	scanner.Scan()
	N, _ = strconv.Atoi(scanner.Text())
	fmt.Println(solve(N))
}
func solve(N int) string {
	if N%3 == 0 {
		return YES
	}
	return NO
}
