package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(l []int) string {
	return ""
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	l := make([]int, 3)
	for i := 0; i < 3; i++ {
		scanner.Scan()
		l[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(solve(l))
}
