package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(P int, Q int, R int) string {
	return ""
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var P int
	scanner.Scan()
	P, _ = strconv.Atoi(scanner.Text())
	var Q int
	scanner.Scan()
	Q, _ = strconv.Atoi(scanner.Text())
	var R int
	scanner.Scan()
	R, _ = strconv.Atoi(scanner.Text())
	fmt.Println(solve(P, Q, R))
}
