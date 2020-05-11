package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mpppk/atcoder-workspace/lib"
)

const YES = "Yes"
const NO = "No"

func solve(S string, T string) string {
	return lib.TernaryOPString(S == T[:len(T)-1], YES, NO)
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
