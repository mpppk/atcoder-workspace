package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MOD = 1000000007

var SIntListInitialValue int
var SIntListIsInitialized bool

type SIntList []int

func lib_ModSum(a, b, mod int) int {
	return ((a % mod) + (b % mod)) % mod
}
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
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var N int
	scanner.Scan()
	N, _ = strconv.Atoi(scanner.Text())
	var M int
	scanner.Scan()
	M, _ = strconv.Atoi(scanner.Text())
	a := make([]int, M)
	for i := 0; i < M; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(solve(N, M, a))
}
func solve(N int, M int, a []int) int {
	dp := lib_NewSIntList(N+1, 0)
	m := map[int]bool{}
	for _, ai := range a {
		m[ai] = true
	}
	if _, ok := m[1]; !ok {
		dp[1] = 1
	}
	if N == 1 {
		return dp[1]
	}
	if _, ok := m[2]; !ok {
		dp[2] = dp[1] + 1
	}
	if N == 2 {
		return dp[2]
	}
	for i := 3; i <= N; i++ {
		if _, ok := m[i]; ok {
			continue
		}
		dp[i] = lib_ModSum(dp[i-1], dp[i-2], MOD)
	}
	return dp[N]
}
