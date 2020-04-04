package main

import (
	"bufio"
	"fmt"
	"github.com/mpppk/atcoder/done/abc161/B"
	"os"
	"strconv"
)

const YES = "Yes"
const NO = "No"

func solve(N int64, M int64, A []int64) string {
	allVoteNum := B.lib_SumInt64(A)

	cnt := int64(0)
	for _, ai := range A {
		v := float64(allVoteNum) / float64(4*M)
		if float64(ai) >= v {
			cnt++
		}
	}
	if cnt >= M {
		return YES
	} else {
		return NO
	}
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
	var M int64
	scanner.Scan()
	M, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	A := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		A[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	fmt.Println(solve(N, M, A))
}
