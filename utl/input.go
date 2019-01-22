package utl

import (
	"bufio"
	"errors"
	"fmt"
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

type Input struct {
	lines [][]string
}

func (i *Input) validateIndex(index int) error {
	if index >= len(i.lines) {
		return errors.New(fmt.Sprintf("index(%d) is larger than lines(%d)", index, len(i.lines)))
	}

	if index < 0 {
		return errors.New(fmt.Sprintf("index is under zero: %d", index))
	}
	return nil
}

func (i *Input) GetLine(index int) ([]string, error) {
	if err := i.validateIndex(index); err != nil {
		return nil, err
	}
	return i.lines[index], nil
}

func (i *Input) GetIntLine(index int) ([]int, error) {
	if err := i.validateIndex(index); err != nil {
		return nil, err
	}

	newLine, err := toIntLine(i.lines[index])
	if err != nil {
		return nil, fmt.Errorf("%dth index: %v", index, err)
	}
	return newLine, nil
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
