package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var n int64
	scanner.Scan()
	n, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Println(solve(n))
}
func solve(n int64) int64 {
	return n - 1
}
