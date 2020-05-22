package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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
func solve(N int) int {
	allCnt := 0
	for n := 1; n <= N; n += 2 {
		cnt := 0
		for i := 1; i <= N; i += 2 {
			if n%i == 0 {
				cnt++
			}
		}
		if cnt == 8 {
			allCnt++
		}
	}
	return allCnt
}
