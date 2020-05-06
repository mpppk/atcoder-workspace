package main

import (
	"bufio"
	"container/list"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	Row int
	Col int
}

type Input struct {
	lines [][]string // Input は複数行からなる文字列から値をいい感じに取り出すためのメソッドを提供します.

}

func (i *Input) GetFirstAndSecondIntValue(rowIndex int) (int, int, error) {
	first, err := i.GetIntValue(rowIndex, 0)
	if err != nil {
		return first, 0, err
	}
	second, err := i.GetIntValue(rowIndex, 1)
	return first, second, err
}

func (i *Input) GetIntValue(rowIndex, colIndex int) (int, error) {
	line, err := i.GetLine(rowIndex)
	if err != nil {
		return 0, err
	}
	if colIndex < 0 || colIndex >= len(line) {
		return 0, fmt.Errorf("Invalid col index: %v ", colIndex)
	}
	v, err := strconv.ParseInt(line[colIndex], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to convert string to int: %v, %v", line[colIndex], err)
	}
	return int(v), nil
}

func (i *Input) GetLine(index int) ([]string, // GetFirstIntValue は指定した行の最初の列の値をIntとして返します. 存在しない行を指定した場合失敗します.
	// GetLine は指定された行を返します. 存在しない行がある場合、失敗します.
	error) {
	if err := i.validateRowIndex(index); err != nil {
		return nil, err
	}
	return i.lines[index], nil
}

func (i *Input) GetStringLinesFrom(fromIndex int) (newLines [][]string, // GetStringLinesFrom は指定された行を含む、それ以降の行を返します.存在しない行を指定した場合失敗します.
	err error) {
	for index := range i.lines {
		if index < fromIndex {
			continue
		}
		newLine, err := i.GetLine(index)
		if err != nil {
			return nil, err
		}
		newLines = append(newLines, newLine)
	}
	return
}

func (i *Input) MustGetFirstAndSecondIntValue(rowIndex int) (int, int) {
	_v0, _v1, _err := i.GetFirstAndSecondIntValue(rowIndex)
	if _err != nil {
		panic(_err)
	}
	return _v0, _v1
}

func (i *Input) MustReadAsStringGridFrom(fromIndex int) [][]string { // GetFirstIntValue は指定した行の最初の列の値をIntとして返します. 存在しない行を指定した場合失敗します.
	// MustReadAsStringGridFrom は指定された行以降の行を、一文字ずつのsliceとして返します.

	_v0, _err := i.ReadAsStringGridFrom(fromIndex)
	if _err != nil {
		panic(_err)
	}
	return _v0
}

func (i *Input) ReadAsStringGridFrom(fromIndex int) ([][]string, // ReadAsStringGridFrom は指定された行以降の行を、一文字ずつのsliceとして返します.
	error) {
	lines, err := i.GetStringLinesFrom(fromIndex)
	if err != nil {
		return nil, err
	}
	var m [][]string
	for _, line := range lines {
		if len(line) > 1 {
			return nil, fmt.Errorf("unexpected length line: %v", line)
		}
		var mLine []string
		for _, r := range line[0] {
			mLine = append(mLine, string(r))
		}
		m = append(m, mLine)
	}
	return m, nil
}

func // FindPosFromStringGrid は、stringの二次元Sliceから、与えられた文字列を持つ要素のindexを返します.
lib_FindPosFromStringGrid(m [][]string, s string) (row int, col int) {
	for rowIndex, row := range m {
		for colIndex, p := range row {
			if p == s {
				return rowIndex, colIndex
			}
		}
	}
	panic(s + " not found")
}

func // MustNewInputFromReader はreaderから入力を読み込み、Inputを生成します.
lib_MustNewInputFromReader(reader *bufio.Reader) *Input {
	_v0, _err := lib_NewInputFromReader(reader)
	if _err != nil {
		panic(_err)
	}
	return _v0
}
func lib_New2DimBoolSlice(row, col int) [][]bool {
	ret := make([][]bool, row)
	for r := 0; r < row; r++ {
		ret[r] = make([]bool, col, col)
	}
	return ret
}

func // NewInputFromReader はreaderから入力を読み込み、Inputを生成します.
lib_NewInputFromReader(reader *bufio.Reader) (*Input, error) {
	lines, err := lib_toLinesFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to create new Input from reader: %v", err)
	}
	return &Input{lines: lines}, nil
}

func // TrimSpaceAndNewLineCodeAndTab は、スペース,改行,タブをTrimします.
lib_TrimSpaceAndNewLineCodeAndTab(s string) string {
	return strings.TrimFunc(s, func(r rune) bool {
		return r == ' ' || r == '\r' || r == '\n' || r == '\t'
	})
}
func lib_readLineAsChunks(reader *bufio.Reader) (chunks []string, err error) {
	for {
		chunk, isPrefix, err := reader.ReadLine()
		if err != nil {
			return nil, err
		}
		chunks = append(chunks, string(chunk))
		if !isPrefix {
			return chunks, nil
		}
	}
}
func lib_toLinesFromReader(reader *bufio.Reader) (lines [][]string, err error) {
	for {
		chunks, err := lib_readLineAsChunks(reader)
		if err == io.EOF {
			return lines, nil
		}
		if err != nil {
			return nil, fmt.Errorf("failed to read line from reader: %v", err)
		}
		lineStr := lib_TrimSpaceAndNewLineCodeAndTab(strings.Join(chunks, ""))
		line := strings.Split(lineStr, " ")
		lines = append(lines, line)
	}
}
func main() {
	input := lib_MustNewInputFromReader(bufio.NewReader(io.Reader(os.Stdin)))
	fmt.Println(solve(input))
}
func solve(input *Input) string {
	H, W := input.MustGetFirstAndSecondIntValue(0)
	grid := input.MustReadAsStringGridFrom(1)
	startRow, startCol := lib_FindPosFromStringGrid(grid, "s")
	endRow, endCol := lib_FindPosFromStringGrid(grid, "g")
	vis := lib_New2DimBoolSlice(H, W)
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
func (i *Input) validateRowIndex(index int) error {
	if index >= len(i.lines) {
		return errors.New(fmt.Sprintf("index(%d) is larger than lines(%d)", index, len(i.lines)))
	}
	if index < 0 {
		return errors.New(fmt.Sprintf("index is under zero: %d", index))
	}
	return nil
}
