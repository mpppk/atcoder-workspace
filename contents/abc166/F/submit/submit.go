package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const YES = "Yes"
const NO = "No"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var N int64
	scanner.Scan()
	N, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var A int64
	scanner.Scan()
	A, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var B int64
	scanner.Scan()
	B, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var C int64
	scanner.Scan()
	C, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	s := make([]string, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		s[i] = scanner.Text()
	}
	solve(N, A, B, C, s)
}
func solve(N int64, A int64, B int64, C int64, s []string) {
	m := map[rune]int64{'A': A, 'B': B, 'C': C}
	var res []rune
	for i := int64(0); i < N; i++ {
		ss := []rune(s[i])
		s1 := ss[0]
		s2 := ss[1]
		if m[s1] == 0 && m[s2] == 0 {
			fmt.Println(NO)
			return
		}
		if m[s1] == 0 {
			m[s1]++
			m[s2]--
			res = append(res, s1)
			continue
		}
		if m[s2] == 0 {
			m[s1]--
			m[s2]++
			res = append(res, s2)
			continue
		}
		if i == N-1 {
			m[s1]++
			m[s2]--
			res = append(res, s1)
			continue
		}
		if strings.ContainsRune(s[i+1], s1) {
			m[s1]++
			m[s2]--
			res = append(res, s1)
			continue
		}
		if strings.ContainsRune(s[i+1], s2) {
			m[s1]--
			m[s2]++
			res = append(res, s2)
			continue
		}
	}
	fmt.Println(YES)
	for _, r := range res {
		fmt.Println(string(r))
	}
}
