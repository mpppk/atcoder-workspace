package utl

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func toLines(scanner *bufio.Scanner) [][]string {
	var lines [][]string
	for scanner.Scan() {
		text := TrimSpaceAndNewLineCodeAndTab(scanner.Text())
		if len(text) == 0 {
			lines = append(lines, []string{})
			continue
		}
		line := strings.Split(text, " ")
		lines = append(lines, line)
	}
	return lines
}

func toLinesFromReader(reader *bufio.Reader) (lines [][]string, err error) {
	for {
		chunks, err := readLineAsChunks(reader)
		if err == io.EOF {
			return lines, nil
		}

		if err != nil {
			return nil, fmt.Errorf("failed to read line from reader: %v", err)
		}
		lineStr := TrimSpaceAndNewLineCodeAndTab(strings.Join(chunks, ""))
		line := strings.Split(lineStr, " ")
		lines = append(lines, line)
	}
}

func readLineAsChunks(reader *bufio.Reader) (chunks []string, err error) {
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

type Input struct {
	lines [][]string
}

func (i *Input) validateColIndex(index int) error {
	if index < 0 {
		return errors.New(fmt.Sprintf("index is under zero: %d", index))
	}

	return nil
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

func (i *Input) GetLines(startRowIndex, endRowIndex int) ([][]string, error) {
	if err := i.validateRowIndex(startRowIndex); err != nil {
		return nil, fmt.Errorf("invalid start row index: %v", err)
	}
	if err := i.validateRowIndex(endRowIndex - 1); err != nil {
		return nil, fmt.Errorf("invalid end row index: %v", err)
	}
	return i.lines[startRowIndex:endRowIndex], nil
}

func (i *Input) MustGetLines(rowIndex, colIndex int) [][]string {
	lines, err := i.GetLines(rowIndex, colIndex)
	PanicIfErrorExist(err)
	return lines
}

func (i *Input) MustGetIntLines(rowIndex, colIndex int) (lines [][]int) {
	strLines := i.MustGetLines(rowIndex, colIndex)
	for _, strLine := range strLines {
		line, err := toIntLine(strLine)
		PanicIfErrorExist(err)
		lines = append(lines, line)
	}
	return lines
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

func (i *Input) MustGetIntValue(rowIndex, colIndex int) int {
	value := i.MustGetValue(rowIndex, colIndex)
	intV, err := strconv.Atoi(value)
	PanicIfErrorExist(err)
	return intV
}

func (i *Input) MustGetFirstIntValue(rowIndex int) int {
	return i.MustGetIntValue(rowIndex, 0)
}

func (i *Input) MustGetLastIntValue(rowIndex int) int {
	v := i.MustGetLastValue(rowIndex)
	intV, err := strconv.Atoi(v)
	PanicIfErrorExist(err)
	return intV
}

func (i *Input) MustGetValue(rowIndex, colIndex int) string {
	line, err := i.GetValue(rowIndex, colIndex)
	PanicIfErrorExist(err)
	return line
}

func (i *Input) MustGetFirstValue(rowIndex int) string {
	return i.MustGetValue(rowIndex, 0)
}

func (i *Input) MustGetLastValue(rowIndex int) string {
	line := i.MustGetLine(rowIndex)
	return line[len(line)-1]
}

func (i *Input) GetColLine(colIndex int) (newLine []string, err error) {
	if err := i.validateRowIndex(colIndex); err != nil {
		return nil, err
	}

	for i, line := range i.lines {
		if len(line) <= colIndex {
			return nil, errors.New(fmt.Sprintf("col index(%d) is larger than %dth line length(%d)", colIndex, i, len(line)))
		}
		newLine = append(newLine, line[colIndex])
	}

	return newLine, nil
}

func (i *Input) GetColIntLine(colIndex int) (newLine []int, err error) {
	strLine, err := i.GetColLine(colIndex)
	if err != nil {
		return nil, err
	}
	newLine, err = toIntLine(strLine)
	if err != nil {
		return nil, fmt.Errorf("%dth col index: %v", colIndex, err)
	}
	return
}

func (i *Input) MustGetColLine(colIndex int) (newLine []string) {
	newLine, err := i.GetColLine(colIndex)
	PanicIfErrorExist(err)
	return
}

func (i *Input) MustGetColIntLine(colIndex int) (newLine []int) {
	newLine, err := i.GetColIntLine(colIndex)
	PanicIfErrorExist(err)
	return
}

func (i *Input) GetLine(index int) ([]string, error) {
	if err := i.validateRowIndex(index); err != nil {
		return nil, err
	}
	return i.lines[index], nil
}

func (i *Input) MustGetLine(index int) []string {
	line, err := i.GetLine(index)
	PanicIfErrorExist(err)
	return line
}

func (i *Input) GetIntLine(index int) ([]int, error) {
	if err := i.validateRowIndex(index); err != nil {
		return nil, err
	}

	newLine, err := toIntLine(i.lines[index])
	if err != nil {
		return nil, fmt.Errorf("%dth index: %v", index, err)
	}
	return newLine, nil
}

func (i *Input) GetInt8Line(index int) ([]int8, error) {
	if err := i.validateRowIndex(index); err != nil {
		return nil, err
	}

	newLine, err := toInt8Line(i.lines[index])
	if err != nil {
		return nil, fmt.Errorf("%dth index: %v", index, err)
	}
	return newLine, nil
}

func (i *Input) GetInt16Line(index int) ([]int16, error) {
	if err := i.validateRowIndex(index); err != nil {
		return nil, err
	}

	newLine, err := toInt16Line(i.lines[index])
	if err != nil {
		return nil, fmt.Errorf("%dth index: %v", index, err)
	}
	return newLine, nil
}

func (i *Input) GetInt32Line(index int) ([]int32, error) {
	if err := i.validateRowIndex(index); err != nil {
		return nil, err
	}

	newLine, err := toInt32Line(i.lines[index])
	if err != nil {
		return nil, fmt.Errorf("%dth index: %v", index, err)
	}
	return newLine, nil
}

func (i *Input) GetInt64Line(index int) ([]int64, error) {
	if err := i.validateRowIndex(index); err != nil {
		return nil, err
	}

	newLine, err := toInt64Line(i.lines[index])
	if err != nil {
		return nil, fmt.Errorf("%dth index: %v", index, err)
	}
	return newLine, nil
}

func (i *Input) MustGetIntLine(index int) []int {
	line, err := i.GetIntLine(index)
	PanicIfErrorExist(err)
	return line
}
func (i *Input) MustGetInt8Line(index int) []int8 {
	line, err := i.GetInt8Line(index)
	PanicIfErrorExist(err)
	return line
}
func (i *Input) MustGetInt16Line(index int) []int16 {
	line, err := i.GetInt16Line(index)
	PanicIfErrorExist(err)
	return line
}
func (i *Input) MustGetInt32Line(index int) []int32 {
	line, err := i.GetInt32Line(index)
	PanicIfErrorExist(err)
	return line
}
func (i *Input) MustGetInt64Line(index int) []int64 {
	line, err := i.GetInt64Line(index)
	PanicIfErrorExist(err)
	return line
}

func toInt8Line(line []string) (int8Line []int8, err error) {
	newLine, err := toSpecificBitIntLine(line, 8)
	if err != nil {
		return nil, err
	}
	for _, v := range newLine {
		int8Line = append(int8Line, int8(v))
	}
	return
}

func toInt16Line(line []string) (int16Line []int16, err error) {
	newLine, err := toSpecificBitIntLine(line, 16)
	if err != nil {
		return nil, err
	}
	for _, v := range newLine {
		int16Line = append(int16Line, int16(v))
	}
	return
}

func toInt32Line(line []string) (int32Line []int32, err error) {
	newLine, err := toSpecificBitIntLine(line, 32)
	if err != nil {
		return nil, err
	}
	for _, v := range newLine {
		int32Line = append(int32Line, int32(v))
	}
	return
}

func toInt64Line(line []string) ([]int64, error) {
	return toSpecificBitIntLine(line, 64)
}

func toIntLine(line []string) (intLine []int, err error) {
	for j, v := range line {
		intV, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf(fmt.Sprintf("%dth value: %v", j, err.Error()))
		}
		intLine = append(intLine, intV)
	}
	return intLine, nil
}

func toSpecificBitIntLine(line []string, bitSize int) (intLine []int64, err error) {
	for j, v := range line {
		intV, err := strconv.ParseInt(v, 10, bitSize)
		if err != nil {
			return nil, fmt.Errorf(fmt.Sprintf("%dth value: %v", j, err.Error()))
		}
		intLine = append(intLine, intV)
	}
	return intLine, nil
}

func NewInput(scanner *bufio.Scanner) *Input {
	return &Input{
		lines: toLines(scanner),
	}
}

func NewInputFromReader(reader *bufio.Reader) (*Input, error) {
	lines, err := toLinesFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to create new Input from reader: %v", err)
	}
	return &Input{
		lines: lines,
	}, nil
}
