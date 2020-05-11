package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	lines [][]string // Input は複数行からなる文字列から値をいい感じに取り出すためのメソッドを提供します.

}

func getMax(curPrice, X int, skills []int, lines [][]int) int {
	if len(lines) == 0 {
		if satisfySkill(skills, X) {
			return curPrice
		}
		return math.MaxInt32
	}
	curPrice1 := getMax(curPrice, X, skills, lines[1:])
	C, A := lines[0][0], lines[0][1:]
	newSkills := make([]int, len(skills))
	copy(newSkills, skills)
	for i, a := range A {
		newSkills[i] += a
	}
	curPrice2 := getMax(curPrice+C, X, newSkills, lines[1:])
	if curPrice1 > curPrice2 {
		return curPrice2
	} else {
		return curPrice1
	}
}
func (i *Input) GetFirstAndSecondIntValue(rowIndex int) (int, int, error) {
	first, err := i.GetIntValue(rowIndex, 0)
	if err != nil {
		return first, 0, err
	}
	second, err := i.GetIntValue(rowIndex, 1)
	return first, second, err
}
func (i *Input) GetFromFirstToThirdIntValue(rowIndex int) (int, int, int, error) {
	first, second, err := i.GetFirstAndSecondIntValue(rowIndex)
	if err != nil {
		return first, second, 0, err
	}
	third, err := i.GetIntValue(rowIndex, 2)
	return first, second, third, err
}
func (i *Input) GetIntLine(index int) ([]int, error) {
	if err := i.validateRowIndex(index); err != nil {
		return nil, err
	}
	newLine, err := lib_StringSliceToIntSlice(i.lines[index])
	if err != nil {
		return nil, fmt.Errorf("%dth index: %v", index, err)
	}
	return newLine, nil
}
func (i *Input) GetIntLineRange(fromRowIndex, rangeNum int) (newLines [][]int, err error) {
	cnt := 0
	for index := range i.lines {
		if index < fromRowIndex {
			continue
		}
		newLine, err := i.GetIntLine(index)
		if err != nil {
			return nil, err
		}
		newLines = append(newLines, newLine)
		cnt++
		if cnt >= rangeNum {
			return newLines, nil
		}
	}
	return
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
func (i *Input) MustGetFromFirstToThirdIntValue(rowIndex int) (int, int, int) {
	_v0, _v1, _v2, _err := i.GetFromFirstToThirdIntValue(rowIndex)
	if _err != nil {
		panic(_err)
	}
	return _v0, _v1, _v2
}
func (i *Input) MustGetIntLineRange(fromRowIndex, rangeNum int) (newLines [][]int) {
	newLines, err := i.GetIntLineRange(fromRowIndex, rangeNum)
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
func lib_StringSliceToIntSlice(line []string) (ValueLine []int, err error) {
	newLine, err := lib_toSpecificBitIntLine(line, 64)
	if err != nil {
		return nil, err
	}
	for _, v := range newLine {
		ValueLine = append(ValueLine, int(v))
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
func satisfySkill(skills []int, X int) bool {
	for _, skill := range skills {
		if skill < X {
			return false
		}
	}
	return true
}
func solve(input *Input) int {
	N, M, X := input.MustGetFromFirstToThirdIntValue(0)
	lines := input.MustGetIntLineRange(1, N)
	skills := make([]int, M)
	price := getMax(0, X, skills, lines)
	if price == math.MaxInt32 {
		return -1
	} else {
		return price
	}
}
