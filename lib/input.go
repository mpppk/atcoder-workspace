//go:generate goofy mustify --file input.go
package lib

import (
	"bufio"
	"errors"
	"fmt"
	"io"
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

func (i *Input) GetColLine(colIndex int) (newLine []string, err error) {
	if err := i.validateColIndex(colIndex); err != nil {
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

func (i *Input) GetLine(index int) ([]string, error) {
	if err := i.validateRowIndex(index); err != nil {
		return nil, err
	}
	return i.lines[index], nil
}

func (i *Input) ReadAsStringGridFrom(fromIndex int) ([][]string, error) {
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
