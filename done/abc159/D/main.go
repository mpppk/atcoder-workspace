package main

import (
	"bufio"
	"fmt"
	"github.com/mpppk/atcoder/done/abc159/D"
	"os"
	"strconv"
)

func solve(N int64, A []int64) {
	// 数字ごとにグルーピング
	m := countGroup(A)

	all := int64(0)
	for _, count := range m {
		if count >= 2 {
			all += D.lib_MustCombination(count, 2)
		}
	}

	for _, ai := range A {
		fmt.Println(all - m[ai] + 1)
	}
}

func countGroup(values []int64) map[int64]int64 {
	m := map[int64]int64{}
	for _, value := range values {
		m[value]++
	}
	return m
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
	A := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		A[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, A)
}
