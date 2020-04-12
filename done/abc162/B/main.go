package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64) int {
	cnt := 0
	for i := 1; i <= int(N); i++ {
		if i%3 == 0 || i%5 == 0 {
			continue
		}
		cnt += i
	}
	return cnt
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
