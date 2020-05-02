package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(A string) {
	if A == "a" {
		fmt.Println("-1")
		return
	}

	fmt.Println("a")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var A string
	scanner.Scan()
	A = scanner.Text()
	solve(A)
}
