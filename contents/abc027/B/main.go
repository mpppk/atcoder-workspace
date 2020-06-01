package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/mpppk/atcoder-workspace/lib"
)

func solve(N int, a []int) int {
	aSum := lib.SumInt(a)
	if aSum%N != 0 {
		return -1
	}
	perNum := aSum / N

	groupNum, groupSize, bridgeNum := 0, 0, 0
	for i := 0; i < N; i++ {
		bridgeNum += lib.TernaryOPInt(groupSize != 0, 1, 0)
		groupNum += a[i]
		groupSize++
		if groupNum%groupSize == 0 && groupNum/groupSize == perNum {
			groupNum = 0
			groupSize = 0
		}
	}
	return bridgeNum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var N int
	scanner.Scan()
	N, _ = strconv.Atoi(scanner.Text())
	a := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(solve(N, a))
}
