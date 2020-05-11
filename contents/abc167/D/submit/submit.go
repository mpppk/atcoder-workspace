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

type Input struct {
	lines [][]string // Input は複数行からなる文字列から値をいい感じに取り出すためのメソッドを提供します.

}

func (i *Input) GetFirstAndSecondInt64Value(rowIndex int) (int64, int64, error) {
	first, err := i.GetInt64Value(rowIndex, 0)
	if err != nil {
		return first, 0, err
	}
	second, err := i.GetInt64Value(rowIndex, 1)
	return first, second, err
}
func (i *Input) GetInt64Line(index int) ([]int64, error) {
	if err := i.validateRowIndex(index); err != nil {
		return nil, err
	}
	newLine, err := lib_StringSliceToInt64Slice(i.lines[index])
	if err != nil {
		return nil, fmt.Errorf("%dth index: %v", index, err)
	}
	return newLine, nil
}
func (i *Input) GetInt64Value(rowIndex, colIndex int) (int64, error) {
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
	return int64(v), nil
}
func (i *Input) GetLine(index int) ([]string, error) {
	if err := i.validateRowIndex(index); err != nil {
		return nil, err
	}
	return i.lines[index], nil
}
func (i *Input) MustGetFirstAndSecondInt64Value(rowIndex int) (int64, int64) {
	_v0, _v1, _err := i.GetFirstAndSecondInt64Value(rowIndex)
	if _err != nil {
		panic(_err)
	}
	return _v0, _v1
}
func (i *Input) MustGetInt64Line(index int) []int64 {
	_v0, _err := i.GetInt64Line(index)
	if _err != nil {
		panic(_err)
	}
	return _v0
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
func lib_StringSliceToInt64Slice(line []string) (ValueLine []int64, err error) {
	newLine, err := lib_toSpecificBitIntLine(line, 64)
	if err != nil {
		return nil, err
	}
	for _, v := range newLine {
		ValueLine = append(ValueLine, int64(v))
	}
	return
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
func lib_toSpecificBitIntLine(line []string, bitSize int) (intLine []int64, err error) {
	for j, v := range line {
		intV, err := strconv.ParseInt(v, 10, bitSize)
		if err != nil {
			return nil, fmt.Errorf(fmt.Sprintf("%dth value: %v", j, err.Error()))
		}
		intLine = append(intLine, intV)
	}
	return intLine, nil
}
func main() {
	input := lib_MustNewInputFromReader(bufio.NewReader(io.Reader(os.Stdin)))
	fmt.Println(solve(input))
}
func solve(input *Input) int64 {
	_, K := input.MustGetFirstAndSecondInt64Value(0)
	A := input.MustGetInt64Line(1)
	route := map[int64]int64{}
	revRoute := map[int64]int64{}
	curPos := int64(1)
	cnt := int64(0)
	loopNum := int64(0)
	for {
		route[curPos] = cnt
		revRoute[cnt] = curPos
		cnt++
		last := A[curPos-1]
		if loopStart, ok := route[last]; ok {
			loopNum = cnt - loopStart
			break
		}
		curPos = A[curPos-1]
	}
	if K < cnt {
		return revRoute[K]
	}
	index := (K - cnt) % loopNum
	return revRoute[(cnt-loopNum)+index]
}
