package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int, S string) int {
	cnt := 0
	for pin := 0; pin < 1000; pin++ {
		pinRunes := []rune(fmt.Sprintf("%03d", pin))
		pinIndex := 0
		for _, r := range S {
			if r == pinRunes[pinIndex] {
				pinIndex++
			}
			if pinIndex == 3 {
				cnt++
				break
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
	var N int
	scanner.Scan()
	N, _ = strconv.Atoi(scanner.Text())
	var S string
	scanner.Scan()
	S = scanner.Text()
	fmt.Println(solve(N, S))
}
