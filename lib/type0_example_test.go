package lib

import "fmt"

func ExampleNewBitGenerator() {
	gen := NewBitGenerator(2)
	for b := gen.Next(); b != nil; b = gen.Next() {
		fmt.Println(b)
	}

	// Output:
	// [false false]
	// [true false]
	// [false true]
	// [true true]
}
