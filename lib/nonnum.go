package lib

type ZZZList []ZZZ
type ZZZ2DList [][]ZZZ

func NewZZZList(length int) ZZZList {
	return make(ZZZList, length, length)
}

func NewZZZ2DList(row, col int, initialValue ZZZ) ZZZ2DList {
	return New2DimZZZSliceWithInitialValue(row, col, initialValue)
}

func (m ZZZ2DList) ChBy(row, col int, f func(cur, new ZZZ) bool, value ZZZ, values ...ZZZ) (replaced bool) {
	maxBy := func(f func(cur, new ZZZ) bool, values ...ZZZ) ZZZ {
		max := values[0]
		for _, v := range values {
			if f(max, v) {
				max = v
			}
		}
		return max
	}
	max := maxBy(f, append(values, value)...)
	if f(m[row][col], max) {
		m[row][col] = max
		return true
	}
	return false
}
