package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func lib_CountInt(values []int) map[int]int {
	m := map[int]int{}
	for _, value := range values {
		m[value]++
	}
	return m
}
func lib_PowInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
func lib_PrimeFactorsInt(n int) (pfs []int) {
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}
	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}
	if n > 2 {
		pfs = append(pfs, n)
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
	var P int
	scanner.Scan()
	P, _ = strconv.Atoi(scanner.Text())
	fmt.Println(solve(N, P))
}
func solve(N int, P int) int {
	pfs := lib_CountInt(lib_PrimeFactorsInt(P))
	if len(pfs) == 0 {
		return 1
	}
	res := 1
	for pf, cnt := range pfs {
		n := cnt / N
		if n > 0 {
			res *= lib_PowInt(pf, n)
		}
	}
	return res
}
