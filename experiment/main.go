package main

import (
	"fmt"

	"github.com/mpppk/atcoder-workspace/lib"
)

func main() {
	//fmt.Println(702 % 26)
	//fmt.Println((702 / 26) % 26)
	for i := 0; ; i++ {
		if lib.Pow10Int(9) < lib.PowInt(2, i) {
			fmt.Println(i)
			break
		}
	}
	//fmt.Println(1000000000000001 - lib.PowInt(26, 30))
	//fmt.Println(18277 % 26)
	//fmt.Println(18277 % (26 * 26))
	//fmt.Println(18277 % (26 * 26 * 26))

	//N := 18288
	//for {
	//	fmt.Println(18288 % 26)
	//
	//}
}
