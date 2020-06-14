package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const YES = "Yes"
const NO = "No"

func solve(X int, Y int) string {
	for i := 0; i <= X; i++ {
		if 2*i+4*(X-i) == Y {
			return YES
		}
	}
	return NO
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
	var Y int
	scanner.Scan()
	Y, _ = strconv.Atoi(scanner.Text())
	fmt.Println(solve(X, Y))
}
