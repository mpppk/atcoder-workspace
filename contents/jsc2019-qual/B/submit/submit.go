package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MOD = 1000000007

func lib_Combination(n, r int) (int, error) {
	if n < r {
		return 0, fmt.Errorf("r(%d) is larger than n(%d)", r, n)
	}
	if n == r {
		return 1, nil
	}
	rRangeFac, err := lib_RangeFactorial(n, r)
	if err != nil {
		return 0, fmt.Errorf("failed in Combination: %v", err)
	}
	rFac, err := lib_Factorial(r)
	if err != nil {
		return 0, fmt.Errorf("failed in Combination: %v", err)
	}
	return rRangeFac / rFac, nil
}
func lib_CountInt(values []int) map[int]int {
	m := map[int]int{}
	for _, value := range values {
		m[value]++
	}
	return m
}
func lib_Factorial(n int) (f int, err error) {
	if n > 20 {
		return 0, fmt.Errorf("too large Factorical n: %d", n)
	}
	f = 1
	for i := 2; i <= n; i++ {
		f = f * i
	}
	return
}
func lib_ModMul(a, b, mod int) int {
	return (a % mod) * (b % mod) % mod
}
func lib_ModSum(a, b, mod int) int {
	return ((a % mod) + (b % mod)) % mod
}
func lib_MustCombination(n, r int) int {
	_v0, _err := lib_Combination(n, r)
	if _err != nil {
		panic(_err)
	}
	return _v0
}
func lib_RangeFactorial(n, num int) (f int, err error) {
	f = 1
	for i := 0; i < num; i++ {
		f *= n - i
	}
	return
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
func solve(N int, K int, A []int) int {
	inNum := 0
	for j := 1; j < len(A); j++ {
		for i := 0; i < j; i++ {
			if A[i] > A[j] {
				inNum++
			}
		}
	}
	inSum := lib_ModMul(inNum, K, MOD)
	if K == 1 {
		return inSum
	}
	m := lib_CountInt(A)
	list := make([]int, 2000+5, 2000+5)
	for a, num := range m {
		for i := 1; i < a; i++ {
			list[i] = lib_ModSum(list[i], num, MOD)
		}
	}
	per := 0
	for _, a := range A {
		per = lib_ModSum(per, list[a], MOD)
	}
	return (inSum + lib_ModMul(per, lib_MustCombination(K, 2), MOD)) % MOD
}
