package lib

//go:generate genny -in=$GOFILE -out=../gen-$GOFILE gen "AAA=int,int16,int32,int64,float32,float64"

import (
	"errors"
	"fmt"
	"strconv"
)

// SumAAA は、与えられた値の合計値を返します.
func SumAAA(values []AAA) AAA {
	var sum AAA = 0
	for _, value := range values {
		sum += value
	}
	return sum
}

// FilterAAA は、与えられた値それぞれを関数へ適用し、結果がtrueになる値のSliceを返します.
func FilterAAA(values []AAA, f func(v AAA) bool) (newValues []AAA) {
	for _, value := range values {
		if f(value) {
			newValues = append(newValues, value)
		}
	}
	return
}

// FilterAAASlice は、与えられたSliceそれぞれを関数へ適用し、結果がtrueになるSliceのSliceを返します.
func FilterAAASlice(values [][]AAA, f func(v []AAA) bool) (newValues [][]AAA) {
	for _, value := range values {
		if f(value) {
			newValues = append(newValues, value)
		}
	}
	return
}

// UniqAAA は、与えられた値から重複を取り除いて返します.
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

// SubtractAAABy は、values1に関数を適用した値からvalues2に関数を適用した値を引いた結果を返します.
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

// SubtractAAA は、values1それぞれの要素からvalues2それぞれの要素を引いた結果を返します.
func SubtractAAA(values1 []AAA, values2 []AAA) (newValues []AAA, err error) {
	return SubtractAAABy(values1, values2, func(v AAA) AAA {
		return v
	})
}

// RDiffAAABy は、valuesそれぞれに関数を適用した結果の、隣り合う要素の差(n-(n-1))を返します.
func RDiffAAABy(values []AAA, f func(v AAA) AAA) (newValues []AAA, err error) {
	diffValues := append([]AAA{0}, values...)
	newValues, err = SubtractAAABy(values, diffValues[:len(diffValues)-1], f)
	if err != nil {
		return nil, fmt.Errorf("failed to RDiff: %v", err)
	}
	return newValues[1:], nil
}

// RDiffAAA は、隣り合う要素の差(n-(n-1))を返します.
func RDiffAAA(values []AAA) (newValues []AAA, err error) {
	return RDiffAAABy(values, func(v AAA) AAA {
		return v
	})
}

// StringToAAASlice は、Stringの各文字をAAAへ変換したSliceを返します.
func StringToAAASlice(s string) (ValueLine []AAA, err error) {
	for _, r := range s {
		v, err := strconv.ParseInt(string(r), 10, 64)
		if err != nil {
			return nil, err
		}
		ValueLine = append(ValueLine, AAA(v))
	}
	return
}

// StringSliceToAAASlice は、String sliceをAAA sliceへ変換します.
func StringSliceToAAASlice(line []string) (ValueLine []AAA, err error) {
	newLine, err := toSpecificBitIntLine(line, 64)
	if err != nil {
		return nil, err
	}
	for _, v := range newLine {
		ValueLine = append(ValueLine, AAA(v))
	}
	return
}

func NewAAAGridMap(grid [][]string, defaultValue AAA) (m [][]AAA) {
	for _, line := range grid {
		var newLine []AAA
		for range line {
			newLine = append(newLine, defaultValue)
		}
		m = append(m, newLine)
	}
	return
}

// AAARange は、startからendまで、stepずつ増加する数列を返します.(endは含まない)
func AAARange(start, end, step AAA) ([]AAA, error) {
	if end < start {
		return nil, fmt.Errorf("end(%v) is bigger than start(%v)", end, start)
	}
	s := make([]AAA, 0, int(1+(end-start)/step))
	for start < end {
		s = append(s, start)
		start += step
	}
	return s, nil
}
