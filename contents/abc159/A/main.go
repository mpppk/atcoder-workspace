package main

import (
	"bufio"
	"fmt"
	"github.com/mpppk/atcoder/done/abc159/A"
	"os"
	"strconv"
)

func solve(N int64, M int64) int {
	nNum := 0
	if N > 1 {
		nNum = A.lib_MustCombination(N, 2)
	}
	mNum := 0
	if m > 1 {
		mNum = A.lib_MustCombination(M, 2)
	}

	return nNum + mNum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var N int64
	scanner.Scan()
	N, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var M int64
	scanner.Scan()
	M, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Println(solve(N, M))
}
