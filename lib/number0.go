package lib

import (
	"fmt"
	"strconv"
)

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

func BitEnumeration(digits uint) (enums [][]bool) {
	if digits == 0 {
		return [][]bool{}
	}

	for i := 0; i < 1<<digits; i++ {
		e := []bool{}
		for d := uint(0); d < digits; d++ {
			e = append(e, i>>d&1 == 1)
		}
		enums = append(enums, e)
	}
	return
}
