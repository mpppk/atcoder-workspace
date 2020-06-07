package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(a int, b int) string {
	return ""
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var a int
	scanner.Scan()
	a, _ = strconv.Atoi(scanner.Text())
	var b int
	scanner.Scan()
	b, _ = strconv.Atoi(scanner.Text())
	fmt.Println(solve(a, b))
}
