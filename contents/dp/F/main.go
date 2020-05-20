package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"

	"github.com/mpppk/atcoder-workspace/lib"
)

func solve(s string, t string) string {
	slen, tlen := utf8.RuneCountInString(s), utf8.RuneCountInString(s)
	m := lib.NewInt2DList(slen+5, tlen+5, 0)
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

	// 復元
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

// TLE
//func solve(s string, t string) string {
//	slen, tlen := utf8.RuneCountInString(s), utf8.RuneCountInString(t)
//	sm := lib.NewString2DList(slen+5, tlen+5, "")
//
//	compare := func(cur, new string) bool {
//		return len(cur) < len(new)
//	}
//
//	for si, sr := range s {
//		for ti, tr := range t {
//			sm.ChBy(si+1, ti+1, compare, sm[si][ti], sm[si+1][ti], sm[si][ti+1])
//			if sr == tr {
//				sm.ChBy(si+1, ti+1, compare, sm[si][ti]+string(sr))
//			}
//		}
//	}
//	return sm[slen][tlen]
//}

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
