package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func solve(input *lib_Input) (results []string) {
	N := input.MustGetFirstIntValue(0)
	queries := input.MustGetIntLinesFrom(1)
	uf := lib_NewUnionFindInt(lib_IntRange(0, N, 1))
	for _, query := range queries {
		p, a, b := query[0], query[1], query[2]
		if p == 0 {
			uf.Unite(a, b)
			continue
		}
		results = append(results, lib_ToYesNo(uf.IsSameGroup(a, b)))
	}
	return results
}

func main() {
	input := lib_MustNewInputFromReader(bufio.NewReader(io.Reader(os.Stdin)))
	for _, answer := range solve(input) {
		fmt.Println(answer)
	}
}
