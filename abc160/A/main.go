package main

import (
	"bufio"
	"fmt"
	"os"
)

const YES = "Yes"
const NO = "No"

func solve(S string) string {
	if S[2] == S[3] && S[4] == S[5] {
		return YES
	} else {
		return NO
	}
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
