package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MOD = 1000000007

func solve(H int64, W int64, a [][]string) string {
	return ""
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var H int64
	scanner.Scan()
	H, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var W int64
	scanner.Scan()
	W, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	a := make([][]string, H)
	for i := int64(0); i < H; i++ {
		a[i] = make([]string, 1)
	}
	for i := int64(0); i < H; i++ {
		for j := int64(0); j < 1; j++ {
			scanner.Scan()
			a[i][j] = scanner.Text()
		}
	}
	fmt.Println(solve(H, W, a))
}
