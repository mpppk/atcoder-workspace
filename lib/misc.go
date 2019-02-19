package lib

import "math"

func GetEachDigitSumAAA(n AAA) (sum AAA) {
	for _, digit := range ToDigitSliceAAA(n) {
		sum += AAA(digit)
	}
	return
}

func ToDigitSliceAAA(n AAA) (digits []int8) {
	nn := int64(n)
	for {
		if nn <= 0 {
			return ReverseInt8(digits)
		}
		digit := int8(nn % 10) // FIXME
		digits = append(digits, digit)
		nn /= 10
	}
}

func DigitsToAAA(digits []int8) AAA {
	v := AAA(0)
	for i, digit := range digits {
		v += AAA(float64(digit) * math.Pow(10, float64(len(digits)-i-1)))
	}
	return v
}
