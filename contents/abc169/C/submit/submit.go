package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var A int
	scanner.Scan()
	A, _ = strconv.Atoi(scanner.Text())
	var B float64
	scanner.Scan()
	B, _ = strconv.ParseFloat(scanner.Text(), 64)
	fmt.Println(solve(A, B))
}
func solve(A int, B float64) int {
	return A * int(math.Round(B*100)) / 100
}
