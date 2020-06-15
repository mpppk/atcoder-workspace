package lib_test

import (
	"fmt"

	"github.com/mpppk/atcoder-workspace/lib"
)

func ExampleAdjacencyList() {
	x := []int{1, 2, 3}
	y := []int{4, 4, 5}
	list, _ := lib.AdjacencyList(x, y, 5+1)
	fmt.Println(list)

	// Output:
	// [[] [4] [4] [5] [1 2] [3]]
}
