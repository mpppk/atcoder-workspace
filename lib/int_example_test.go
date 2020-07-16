package lib

import "fmt"

func ExampleAAASliceToMap() {
	m := AAASliceToMap([]AAA{1, 2, 3})
	fmt.Println(m)

	// Output:
	// map[1:{} 2:{} 3:{}]
}
