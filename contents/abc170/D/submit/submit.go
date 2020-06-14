package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func lib_CountIntAsList(values []int, max int) []int {
	m := make([]int, max, max)
	for _, value := range values {
		m[value]++
	}
	return m
}
func lib_Divisor(n int) (res []int) {
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			res = append(res, i)
			if n/i != i {
				res = append(res, n/i)
			}
		}
	}
	return
}
func lib_Pow10Int(y int) int {
	return int(math.Pow10(int(y)))
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
	A := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		A[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(solve(N, A))
}
func solve(N int, A []int) (cnt int) {
	aCounts := lib_CountIntAsList(A, lib_Pow10Int(6)+5)
	for _, a := range A {
		if aCounts[a] != 1 {
			continue
		}
		valid := true
		for _, e := range lib_Divisor(a) {
			if aCounts[e] != 0 && e != a {
				valid = false
				break
			}
		}
		if valid {
			cnt++
		}
	}
	return
}
