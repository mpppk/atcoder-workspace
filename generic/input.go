//go:generate goofy mustify --file input.go
package generic

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

func (i *Input) GetFirstValue(rowIndex int) (string, error) {
	return i.GetValue(rowIndex, 0)
}

//func (i *Input) GetFirstIntValue(rowIndex int) (int, error) {
//	return i.GetIntValue(rowIndex, 0)
//}

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

func (i *Input) GetLine(index int) ([]string, error) {
	if err := i.validateRowIndex(index); err != nil {
		return nil, err
	}
	return i.lines[index], nil
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
