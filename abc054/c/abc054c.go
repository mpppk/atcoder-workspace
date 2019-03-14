package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func solve(input *lib_Input) int {
	N := input.MustGetFirstIntValue(0)
	edges := input.MustGetIntLinesFrom(1)
	edges = lib_MapIntSlice(edges, func(edge []int) []int {
		return []int{edge[0] - 1, edge[1] - 1}
	})
	graph := lib_MustNewGraph(N, edges, false)
	paths := lib_IntPermutation(lib_IntRange(2, N+1, 1), N-1)
	paths = lib_MapIntSlice(paths, func(path []int) []int {
		return append([]int{1}, path...)
	})
	paths = lib_MapIntSlice(paths, func(path []int) []int {
		return lib_MapInt(path, func(v int) int {
			return v - 1
		})
	})

	currentPathCnt := 0
	for _, path := range paths {
		if graph.IsValidPath(path) {
			currentPathCnt++
		}
	}
	return currentPathCnt
}

func main() {
	input := lib_MustNewInputFromReader(bufio.NewReader(io.Reader(os.Stdin)))
	fmt.Println(solve(input))
}
