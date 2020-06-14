package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/mpppk/atcoder-workspace/lib"
)

func solve(N int, A []int) (cnt int) {
	aCounts := lib.CountIntAsList(A, lib.Pow10Int(6)+5)
	for _, a := range A {
		if aCounts[a] != 1 {
			continue
		}
		valid := true
		for _, e := range lib.Divisor(a) {
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
