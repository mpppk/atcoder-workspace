package lib

type BBB3ToAAACache map[BBB]map[BBB]map[BBB]AAA

func (c BBB3ToAAACache) Has(k1, k2, k3 BBB) bool {
	if _, ok := c[k1]; !ok {
		return false
	}

	if _, ok := c[k1][k2]; !ok {
		return false
	}

	_, ok := c[k1][k2][k3]
	return ok
}

func (c BBB3ToAAACache) Get(k1, k2, k3 BBB) (AAA, bool) {
	if _, ok := c[k1]; !ok {
		return 0, false
	}

	if _, ok := c[k1][k2]; !ok {
		return 0, false
	}

	v, ok := c[k1][k2][k3]
	return v, ok
}

func (c BBB3ToAAACache) Set(k1, k2, k3 BBB, v AAA) AAA {
	if _, ok := c[k1]; !ok {
		c[k1] = map[BBB]map[BBB]AAA{}
	}
	if _, ok := c[k1][k2]; !ok {
		c[k1][k2] = map[BBB]AAA{}
	}
	c[k1][k2][k3] = v
	return v
}

// SumAAAByBBB は、与えられた値それぞれに関数を適用した結果を足し合わせて返します.
func SumAAAByBBB(values []AAA, f func(v AAA) BBB) BBB {
	var sum BBB = 0
	for _, value := range values {
		sum += f(value)
	}
	return sum
}

// AAASliceToBBBSlice は、AAA SliceをBBB Sliceに変換します.
func AAASliceToBBBSlice(values []AAA) (newValues []BBB) {
	for _, value := range values {
		newValues = append(newValues, BBB(value))
	}
	return
}

// MapBBBSliceToAAA は、BBBの二次元Sliceの各Sliceに対して関数を適用した結果を返します.
func MapBBBSliceToAAA(values [][]BBB, f func(v []BBB) AAA) (newValues []AAA) {
	for _, value := range values {
		newValues = append(newValues, f(value))
	}
	return
}

// MaxAAAByBBBSlice は、BBBの二次元Sliceの各Sliceに対して関数を適用した結果の最大値を返します.
func MaxAAAByBBBSlice(values [][]BBB, f func(vs []BBB) AAA) (max AAA, err error) {
	return MaxAAA(MapBBBSliceToAAA(values, f)...)
}
