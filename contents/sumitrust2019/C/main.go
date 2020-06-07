package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(X int) string {
	return ""
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var X int
	scanner.Scan()
	X, _ = strconv.Atoi(scanner.Text())
	fmt.Println(solve(X))
}
