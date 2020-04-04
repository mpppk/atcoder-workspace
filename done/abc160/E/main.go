package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(X int64, Y int64, A int64, B int64, C int64, p []int64, q []int64, r []int64) string {
	return ""
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
	var A int64
	scanner.Scan()
	A, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var B int64
	scanner.Scan()
	B, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var C int64
	scanner.Scan()
	C, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	p := make([]int64, A)
	for i := int64(0); i < A; i++ {
		scanner.Scan()
		p[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	q := make([]int64, B)
	for i := int64(0); i < B; i++ {
		scanner.Scan()
		q[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	r := make([]int64, C)
	for i := int64(0); i < C; i++ {
		scanner.Scan()
		r[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	fmt.Println(solve(X, Y, A, B, C, p, q, r))
}
