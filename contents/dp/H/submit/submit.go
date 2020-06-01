package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const MOD = 1000000007

type Input struct {
	lines [][]string // Input は複数行からなる文字列から値をいい感じに取り出すためのメソッドを提供します.

}
type Int2DList [][]int
type IntList []int

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
func (i *Input) GetLine(index int) ([]string, error) {
	if err := i.validateRowIndex(index); err != nil {
		return nil, err
	}
	return i.lines[index], nil
}
func (i *Input) GetStringLinesFrom(fromIndex int) (newLines [][]string, err error) {
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
func (i *Input) MustGetStringLinesFrom(fromIndex int) (newLines [][]string) {
	newLines, err := i.GetStringLinesFrom(fromIndex)
	if err != nil {
		panic(err)
	}
	return newLines
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
func lib_MustNewInputFromReader(reader *bufio.Reader) *Input {
	_v0, _err := lib_NewInputFromReader(reader)
	if _err != nil {
		panic(_err)
	}
	return _v0
}
func lib_NewInputFromReader(reader *bufio.Reader) (*Input, error) {
	lines, err := lib_toLinesFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to create new Input from reader: %v", err)
	}
	return &Input{lines: lines}, nil
}
func lib_NewInt2DList(length1, length2 int, initialValue int) Int2DList {
	ret := make([][]int, length1, length1)
	for i := 0; i < length1; i++ {
		ret[i] = lib_NewIntList(length2, initialValue)
	}
	return ret
}
func lib_NewIntList(length int, initialValue int) IntList {
	ret := make([]int, length, length)
	for i := 0; i < length; i++ {
		ret[i] = initialValue
	}
	return ret
}
func lib_TrimSpaceAndNewLineCodeAndTab(s string) string {
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
func solve(input *Input) int {
	H, W := input.MustGetFirstAndSecondIntValue(0)
	lines := input.MustGetStringLinesFrom(1)
	var a [][]rune
	for _, line := range lines {
		a = append(a, []rune(line[0]))
	}
	m := lib_NewInt2DList(H+5, W+5, 0)
	m[0][0] = 1
	for h := 0; h < H; h++ {
		for w := 0; w < W; w++ {
			if h < H-1 && a[h+1][w] != '#' {
				m[h+1][w] += m[h][w]
				if m[h+1][w] >= MOD {
					m[h+1][w] %= MOD
				}
			}
			if w < W-1 && a[h][w+1] != '#' {
				m[h][w+1] += m[h][w]
				if m[h][w+1] >= MOD {
					m[h][w+1] %= MOD
				}
			}
		}
	}
	return m[H-1][W-1]
	return 0
}
