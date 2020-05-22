package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/mpppk/atcoder-workspace/lib"
)

func solve(N int, P int) int {
	pfs := lib.CountInt(lib.PrimeFactorsInt(P))
	if len(pfs) == 0 {
		return 1
	}
	res := 1
	for pf, cnt := range pfs {
		n := cnt / N
		if n > 0 {
			res *= lib.PowInt(pf, n)
		}
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var N int
	scanner.Scan()
	N, _ = strconv.Atoi(scanner.Text())
	var P int
	scanner.Scan()
	P, _ = strconv.Atoi(scanner.Text())
	fmt.Println(solve(N, P))
}
