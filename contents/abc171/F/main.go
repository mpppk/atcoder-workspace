package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MOD = 1000000007

func solve(K int, S string) string {
	return ""
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var K int
	scanner.Scan()
	K, _ = strconv.Atoi(scanner.Text())
	var S string
	scanner.Scan()
	S = scanner.Text()
	fmt.Println(solve(K, S))
}
