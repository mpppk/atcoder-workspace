package slice

//go:generate genny -in=$GOFILE -out=../gen-$GOFILE gen "Value=int,int16,int32,int64,float32,float64"

import (
	"github.com/cheekybits/genny/generic"
)

type Value generic.Number

func SumValue(values []Value) Value {
	var sum Value = 0
	for _, value := range values {
		sum += value
	}
	return sum
}

func FilterValue(values []Value, f func(v Value) bool) (newValues []Value) {
	for _, value := range values {
		if f(value) {
			newValues = append(newValues, value)
		}
	}
	return
}

func MapValue(values []Value, f func(v Value) Value) (newValues []Value) {
	for _, value := range values {
		newValues = append(newValues, f(value))
	}
	return
}

func ReverseValue(values []Value) []Value {
	newValues := CopyValue(values)
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		newValues[i], newValues[j] = values[j], values[i]
	}
	return newValues
}

func CopyValue(values []Value) []Value {
	dst := make([]Value, len(values))
	copy(dst, values)
	return dst
}

func UniqValue(values []Value) (newValues []Value) {
	m := map[Value]bool{}
	for _, value := range values {
		m[value] = true
	}

	for key := range m {
		newValues = append(newValues, key)
	}
	return
}
