package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/mpppk/atcoder-workspace/lib"
)

func solve(N int) string {
	digits := lib.ToDigitSliceInt(N)
	switch digits[len(digits)-1] {
	case 2, 4, 5, 7, 9:
		return "hon"
	case 0, 1, 6, 8:
		return "pon"
	case 3:
		return "bon"

	}
	return ""
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
	fmt.Println(solve(N))
}
