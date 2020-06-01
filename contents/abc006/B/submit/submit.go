package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MOD = 10007

var SInt16ListInitialValue int16
var SInt16ListIsInitialized bool

type SInt16List []int16

func lib_NewSInt16List(length int, initialValue int16) SInt16List {
	if SInt16ListIsInitialized {
		panic("SInt16List is already used")
	}
	ret := make([]int16, length, length)
	for i := 0; i < length; i++ {
		ret[i] = initialValue
	}
	SInt16ListInitialValue = initialValue
	return ret
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var n int
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	fmt.Println(solve(n))
}
func solve(n int) int16 {
	dp := lib_NewSInt16List(n+5, 0)
	dp[2], dp[3] = 1, 1
	for i := 4; i < n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % MOD
		dp[i] += dp[i-3] % MOD
		dp[i] %= MOD
	}
	return dp[n-1]
}
