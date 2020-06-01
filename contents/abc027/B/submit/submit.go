package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func lib_SumInt(values []int) int {
	var sum int = 0
	for _, value := range values {
		sum += value
	}
	return sum
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
func solve(N int, a []int) int {
	aSum := lib_SumInt(a)
	if aSum%N != 0 {
		return -1
	}
	perNum := aSum / N
	groupNum, groupSize, bridgeNum := 0, 0, 0
	for i := 0; i < N; i++ {
		if groupSize != 0 {
			bridgeNum++
		}
		groupNum += a[i]
		groupSize++
		if groupNum%groupSize == 0 && groupNum/groupSize == perNum {
			groupNum = 0
			groupSize = 0
		}
	}
	return bridgeNum
}
