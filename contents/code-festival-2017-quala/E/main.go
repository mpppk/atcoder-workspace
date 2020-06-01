package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MOD = 998244353

func solve(N int, M int, A string, B string, C string, D string) string {
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
	var M int
	scanner.Scan()
	M, _ = strconv.Atoi(scanner.Text())
	var A string
	scanner.Scan()
	A = scanner.Text()
	var B string
	scanner.Scan()
	B = scanner.Text()
	var C string
	scanner.Scan()
	C = scanner.Text()
	var D string
	scanner.Scan()
	D = scanner.Text()
	fmt.Println(solve(N, M, A, B, C, D))
}
