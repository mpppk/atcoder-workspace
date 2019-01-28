package slice

import (
	"errors"

	"github.com/cheekybits/genny/generic"
)

type Type generic.Type

func CopyType(values []Type) []Type {
	dst := make([]Type, len(values))
	copy(dst, values)
	return dst
}

func ReverseType(values []Type) []Type {
	newValues := CopyType(values)
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		newValues[i], newValues[j] = values[j], values[i]
	}
	return newValues
}

func MapType(values []Type, f func(v Type) Type) (newValues []Type) {
	for _, value := range values {
		newValues = append(newValues, f(value))
	}
	return
}

func MapTypeSlice(values [][]Type, f func(v []Type) Type) (newValues []Type) {
	for _, value := range values {
		newValues = append(newValues, f(value))
	}
	return
}

func ZipType(valuesList ...[]Type) (newValuesList [][]Type, err error) {
	if len(valuesList) == 0 {
		return nil, errors.New("empty values list are given to ZipType")
	}
	valuesLen := len(valuesList[0])
	for _, values := range valuesList {
		if (len(values)) != valuesLen {
			return nil, errors.New("different lengths values are given to ZipType")
		}
	}

	for i := 0; i < valuesLen; i++ {
		var newValues []Type
		for _, values := range valuesList {
			newValues = append(newValues, values[i])
		}
		newValuesList = append(newValuesList, newValues)
	}
	return
}

func SomeType(values []Type, f func(v Type) bool) bool {
	for _, value := range values {
		if f(value) {
			return true
		}
	}
	return false
}

func SomeTypeSlice(valuesList [][]Type, f func(v []Type) bool) bool {
	for _, values := range valuesList {
		if f(values) {
			return true
		}
	}
	return false
}

func EveryType(values []Type, f func(v Type) bool) bool {
	for _, value := range values {
		if !f(value) {
			return false
		}
	}
	return true
}

func EveryTypeSlice(valuesList [][]Type, f func(v []Type) bool) bool {
	for _, values := range valuesList {
		if !f(values) {
			return false
		}
	}
	return true
}
