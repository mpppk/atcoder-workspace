package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func solve(A int64, B int64) int64 {
	minBprice := int64(math.Trunc(float64(A)) / 0.1)

	for p := minBprice; ; p++ {
		if calcA(p) > A || calcB(p) > B {
			return -1
		}
		if calcA(p) == A && calcB(p) == B {
			return p
		}
	}
}

func calcA(p int64) int64 {
	return int64(math.Trunc(float64(p)) * 0.08)
}

func calcB(p int64) int64 {
	return int64(math.Trunc(float64(p)) * 0.1)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var A int64
	scanner.Scan()
	A, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var B int64
	scanner.Scan()
	B, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Println(solve(A, B))
}
