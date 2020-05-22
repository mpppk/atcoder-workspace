package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(b string) string {
	return ""
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var b string
	scanner.Scan()
	b = scanner.Text()
	fmt.Println(solve(b))
}
