package main

import (
	"bufio"
	"container/list"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/mpppk/atcoder-workspace/lib"
)

const YES = "Yes"

func solve(N int, M int, A []int, B []int) {
	m := lib.NewIntList(N+10, math.MaxInt64/2)
	m[1] = 0
	routes := map[int][]int{}
	for i := 0; i < len(A); i++ {
		a, b := A[i], B[i]
		routes[a] = append(routes[a], b)
		routes[b] = append(routes[b], a)
	}

	queue := list.New()
	queue.PushBack(1)
	sirube := lib.NewIntList(N+10, -1)
	sirube[1] = 0
	for queue.Len() > 0 {
		room := queue.Front()
		queue.Remove(room)
		roomId := room.Value.(int)
		neighbors := routes[roomId]
		for _, n := range neighbors {
			if sirube[n] != -1 {
				continue
			}
			if m.ChMin(n, m[roomId]+1) {
				sirube[n] = roomId
			}
			queue.PushBack(n)
		}
	}
	fmt.Println(YES)
	for i := 2; i <= N; i++ {
		fmt.Println(sirube[i])
	}
}

//func solve(N int, M int, A []int, B []int) {
//	m := lib.NewIntMap(N)
//	m[1] = 0
//	routes := map[int][]int{}
//	for i := 0; i < len(A); i++ {
//		a, b := A[i], B[i]
//		routes[a] = append(routes[a], b)
//		routes[b] = append(routes[b], a)
//	}
//
//	queue := list.New()
//	queue.PushBack(1)
//	visited := map[int]bool{}
//	sirube := lib.NewIntMap(N)
//	for queue.Len() > 0 {
//		room := queue.Front()
//		queue.Remove(room)
//		roomId := room.Value.(int)
//		visited[roomId] = true
//		neighbors := routes[roomId]
//		for _, n := range neighbors {
//			if replaced, _ := m.ChMin(n, m[n]+1); replaced {
//				sirube[n] = roomId
//			}
//			if !visited[n] {
//				queue.PushBack(n)
//			}
//		}
//	}
//	fmt.Println(YES)
//	for i := 2; i <= N; i++ {
//		fmt.Println(sirube[i])
//	}
//}

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
