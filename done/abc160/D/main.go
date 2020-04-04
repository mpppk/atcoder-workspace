package main

import (
	"bufio"
	"fmt"
	"github.com/mpppk/atcoder/done/abc160/D"
	"os"
	"strconv"
)

func solve(N int, X int, Y int) {
	counter := D.lib_NewCounterInt()
	counter.Func2 = func(i, j int) int {
		return calcDist(i, j, X, Y)
	}
	for i := 1; i <= N; i++ {
		for j := i + 1; j <= N; j++ {
			counter.CountBy2Int(i, j)
		}
	}
	for i := 1; i < N; i++ {
		fmt.Println(counter.M[i])
	}
}

func calcDist(i, j, X, Y int) int {
	normalDist := D.lib_AbsInt(i - j)
	XDist := D.lib_AbsInt(X - i)
	YDist := D.lib_AbsInt(Y - j)
	shortcutDist := XDist + YDist + 1
	XDist2 := D.lib_AbsInt(X - j)
	YDist2 := D.lib_AbsInt(Y - i)
	shortcutDist2 := XDist2 + YDist2 + 1
	return D.lib_MustMinInt([]int{normalDist, shortcutDist, shortcutDist2})
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
	var X int64
	scanner.Scan()
	X, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var Y int64
	scanner.Scan()
	Y, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(int(N), int(X), int(Y))
}
