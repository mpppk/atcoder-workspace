package lib

func AAASliceToBBBSlice(values []AAA) (newValues []BBB) {
	for _, value := range values {
		newValues = append(newValues, BBB(value))
	}
	return
}

func MapBBBSliceToAAA(values [][]BBB, f func(v []BBB) AAA) (newValues []AAA) {
	for _, value := range values {
		newValues = append(newValues, f(value))
	}
	return
}

func MapBBBSlice2ToAAA(values [][][]BBB, f func(v [][]BBB) AAA) (newValues []AAA) {
	for _, value := range values {
		newValues = append(newValues, f(value))
	}
	return
}

func MaxAAAByBBBSlice(values [][]BBB, f func(vs []BBB) AAA) (max AAA, err error) {
	return MaxAAA(MapBBBSliceToAAA(values, f))
}

func MaxAAAByBBBSlice2(values [][][]BBB, f func(vs [][]BBB) AAA) (max AAA, err error) {
	return MaxAAA(MapBBBSlice2ToAAA(values, f))
}
