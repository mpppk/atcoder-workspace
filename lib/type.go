package lib

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

func MapZZZSlice(values [][]ZZZ, f func(v []ZZZ) []ZZZ) (newValues [][]ZZZ) {
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

func UnsetZZZ(values []ZZZ, i int) ([]ZZZ, error) {
	if i < 0 {
		return nil, fmt.Errorf("negative number index is given: %d", i)
	}

	if i >= len(values) {
		return nil, fmt.Errorf("index(%d) is larger than slice length(%d)", i, len(values))
	}

	if len(values) == 1 {
		return []ZZZ{}, nil
	}

	if i == len(values)-1 {
		return values[:len(values)-1], nil
	}

	newValues := make([]ZZZ, len(values))
	copy(newValues, values)
	return append(newValues[:i], newValues[i+1:]...), nil
}

func ZZZCombination(values []ZZZ, r int) (combinations [][]ZZZ, err error) {
	if r == 1 {
		for _, value := range values {
			combinations = append(combinations, []ZZZ{value})
		}
		return
	}

	for i := range values {
		newValues := values
		for j := 0; j <= i; j++ {
			newValues = newValues[1:]
		}
		partialCombinations, err := ZZZCombination(newValues, r-1)
		if err != nil {
			return nil, err
		}

		for _, pc := range partialCombinations {
			newC := append(pc, values[i])
			combinations = append(combinations, newC)
		}
	}
	return
}

func ZZZPermutation(values []ZZZ, r int) (permutations [][]ZZZ) {
	if r == 1 {
		for _, value := range values {
			permutations = append(permutations, []ZZZ{value})
		}
		return
	}
	for i := range values {
		newValues := ZZZRemoveFromSlice(values, i)
		for _, pc := range ZZZPermutation(newValues, r-1) {
			newC := append(pc, values[i])
			permutations = append(permutations, newC)
		}
	}
	return
}

func ZZZSliceCombination(values [][]ZZZ, r int) (combinations [][][]ZZZ, err error) {
	if r == 1 {
		for _, value := range values {
			combinations = append(combinations, [][]ZZZ{value})
		}
		return
	}

	for i := range values {
		newValues := values
		for j := 0; j <= i; j++ {
			newValues = newValues[1:]
			if err != nil {
				return nil, err
			}
		}
		partialCombinations, err := ZZZSliceCombination(newValues, r-1)
		if err != nil {
			return nil, err
		}

		for _, pc := range partialCombinations {
			newC := append(pc, values[i])
			combinations = append(combinations, newC)
		}
	}
	return
}

func ZZZRemoveFromSlice(slice []ZZZ, i int) []ZZZ {
	n := make([]ZZZ, len(slice))
	copy(n, slice)
	return append(n[:i], n[i+1:]...)
}
