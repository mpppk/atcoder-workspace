package generic

import (
	"errors"
	"fmt"
)

func ReduceZZZ(values []ZZZ, f func(acc, cur ZZZ) ZZZ, initial ZZZ) (newValue ZZZ) {
	newValue = initial
	for _, value := range values {
		newValue = f(newValue, value)
	}
	return
}

func ReduceZZZSlice(values [][]ZZZ, f func(acc ZZZ, cur []ZZZ) ZZZ, initial ZZZ) (newValue ZZZ) {
	newValue = initial
	for _, value := range values {
		newValue = f(newValue, value)
	}
	return
}

func CopyZZZ(values []ZZZ) []ZZZ {
	dst := make([]ZZZ, len(values))
	copy(dst, values)
	return dst
}

func ReverseZZZ(values []ZZZ) []ZZZ {
	newValues := CopyZZZ(values)
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		newValues[i], newValues[j] = values[j], values[i]
	}
	return newValues
}

func MapZZZ(values []ZZZ, f func(v ZZZ) ZZZ) (newValues []ZZZ) {
	for _, value := range values {
		newValues = append(newValues, f(value))
	}
	return
}

func MapTypeZZZ(values [][]ZZZ, f func(v []ZZZ) ZZZ) (newValues []ZZZ) {
	for _, value := range values {
		newValues = append(newValues, f(value))
	}
	return
}

func ZipZZZ(valuesList ...[]ZZZ) (newValuesList [][]ZZZ, err error) {
	if len(valuesList) == 0 {
		return nil, errors.New("empty values list are given to ZipZZZ")
	}
	valuesLen := len(valuesList[0])
	for _, values := range valuesList {
		if (len(values)) != valuesLen {
			return nil, errors.New("different lengths values are given to ZipZZZ")
		}
	}

	for i := 0; i < valuesLen; i++ {
		var newValues []ZZZ
		for _, values := range valuesList {
			newValues = append(newValues, values[i])
		}
		newValuesList = append(newValuesList, newValues)
	}
	return
}

func SomeZZZ(values []ZZZ, f func(v ZZZ) bool) bool {
	for _, value := range values {
		if f(value) {
			return true
		}
	}
	return false
}

func SomeZZZSlice(valuesList [][]ZZZ, f func(v []ZZZ) bool) bool {
	for _, values := range valuesList {
		if f(values) {
			return true
		}
	}
	return false
}

func EveryZZZ(values []ZZZ, f func(v ZZZ) bool) bool {
	for _, value := range values {
		if !f(value) {
			return false
		}
	}
	return true
}

func EveryZZZSlice(valuesList [][]ZZZ, f func(v []ZZZ) bool) bool {
	for _, values := range valuesList {
		if !f(values) {
			return false
		}
	}
	return true
}

func ChunkZZZByBits(values []ZZZ, bits []bool) (newValues [][]ZZZ, err error) {
	if len(values) != len(bits)+1 {
		return nil, errors.New(fmt.Sprintf("there are different length between values(%d) and bits(%d)", len(values), len(bits)))
	}

	var chunk []ZZZ
	for i, bit := range bits {
		chunk = append(chunk, values[i])
		if bit {
			newValues = append(newValues, chunk)
			chunk = []ZZZ{}
		}
	}
	chunk = append(chunk, values[len(values)-1])
	newValues = append(newValues, chunk)
	return
}
