package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(L int64) float64 {
	h := float64(L) / 3
	return h * h * h
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var L int64
	scanner.Scan()
	L, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Println(solve(L))
}
