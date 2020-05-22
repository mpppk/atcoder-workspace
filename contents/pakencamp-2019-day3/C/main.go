package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/mpppk/atcoder-workspace/lib"
)

func solve(N int, M int, A [][]int) int {
	max := 0
	for _, pair := range lib.MustIntCombination(lib.MustIntRange(1, M+1, 1), 2) {
		sum := 0
		for i := 0; i < N; i++ {
			t1, t2 := pair[0]-1, pair[1]-1
			sum += lib.MustMaxInt(A[i][t1], A[i][t2])
		}
		max = lib.MustMaxInt(max, sum)
	}
	return max
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
	A := make([][]int, N)
	for i := 0; i < N; i++ {
		A[i] = make([]int, M)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			scanner.Scan()
			A[i][j], _ = strconv.Atoi(scanner.Text())
		}
	}
	fmt.Println(solve(N, M, A))
}
