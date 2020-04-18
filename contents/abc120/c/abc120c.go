package main

import (
	"bufio"
	"container/list"
	"fmt"
	"github.com/mpppk/atcoder/done/abc120/c"
	"io"
	"os"
)

func solve(input *c.lib_Input) int {
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
	input := c.lib_MustNewInputFromReader(bufio.NewReader(io.Reader(os.Stdin)))
	fmt.Println(solve(input))
}
