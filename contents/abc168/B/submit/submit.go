package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode/utf8"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var K int
	scanner.Scan()
	K, _ = strconv.Atoi(scanner.Text())
	var S string
	scanner.Scan()
	S = scanner.Text()
	fmt.Println(solve(K, S))
}
func solve(K int, S string) string {
	slen := utf8.RuneCountInString(S)
	if slen <= K {
		return S
	}
	return S[:K] + "..."
}
