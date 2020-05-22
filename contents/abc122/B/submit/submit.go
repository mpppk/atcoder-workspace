package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func lib_MaxInt(values ...int) (max int, err error) {
	if len(values) == 0 {
		return 0, errors.New("empty slice is given")
	}
	max = values[0]
	for _, value := range values {
		if max < value {
			max = value
		}
	}
	return
}
func lib_MustMaxInt(values ...int) (max int) {
	max, err := lib_MaxInt(values...)
	if err != nil {
		panic(err)
	}
	return max
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
	fmt.Println(solve(S))
}
func solve(S string) int {
	maxCnt := 0
	for i, s := range S {
		cnt := 0
		if !strings.ContainsRune("ACGT", s) {
			continue
		}
		for _, s2 := range S[i:] {
			if strings.ContainsRune("ACGT", s2) {
				cnt++
			} else {
				break
			}
		}
		maxCnt = lib_MustMaxInt(maxCnt, cnt)
	}
	return maxCnt
}
