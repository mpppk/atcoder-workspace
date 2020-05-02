package lib

import (
	"fmt"
	"strconv"
)

// GetAAALines は、全ての行をAAAとして取得します.
func (i *Input) GetAAALines() (newLines [][]AAA, err error) {
	return i.GetAAALinesFrom(0)
}

// GetAAALinesFrom は指定された行を含む、それ以降の行をAAAとして返します.存在しない行を指定した場合、失敗します.
func (i *Input) GetAAALinesFrom(fromIndex int) (newLines [][]AAA, err error) {
	for index := range i.lines {
		if index < fromIndex {
			continue
		}
		newLine, err := i.GetAAALine(index)
		if err != nil {
			return nil, err
		}
		newLines = append(newLines, newLine)
	}
	return
}

// GetAAALineRange は指定された行を含むrangeNum行をAAAとして返します. 存在しない行を指定した場合失敗します.
func (i *Input) GetAAALineRange(fromRowIndex, rangeNum int) (newLines [][]AAA, err error) {
	cnt := 0
	for index := range i.lines {
		if index < fromRowIndex {
			continue
		}
		newLine, err := i.GetAAALine(index)
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

// GetAAALine は指定された行をAAAとして返します. 存在しない行がある場合、失敗します.
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

// GetAAAValue は指定された行と列の値をAAAとして返します.存在しない行か列を指定した場合失敗します.
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

// GetFirstAAAValue は指定した行の最初の列の値をAAAとして返します. 存在しない行を指定した場合失敗します.
func (i *Input) GetFirstAAAValue(rowIndex int) (AAA, error) {
	return i.GetAAAValue(rowIndex, 0)
}

// GetColAAALine は指定された列をAAAとして返します。値が存在しない行がある場合失敗します.
func (i *Input) GetColAAALine(colIndex int) (newLine []AAA, err error) {
	strLine, err := i.GetColLine(colIndex)
	if err != nil {
		return nil, err
	}
	newLine, err = StringSliceToAAASlice(strLine)
	if err != nil {
		return nil, fmt.Errorf("%dth col index: %v", colIndex, err)
	}
	return
}
