package main

import (
	"bufio"
	"fmt"
	"os"
)

const YES = "Yes"
const NO = "No"

func solve(S string) string {
	if S == "AAA" || S == "BBB" {
		return NO
	}
	return YES
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var S string
	scanner.Scan()
	S = scanner.Text()
	fmt.Println(solve(S))
}
