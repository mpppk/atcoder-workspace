package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	L = 'L'
	R = 'R'
)

var SIntListInitialValue int
var SIntListIsInitialized bool
var SIntListRecFunc func(int) int = nil

type SIntList []int

func lib_NewSIntList(length int, initialValue int) SIntList {
	if SIntListIsInitialized {
		panic("SIntList is already used")
	}
	ret := make([]int, length, length)
	for i := 0; i < length; i++ {
		ret[i] = initialValue
	}
	SIntListInitialValue = initialValue
	return ret
}
func lib_PrintIntSlice(values []int, sep string) {
	for i := 0; i < len(values)-1; i++ {
		fmt.Print(values[i], sep)
	}
	fmt.Println(values[len(values)-1])
}
func lib_TernaryOPInt(ok bool, v1, v2 int) int {
	if ok {
		return v1
	}
	return v2
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
func (si SIntList) Rec(f func(index int) int) {
	SIntListRecFunc = f
}
func (si SIntList) Update(index int) int {
	if SIntListRecFunc == nil {
		panic("SIntListRecFunc is nil")
	}
	if si[int(index)] != SIntListInitialValue {
		return si[int(index)]
	}
	ret := SIntListRecFunc(index)
	si[int(index)] = ret
	return ret
}
func solve(S string) {
	s := []rune(S)
	N := len(s)
	lastPos := lib_NewSIntList(N, -1)
	lastPos.Rec(func(i int) int {
		nextPos := lib_TernaryOPInt(s[i] == R, i+1, i-1)
		if s[nextPos] != s[i] {
			return i
		}
		lPos := lastPos.Update(nextPos)
		return lPos + lib_TernaryOPInt(s[lPos] == L, -1, 1)
	})
	count := make([]int, N, N)
	for i := 0; i < len(lastPos); i++ {
		count[lastPos.Update(i)]++
	}
	lib_PrintIntSlice(count, " ")
}
