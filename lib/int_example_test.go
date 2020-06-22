package lib

import "fmt"

func ExampleAAASliceToMap() {
	m := AAASliceToMap([]AAA{1, 2, 3})
	fmt.Println(m)

	// Output:
	// map[1:{} 2:{} 3:{}]
}

func ExampleAAAToBits() {
	bits := AAAToBits(5, 5)
	fmt.Println(bits)

	// Output:
	// [true false true false false]
}
