package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/mpppk/atcoder-workspace/lib"
)

func solve(input *lib.Input) int {
	//N := input.MustGetFirstInt64Value(0)
	A := input.MustGetInt64Line(1)

	// j-i=Aj+Ai
	// i+Ai=j-Aj
	m1 := map[int64][]int64{}
	m2 := map[int64][]int64{}
	for i, a := range A {
		i2 := int64(i)
		m1[i2+a] = append(m1[i2+a], i2)
		m2[i2-a] = append(m2[i2-a], i2)
	}

	cnt := 0
	for i, indices := range m1 {
		cnt += len(m2[i]) * len(indices)
	}
	return cnt
}

func main() {
	input := lib.MustNewInputFromReader(bufio.NewReader(io.Reader(os.Stdin)))
	fmt.Println(solve(input))
}
