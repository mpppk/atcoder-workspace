package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(H int, W int, S []string) string {
	return ""
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var H int
	scanner.Scan()
	H, _ = strconv.Atoi(scanner.Text())
	var W int
	scanner.Scan()
	W, _ = strconv.Atoi(scanner.Text())
	S := make([]string, H)
	for i := 0; i < H; i++ {
		scanner.Scan()
		S[i] = scanner.Text()
	}
	fmt.Println(solve(H, W, S))
}
