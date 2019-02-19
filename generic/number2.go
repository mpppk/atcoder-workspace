package generic

func AAASliceToBBBSlice(values []AAA) (newValues []BBB) {
	for _, value := range values {
		newValues = append(newValues, BBB(value))
	}
	return
}
