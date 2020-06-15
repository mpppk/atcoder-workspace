package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func lib_AdjacencyList(x []int, y []int, length int) ([][]int, error) {
	if len(x) != len(y) {
		return nil, fmt.Errorf("x and y lengths are different. x:%d y:%d", len(x), len(y))
	}
	ret := make([][]int, length, length)
	for i := 0; i < len(x); i++ {
		ret[x[i]] = append(ret[x[i]], y[i])
		ret[y[i]] = append(ret[y[i]], x[i])
	}
	return ret, nil
}
func lib_ContainsInt(values []int, v int) bool {
	for _, value := range values {
		if value == v {
			return true
		}
	}
	return false
}
func lib_FlatMapInt(values []int, f func(v int) []int) (newValues []int) {
	for _, value := range values {
		newValues = append(newValues, f(value)...)
	}
	return
}
func lib_MustAdjacencyList(x []int, y []int, length int) [][]int {
	_v0, _err := lib_AdjacencyList(x, y, length)
	if _err != nil {
		panic(_err)
	}
	return _v0
}
func lib_UniqInt(values []int) (newValues []int) {
	m := map[int]bool{}
	for _, value := range values {
		m[value] = true
	}
	for key := range m {
		newValues = append(newValues, key)
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
func solve(N int, M int, A []int, B []int) {
	list := lib_MustAdjacencyList(A, B, N+1)
	for i, friends := range list[1:] {
		friendsOfFriends := lib_FlatMapInt(friends, func(friend int) []int {
			return list[friend]
		})
		friendsOfFriends = lib_UniqInt(friendsOfFriends)
		cnt := len(friendsOfFriends)
		for _, friend := range friendsOfFriends {
			if friend == i+1 || lib_ContainsInt(friends, friend) {
				cnt--
			}
		}
		fmt.Println(cnt)
	}
}
