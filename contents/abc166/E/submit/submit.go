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

func (i *Input) GetInt64Line(index int) ([]int64, // GetInt64Line は指定された行をInt64として返します. 存在しない行がある場合、失敗します.
	error) {
	if err := i.validateRowIndex(index); err != nil {
		return nil, err
	}
	newLine, err := lib_StringSliceToInt64Slice(i.lines[index])
	if err != nil {
		return nil, fmt.Errorf("%dth index: %v", index, err)
	}
	return newLine, nil
}

func (i *Input) MustGetInt64Line(index int) []int64 { // MustGetInt64Line は指定された行をInt64として返します. 存在しない行がある場合、失敗します.

	_v0, _err := i.GetInt64Line(index)
	if _err != nil {
		panic(_err)
	}
	return _v0
}

func // MustNewInputFromReader はreaderから入力を読み込み、Inputを生成します.
lib_MustNewInputFromReader(reader *bufio.Reader) *Input {
	_v0, _err := lib_NewInputFromReader(reader)
	if _err != nil {
		panic(_err)
	}
	return _v0
}

func // NewInputFromReader はreaderから入力を読み込み、Inputを生成します.
lib_NewInputFromReader(reader *bufio.Reader) (*Input, error) {
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
func solve(input *Input) int {
	A := input.MustGetInt64Line(1)
	m1 := map[int64][]int64{}
	m2 := map[int64][]int64{}
	for i, a := range A {
		i2 := int64(i)
		m1[i2+a] = append(m1[i2+a], i2)
		m2[i2-a] = append(m2[i2-a], i2)
	}
	cnt := 0
	for i, indices := range m1 {
		cnt += len(m2[i]) * len(indices)
	}
	return cnt
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
