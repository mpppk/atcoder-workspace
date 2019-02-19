package generic

import (
	"fmt"
	"strconv"
)

func (i *Input) GetAAALine(index int) ([]AAA, error) {
	if err := i.validateRowIndex(index); err != nil {
		return nil, err
	}

	newLine, err := StringSliceToAAASlice(i.lines[index])
	if err != nil {
		return nil, fmt.Errorf("%dth index: %v", index, err)
	}
	return newLine, nil
}

func (i *Input) GetAAAValue(rowIndex, colIndex int) (AAA, error) {
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

	return AAA(v), nil
}

func (i *Input) GetFirstAAAValue(rowIndex int) (AAA, error) {
	return i.GetAAAValue(rowIndex, 0)
}
