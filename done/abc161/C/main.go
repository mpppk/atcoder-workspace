package main

import (
	"bufio"
	"fmt"
	"github.com/mpppk/atcoder/done/abc161/C"
	"os"
	"strconv"
)

func solve(N int64, K int64) int64 {
	N = N % K
	for {
		beforeN := N
		N = C.lib_AbsInt64(N - K)
		if beforeN <= N {
			return beforeN
		}
	}
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
	var K int64
	scanner.Scan()
	K, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Println(solve(N, K))
}
