package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/mpppk/atcoder-workspace/lib"
)

func solve(X int64) {
	for A := int64(0); A < 120; A++ {
		for B := int64(-120); B < 120; B++ {
			if lib.PowInt64(A, 5)-lib.PowInt64(B, 5) == X {
				fmt.Println(A, B)
				return
			}
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var X int64
	scanner.Scan()
	X, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(X)
}
