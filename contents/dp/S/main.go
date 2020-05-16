package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MOD = 1000000007

func solve(K string, D int) string {
	return ""
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var K string
	scanner.Scan()
	K = scanner.Text()
	var D int
	scanner.Scan()
	D, _ = strconv.Atoi(scanner.Text())
	fmt.Println(solve(K, D))
}
