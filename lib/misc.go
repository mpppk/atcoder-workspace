package lib

import "math"

// GetEachDigitSumAAA は、格桁の値の挿話を返します.
// ex) 123 -> 1+2+3=6
func GetEachDigitSumAAA(n AAA) (sum AAA) {
	for _, digit := range ToDigitSliceAAA(n) {
		sum += AAA(digit)
	}
	return
}

// ToDigitSliceAAA は、値を桁ごとのsliceとして返します
// ex) 123 -> []int8{1, 2, 3}
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

// DigitsToAAA は、桁ごとの値を持つsliceを値へ変換します.
// []int8{1, 2, 3} -> 123
func DigitsToAAA(digits []int8) AAA {
	v := AAA(0)
	for i, digit := range digits {
		v += AAA(float64(digit) * math.Pow(10, float64(len(digits)-i-1)))
	}
	return v
}
