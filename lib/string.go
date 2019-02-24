package lib

func FindPosFromStringGrid(m [][]string, s string) (int, int) {
	for rowIndex, row := range m {
		for colIndex, p := range row {
			if p == s {
				return rowIndex, colIndex
			}
		}
	}
	panic(s + " not found")
}
