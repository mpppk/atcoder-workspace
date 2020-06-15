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
