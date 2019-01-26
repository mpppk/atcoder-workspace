package slice

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
