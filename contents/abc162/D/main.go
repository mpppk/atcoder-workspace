package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(N int64, S string) int {
	ss := []rune(S)
	rSum := strings.Count(S, "R")
	gSum := strings.Count(S, "G")
	bSum := strings.Count(S, "B")
	cnt := rSum * gSum * bSum

	for i := 0; i < int(N-2); i++ {
		for j := i + 1; j < int(N-1); j++ {
			k := j + (j - i)
			if k >= int(N) {
				continue
			}
			if ss[i] != ss[j] && ss[i] != ss[k] && ss[j] != ss[k] {
				cnt--
			}
		}
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
	var S string
	scanner.Scan()
	S = scanner.Text()
	fmt.Println(solve(N, S))
}
