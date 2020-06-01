package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/mpppk/atcoder-workspace/lib"
)

func solve(N int, B []int) int {
	adj := make(map[int][]int, N)
	for i, boss := range B {
		adj[boss] = append(adj[boss], i+2)
	}
	return rec(adj, adj[1])
}

func rec(adj map[int][]int, subs []int) (salary int) {
	var salaries []int
	if len(subs) == 0 {
		return 1
	}
	for _, sub := range subs {
		salaries = append(salaries, rec(adj, adj[sub]))
	}
	return lib.MustMaxInt(salaries...) + lib.MustMinInt(salaries...) + 1
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
	B := make([]int, N-2+1)
	for i := 0; i < N-2+1; i++ {
		scanner.Scan()
		B[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(solve(N, B))
}
