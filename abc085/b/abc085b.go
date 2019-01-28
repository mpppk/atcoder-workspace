package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(input *utl_Input) int {
	radii := input.MustGetColIntLine(0)[1:]
	//sort.Sort(sort.Reverse(sort.IntSlice(radii)))
	uniqRadii := utl_UniqInt(radii)
	return len(uniqRadii)
}

func main() {
	input := utl_NewInput(bufio.NewScanner(os.Stdin))
	fmt.Println(solve(input))
}
