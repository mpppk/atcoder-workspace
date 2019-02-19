package generic

import (
	"fmt"
)

func (i *Input) GetAAALine(index int) ([]AAA, error) {
	if err := i.validateRowIndex(index); err != nil {
		return nil, err
	}

	newLine, err := StringToAAALine(i.lines[index])
	if err != nil {
		return nil, fmt.Errorf("%dth index: %v", index, err)
	}
	return newLine, nil
}
