package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(X int64, Y int64, Z int64) {
	fmt.Println(Z, X, Y)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var X int64
	scanner.Scan()
	X, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var Y int64
	scanner.Scan()
	Y, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var Z int64
	scanner.Scan()
	Z, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(X, Y, Z)
}
