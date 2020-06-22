package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

type Input struct {
	lines [][]string // Input は複数行からなる文字列から値をいい感じに取り出すためのメソッドを提供します.

}

func (i *Input) GetFirstValue(rowIndex int) (string, error) {
	return i.GetValue(rowIndex, 0)
}
func (i *Input) GetLine(index int) ([]string, error) {
	if err := i.validateRowIndex(index); err != nil {
		return nil, err
	}
	return i.lines[index], nil
}
func (i *Input) GetValue(rowIndex, colIndex int) (string, error) {
	line, err := i.GetLine(rowIndex)
	if err != nil {
		return "", err
	}
	if colIndex < 0 || colIndex >= len(line) {
		return "", fmt.Errorf("Invalid col index: %v ", colIndex)
	}
	return line[colIndex], nil
}
func (i *Input) MustGetFirstValue(rowIndex int) string {
	_v0, _err := i.GetFirstValue(rowIndex)
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
func solve(input *Input) string {
	a := []rune(input.MustGetFirstValue(0))[0]
	if unicode.IsUpper(a) {
		return "A"
	} else {
		return "a"
	}
}
