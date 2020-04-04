package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(X int64) int64 {
	u1 := (X / 500) * 1000
	amari := X % 500
	return u1 + (amari/5)*5
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
	fmt.Println(solve(X))
}
