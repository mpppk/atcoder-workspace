package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/mpppk/atcoder-workspace/lib"
)

const MOD = 1000000007

func solve(N int, K int, A []int) int {
	inNum := 0
	for j := 1; j < len(A); j++ {
		for i := 0; i < j; i++ {
			if A[i] > A[j] {
				inNum++
			}
		}
	}
	inSum := lib.ModMul(inNum, K, MOD)
	if K == 1 {
		return inSum
	}

	m := lib.CountInt(A)
	list := make([]int, 2000+5, 2000+5)
	for a, num := range m {
		for i := 1; i < a; i++ {
			list[i] = lib.ModSum(list[i], num, MOD)
		}
	}
	per := 0
	for _, a := range A {
		per = lib.ModSum(per, list[a], MOD)
	}

	return (inSum + lib.ModMul(per, lib.MustCombination(K, 2), MOD)) % MOD
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
	var K int
	scanner.Scan()
	K, _ = strconv.Atoi(scanner.Text())
	A := make([]int, N-1-0+1)
	for i := 0; i < N-1-0+1; i++ {
		scanner.Scan()
		A[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(solve(N, K, A))
}
