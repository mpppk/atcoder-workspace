package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

type String2DList [][]string

func lib_CopyString(values []string) []string {
	dst := make([]string, len(values))
	copy(dst, values)
	return dst
}
func lib_New2DimStringSliceWithInitialValue(row, col int, initialValue string) [][]string {
	ret := make([][]string, row)
	values := lib_NewStringSliceWithInitialValue(col, initialValue)
	for r := 0; r < row; r++ {
		ret[r] = lib_CopyString(values)
	}
	return ret
}
func lib_NewString2DList(row, col int, initialValue string) String2DList {
	return lib_New2DimStringSliceWithInitialValue(row, col, initialValue)
}
func lib_NewStringSliceWithInitialValue(length int, initialValue string) []string {
	ret := make([]string, length)
	for i := 0; i < length; i++ {
		ret[i] = initialValue
	}
	return ret
}
func (m String2DList) ChBy(row, col int, f func(cur, new string) bool, value string, values ...string) (replaced bool) {
	maxBy := func(f func(cur, new string) bool, values ...string) string {
		max := values[0]
		for _, v := range values {
			if f(max, v) {
				max = v
			}
		}
		return max
	}
	max := maxBy(f, append(values, value)...)
	if f(m[row][col], max) {
		m[row][col] = max
		return true
	}
	return false
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
	slen, tlen := utf8.RuneCountInString(s), utf8.RuneCountInString(t)
	sm := lib_NewString2DList(slen+5, tlen+5, "")
	compare := func(cur, new string) bool {
		return len(cur) < len(new)
	}
	for si, sr := range s {
		for ti, tr := range t {
			sm.ChBy(si+1, ti+1, compare, sm[si][ti], sm[si+1][ti], sm[si][ti+1])
			if sr == tr {
				sm.ChBy(si+1, ti+1, compare, sm[si][ti]+string(sr))
			}
		}
	}
	return sm[slen][tlen]
}
