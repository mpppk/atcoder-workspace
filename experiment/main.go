package main

import (
	"fmt"

	"github.com/mpppk/atcoder-workspace/lib"
)

func main() {
	max := lib.PowInt64(10, 9)
	for i := 0; ; i++ {
		diff := lib.PowInt(i+1, 5) - lib.PowInt(i, 5)
		if int64(diff) > max {
			fmt.Println(i)
			return
		}
	}
}
