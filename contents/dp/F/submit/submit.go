package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

type Int2DList [][]int
type IntList []int

func (a Int2DList) ChMax(i, j int, value int) bool {
	curV := a[i][j]
	if curV < value {
		a[i][j] = value
		return true
	}
	return false
}
func lib_NewInt2DList(length1, length2 int, initialValue int) Int2DList {
	ret := make([][]int, length1, length1)
	for i := 0; i < length1; i++ {
		ret[i] = lib_NewIntList(length2, initialValue)
	}
	return ret
}
func lib_NewIntList(length int, initialValue int) IntList {
	ret := make([]int, length, length)
	for i := 0; i < length; i++ {
		ret[i] = initialValue
	}
	return ret
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var s string
	scanner.Scan()
	s = scanner.Text()
	var t string
	scanner.Scan()
	t = scanner.Text()
	fmt.Println(solve(s, t))
}
func solve(s string, t string) string {
	slen := utf8.RuneCountInString(s)
	tlen := utf8.RuneCountInString(t)
	m := lib_NewInt2DList(slen+5, tlen+5, 0)
	m[0][0] = 0
	for si, sr := range s {
		for ti, tr := range t {
			m.ChMax(si+1, ti+1, m[si][ti])
			m.ChMax(si+1, ti+1, m[si+1][ti])
			m.ChMax(si+1, ti+1, m[si][ti+1])
			if sr == tr {
				m.ChMax(si+1, ti+1, m[si][ti]+1)
			}
		}
	}
	result := ""
	for i, j := slen, tlen; i > 0 && j > 0; {
		switch m[i][j] {
		case m[i-1][j]:
			i--
		case m[i][j-1]:
			j--
		default:
			i--
			j--
			result = string(s[i]) + result
		}
	}
	return result
}
