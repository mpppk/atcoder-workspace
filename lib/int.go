package lib

import (
	"math"
)

// PrimeFactors はnを素因数分解します.
func PrimeFactorsAAA(n int) (pfs []int) {
	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}

// Pow10AAA は、10のy乗を返します.
func Pow10AAA(y AAA) AAA {
	return AAA(math.Pow10(int(y)))
}

// AAASliceToMap は、与えられた値をkeyとして持つmapを返します
func AAASliceToMap(values []AAA) map[AAA]struct{} {
	m := map[AAA]struct{}{}
	for _, value := range values {
		m[value] = struct{}{}
	}
	return m
}

type SAAAList []AAA

var SAAAListIsInitialized bool
var SAAAListInitialValue AAA
var SAAAListRecFunc func(AAA) AAA = nil

func NewSAAAList(length int, initialValue AAA) SAAAList {
	if SAAAListIsInitialized {
		panic("SAAAList is already used")
	}
	ret := make([]AAA, length, length)
	for i := 0; i < length; i++ {
		ret[i] = initialValue
	}
	SAAAListInitialValue = initialValue
	return ret
}

func (si SAAAList) ChMin(i AAA, value AAA) bool {
	curV := si[int(i)]
	if curV > value {
		si[int(i)] = value
		return true
	}
	return false
}

func (si SAAAList) ChMax(i AAA, value AAA) bool {
	curV := si[int(i)]
	if curV < value {
		si[int(i)] = value
		return true
	}
	return false
}

// Rec はUpdateで利用する再帰関数を登録します. 登録した関数はUpdateでの呼び出し時に自動でメモ化されます.
func (si SAAAList) Rec(f func(index AAA) AAA) {
	SAAAListRecFunc = f
}

// Update は、indexの要素を事前に登録した再帰関数で更新します.
func (si SAAAList) Update(index AAA) AAA {
	if SAAAListRecFunc == nil {
		panic("SAAAListRecFunc is nil")
	}
	if si[int(index)] != SAAAListInitialValue {
		return si[int(index)]
	}
	ret := SAAAListRecFunc(index)
	si[int(index)] = ret
	return ret
}
