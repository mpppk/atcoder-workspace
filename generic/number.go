package generic

//go:generate genny -in=$GOFILE -out=../gen-$GOFILE gen "AAA=int,int16,int32,int64,float32,float64"

import (
	"errors"
	"fmt"
)

func SumAAA(values []AAA) AAA {
	var sum AAA = 0
	for _, value := range values {
		sum += value
	}
	return sum
}

func FilterAAA(values []AAA, f func(v AAA) bool) (newValues []AAA) {
	for _, value := range values {
		if f(value) {
			newValues = append(newValues, value)
		}
	}
	return
}

func UniqAAA(values []AAA) (newValues []AAA) {
	m := map[AAA]bool{}
	for _, value := range values {
		m[value] = true
	}

	for key := range m {
		newValues = append(newValues, key)
	}
	return
}

func SubtractAAABy(values1 []AAA, values2 []AAA, f func(v AAA) AAA) (newValues []AAA, err error) {
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

func SubtractAAA(values1 []AAA, values2 []AAA) (newValues []AAA, err error) {
	return SubtractAAABy(values1, values2, func(v AAA) AAA {
		return v
	})
}

func RDiffAAABy(values []AAA, f func(v AAA) AAA) (newValues []AAA, err error) {
	diffValues := append([]AAA{0}, values...)
	newValues, err = SubtractAAABy(values, diffValues[:len(diffValues)-1], f)
	if err != nil {
		return nil, fmt.Errorf("failed to RDiff: %v", err)
	}
	return newValues[1:], nil
}

func RDiffAAA(values []AAA) (newValues []AAA, err error) {
	return RDiffAAABy(values, func(v AAA) AAA {
		return v
	})
}

func StringToAAALine(line []string) (ValueLine []AAA, err error) {
	newLine, err := toSpecificBitIntLine(line, 64)
	if err != nil {
		return nil, err
	}
	for _, v := range newLine {
		ValueLine = append(ValueLine, AAA(v))
	}
	return
}
