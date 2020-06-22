package main

import (
	"bufio"
	"os"

	"github.com/mpppk/atcoder-workspace/lib"
)

const (
	L = 'L'
	R = 'R'
)

func solve(S string) {
	s := []rune(S)
	N := len(s)

	lastPos := lib.NewSIntList(N, -1)
	lastPos.Rec(func(i int) int {
		nextPos := lib.TernaryOPInt(s[i] == R, i+1, i-1)
		if s[nextPos] != s[i] {
			return i
		}
		lPos := lastPos.Update(nextPos)
		return lPos + lib.TernaryOPInt(s[lPos] == L, -1, 1)
	})

	count := make([]int, N, N)
	for i := 0; i < len(lastPos); i++ {
		count[lastPos.Update(i)]++
	}
	lib.PrintIntSlice(count, " ")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var S string
	scanner.Scan()
	S = scanner.Text()
	solve(S)
}
