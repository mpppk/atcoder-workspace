package generic

//go:generate genny -in=$GOFILE -out=../gen-$GOFILE gen "Value=int,int16,int32,int64,float32,float64"

import (
	"errors"
	"fmt"

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

func SubtractValueBy(values1 []Value, values2 []Value, f func(v Value) Value) (newValues []Value, err error) {
	if len(values1) != len(values2) {
		return nil, errors.New("two values lengths are different")
	}

	for i := 0; i < len(values1); i++ {
		fValue1 := f(values1[i])
		fValue2 := f(values2[i])
		newValues = append(newValues, fValue1-fValue2)
	}
	return newValues, nil
}

func SubtractValue(values1 []Value, values2 []Value) (newValues []Value, err error) {
	return SubtractValueBy(values1, values2, func(v Value) Value {
		return v
	})
}

func RDiffValueBy(values []Value, f func(v Value) Value) (newValues []Value, err error) {
	diffValues := append([]Value{0}, values...)
	newValues, err = SubtractValueBy(values, diffValues[:len(diffValues)-1], f)
	if err != nil {
		return nil, fmt.Errorf("failed to RDiff: %v", err)
	}
	return newValues[1:], nil
}

func RDiffValue(values []Value) (newValues []Value, err error) {
	return RDiffValueBy(values, func(v Value) Value {
		return v
	})
}

func StringToValueLine(line []string) (ValueLine []Value, err error) {
	newLine, err := toSpecificBitIntLine(line, 64)
	if err != nil {
		return nil, err
	}
	for _, v := range newLine {
		ValueLine = append(ValueLine, Value(v))
	}
	return
}
