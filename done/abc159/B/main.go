package main

import (
	"bufio"
	"fmt"
	"github.com/mpppk/atcoder/done/abc159/B"
	"os"
)

const YES = "Yes"
const NO = "No"

func solve(S string) string {
	if S != B.lib_ReverseStr(S) {
		return NO
	}

	runes := []rune(S)
	s1 := string(runes[:(len(runes)-1)/2])
	s2 := string(runes[(len(runes)+3)/2-1:])

	if s1 != B.lib_ReverseStr(s1) || s2 != B.lib_ReverseStr(s2) {
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
