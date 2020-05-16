package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MOD = 1000000007

func solve(H int, W int, a [][]string) string {
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
	a := make([][]string, H)
	for i := 0; i < H; i++ {
		a[i] = make([]string, 1)
	}
	for i := 0; i < H; i++ {
		for j := 0; j < 1; j++ {
			scanner.Scan()
			a[i][j] = scanner.Text()
		}
	}
	fmt.Println(solve(H, W, a))
}
