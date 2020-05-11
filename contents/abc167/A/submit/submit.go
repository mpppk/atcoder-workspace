package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	NO  = "No"
	YES = "Yes"
)

func lib_TernaryOPString(ok bool, v1, v2 string) string {
	if ok {
		return v1
	}
	return v2
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
	var T string
	scanner.Scan()
	T = scanner.Text()
	fmt.Println(solve(S, T))
}
func solve(S string, T string) string {
	return lib_TernaryOPString(S == T[:len(T)-1], YES, NO)
}
