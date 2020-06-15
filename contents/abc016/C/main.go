package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/mpppk/atcoder-workspace/lib"
)

func solve(N int, M int, A []int, B []int) {
	list := lib.MustAdjacencyList(A, B, N+1)
	for i, friends := range list[1:] {
		friendsOfFriends := lib.FlatMapInt(friends, func(friend int) []int { return list[friend] })
		friendsOfFriends = lib.UniqInt(friendsOfFriends)
		cnt := len(friendsOfFriends)
		for _, friend := range friendsOfFriends {
			if friend == i+1 || lib.ContainsInt(friends, friend) {
				cnt--
			}
		}
		fmt.Println(cnt)
	}
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
	A := make([]int, M)
	B := make([]int, M)
	for i := 0; i < M; i++ {
		scanner.Scan()
		A[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		B[i], _ = strconv.Atoi(scanner.Text())
	}
	solve(N, M, A, B)
}
