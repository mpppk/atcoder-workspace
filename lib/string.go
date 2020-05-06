package lib

// FindPosFromStringGrid は、stringの二次元Sliceから、与えられた文字列を持つ要素のindexを返します.
func FindPosFromStringGrid(m [][]string, s string) (row int, col int) {
	for rowIndex, row := range m {
		for colIndex, p := range row {
			if p == s {
				return rowIndex, colIndex
			}
		}
	}
	panic(s + " not found")
}
