package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mpppk/atcoder-workspace/lib"
)

func solve(S string) int {
	maxCnt := 0
	for i, s := range S {
		cnt := 0
		if !strings.ContainsRune("ACGT", s) {
			continue
		}
		for _, s2 := range S[i:] {
			if strings.ContainsRune("ACGT", s2) {
				cnt++
			} else {
				break
			}
		}
		maxCnt = lib.MustMaxInt(maxCnt, cnt)
	}
	return maxCnt
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
