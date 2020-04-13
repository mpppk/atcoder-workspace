package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/mpppk/atcoder-workspace/lib"
)

const YES = "Yes"
const NO = "No"

func solve(N int64) string {
	digits := lib.ToDigitSliceInt64(N)
	if digits[0] == 7 || digits[1] == 7 || digits[2] == 7 {
		return YES
	}
	return NO
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
	fmt.Println(solve(N))
}
