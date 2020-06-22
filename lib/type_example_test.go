package lib

import "fmt"

func ExampleFlatMapZZZ() {
	values := []ZZZ{1, 2, 3}
	ret := FlatMapZZZ(values, func(v ZZZ) []ZZZ {
		return []ZZZ{v, v}
	})
	fmt.Println(ret)

	// Output:
	// [1 1 2 2 3 3]
}

func ExampleJoinZZZ() {
	s := JoinZZZ([]ZZZ{1, 2, 3}, ",")
	fmt.Println(s)

	// Output:
	// 1,2,3
}

func ExamplePrintZZZSlice() {
	PrintZZZSlice([]ZZZ{1, 2, 3}, ",")

	// Output:
	// 1,2,3
}
