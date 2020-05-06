package lib

import (
	"errors"
	"math"
)

// PowAAA は、xのy乗を返します.
func PowAAA(x, y AAA) AAA {
	return AAA(math.Pow(float64(x), float64(y)))
}

// AbsAAA は、与えられた値の絶対値を返します.
func AbsAAA(value AAA) AAA {
	return AAA(math.Abs(float64(value)))
}

// MaxAAA は、与えられた値の最大値を返します.
func MaxAAA(values ...AAA) (max AAA, err error) {
	if len(values) == 0 {
		return 0, errors.New("empty slice is given")
	}

	max = values[0]
	for _, value := range values {
		if max < value {
			max = value
		}
	}
	return
}

// MinAAAは、与えられた値の最小値を返します.
func MinAAA(values ...AAA) (min AAA, err error) {
	if len(values) == 0 {
		return 0, errors.New("empty slice is given")
	}

	min = values[0]
	for _, value := range values {
		if min > value {
			min = value
		}
	}
	return
}
