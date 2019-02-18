package generic

func ZZZSliceToTypeSlice(values []ZZZ) (newValues []Type) {
	for _, value := range values {
		newValues = append(newValues, Type(value))
	}
	return
}
