package slice

import "github.com/cheekybits/genny/generic"

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
