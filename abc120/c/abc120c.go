package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"os"
)

func solve(input *lib_Input) int {
	s := input.MustGetFirstValue(0)
	cubes := []rune(s)
	stack := list.New()

	cnt := 0
	for _, cube := range cubes {
		if stack.Len() == 0 {
			stack.PushBack(cube)
			continue
		}

		beforeCube := stack.Back().Value.(rune)
		if cube != beforeCube {
			cnt += 2
			stack.Remove(stack.Back())
			continue
		}
		stack.PushBack(cube)
	}
	return cnt
}

func main() {
	input := lib_MustNewInputFromReader(bufio.NewReader(io.Reader(os.Stdin)))
	fmt.Println(solve(input))
}
