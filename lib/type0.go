package lib

// ReverseSt　は、引数の文字列を反転させた文字列を返します.
func ReverseStr(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

type bitGenerator struct {
	cur       int
	max       int
	minDigits int
}

// NewBitGenerator は、与えられた桁数のビット列の全組み合わせを返すgeneratorを返します.
// bit全探索を行う際に便利です.
func NewBitGenerator(minDigits int) *bitGenerator {
	return &bitGenerator{
		cur:       0,
		max:       1 << minDigits,
		minDigits: minDigits,
	}
}

// Next は、次のビット組み合わせを返します. 計算量はO(minDigits)です.
func (b *bitGenerator) Next() []bool {
	if b.cur == b.max {
		return nil
	}
	r := IntToBits(b.cur, b.minDigits)
	b.cur++
	return r
}
