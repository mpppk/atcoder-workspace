package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"os"

	"github.com/mpppk/atcoder-workspace/lib"
)

type Pos struct {
	Row int
	Col int
}

func solve(input *lib.Input) string {
	H, W := input.MustGetFirstAndSecondIntValue(0)
	grid := input.MustReadAsStringGridFrom(1)
	startRow, startCol := lib.FindPosFromStringGrid(grid, "s")
	endRow, endCol := lib.FindPosFromStringGrid(grid, "g")
	vis := lib.New2DimBoolSlice(H, W)
	stack := list.New()
	stack.PushBack(&Pos{Row: startRow, Col: startCol})

	for stack.Len() > 0 {
		el := stack.Back()
		pos := el.Value.(*Pos)
		stack.Remove(el)

		if pos.Row == endRow && pos.Col == endCol {
			return "Yes"
		}

		if pos.Row > 0 && !vis[pos.Row-1][pos.Col] && grid[pos.Row-1][pos.Col] != "#" {
			stack.PushBack(&Pos{Row: pos.Row - 1, Col: pos.Col})
			vis[pos.Row-1][pos.Col] = true
		}
		if pos.Row < H-1 && !vis[pos.Row+1][pos.Col] && grid[pos.Row+1][pos.Col] != "#" {
			stack.PushBack(&Pos{Row: pos.Row + 1, Col: pos.Col})
			vis[pos.Row+1][pos.Col] = true
		}
		if pos.Col > 0 && !vis[pos.Row][pos.Col-1] && grid[pos.Row][pos.Col-1] != "#" {
			stack.PushBack(&Pos{Row: pos.Row, Col: pos.Col - 1})
			vis[pos.Row][pos.Col-1] = true
		}
		if pos.Col < W-1 && !vis[pos.Row][pos.Col+1] && grid[pos.Row][pos.Col+1] != "#" {
			stack.PushBack(&Pos{Row: pos.Row, Col: pos.Col + 1})
			vis[pos.Row][pos.Col+1] = true
		}
	}
	return "No"
}

func main() {
	input := lib.MustNewInputFromReader(bufio.NewReader(io.Reader(os.Stdin)))
	fmt.Println(solve(input))
}
