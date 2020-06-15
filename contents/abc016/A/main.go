package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const YES = "YES"
const NO = "NO"

func solve(M int, D int) string {
	return ""
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var M int
	scanner.Scan()
	M, _ = strconv.Atoi(scanner.Text())
	var D int
	scanner.Scan()
	D, _ = strconv.Atoi(scanner.Text())
	fmt.Println(solve(M, D))
}
