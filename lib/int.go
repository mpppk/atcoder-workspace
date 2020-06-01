package lib

import (
	"fmt"
)

// AAAToBits は、0と1からなる文字列などを0->false, 1->trueのbool sliceとして返します.
func AAAToBits(value AAA, minDigits int) (bits []bool) {
	bin := fmt.Sprintf("%b", int(value))
	digits := 0
	for _, b := range bin {
		digits++
		if b == '0' {
			bits = append(bits, false)
		} else if b == '1' {
			bits = append(bits, true)
		} else {
			panic("invalid bit:" + string(b) + ", " + string('0'))
		}
	}

	for minDigits > digits {
		bits = append([]bool{false}, bits...)
		digits++
	}

	return
}

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

func (si SAAAList) Rec(f func(AAA) AAA) {
	SAAAListRecFunc = f
}

func (si SAAAList) Update(v AAA) AAA {
	if SAAAListRecFunc == nil {
		panic("SAAAListRecFunc is nil")
	}
	if si[int(v)] != SAAAListInitialValue {
		return si[int(v)]
	}
	ret := SAAAListRecFunc(v)
	si[int(v)] = ret
	return ret
}
