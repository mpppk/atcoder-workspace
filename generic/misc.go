package generic

import "math"

func GetEachDigitSumValue(n Value) (sum Value) {
	for _, digit := range ToDigitSliceValue(n) {
		sum += Value(digit)
	}
	return
}

func ToDigitSliceValue(n Value) (digits []int8) {
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

func DigitsToValue(digits []int8) Value {
	v := Value(0)
	for i, digit := range digits {
		v += Value(float64(digit) * math.Pow(10, float64(len(digits)-i-1)))
	}
	return v
}
