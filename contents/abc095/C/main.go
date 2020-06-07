package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/mpppk/atcoder-workspace/lib"
)

func solve(A int, B int, C int, X int, Y int) int {
	// 全部別
	noMix := A*X + B*Y

	// 全部mix
	allMix := lib.TernaryOPInt(X > Y, C*2*X, C*2*Y)

	// 少ない方をmix, 多い方を単品
	mix := lib.TernaryOPInt(X > Y, C*2*Y+A*(X-Y), C*2*X+B*(Y-X))

	return lib.MustMinInt(noMix, allMix, mix)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var A int
	scanner.Scan()
	A, _ = strconv.Atoi(scanner.Text())
	var B int
	scanner.Scan()
	B, _ = strconv.Atoi(scanner.Text())
	var C int
	scanner.Scan()
	C, _ = strconv.Atoi(scanner.Text())
	var X int
	scanner.Scan()
	X, _ = strconv.Atoi(scanner.Text())
	var Y int
	scanner.Scan()
	Y, _ = strconv.Atoi(scanner.Text())
	fmt.Println(solve(A, B, C, X, Y))
}
