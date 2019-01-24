package slice

//go:generate genny -in=$GOFILE -out=../gen-$GOFILE gen "Value=int,int16,int32,int64,float32,float64"

import "github.com/cheekybits/genny/generic"

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
