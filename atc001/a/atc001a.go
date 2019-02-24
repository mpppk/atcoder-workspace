package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func solve(input *lib_Input) string {
	lines := input.MustGetStringLinesFrom(1)
	var m [][]string
	for _, line := range lines {
		var mLine []string
		for _, r := range line[0] {
			mLine = append(mLine, string(r))
		}
		m = append(m, mLine)
	}

	startRowIndex, startColIndex := findPos(m, "s")
	goalRowIndex, goalColIndex := findPos(m, "g")

	visitMap := map[[2]int]bool{}
	if isValidRoute(m, visitMap, startRowIndex, startColIndex, goalRowIndex, goalColIndex) {
		return "Yes"
	}
	return "No"
}

func isValidRoute(m [][]string, visitMap map[[2]int]bool, fromRowIndex, fromColIndex, toRowIndex, toColIndex int) bool {
	currentPos := m[fromRowIndex][fromColIndex]
	coor := [2]int{fromRowIndex, fromColIndex}
	if visitMap[coor] {
		return false
	}

	visitMap[coor] = true

	if currentPos == "g" {
		return true
	}

	if currentPos == "#" {
		return false
	}

	// 上移動
	if fromRowIndex > 0 && isValidRoute(m, visitMap, fromRowIndex-1, fromColIndex, toRowIndex, toColIndex) {
		return true
	}
	// 下移動
	if fromRowIndex < len(m)-1 && isValidRoute(m, visitMap, fromRowIndex+1, fromColIndex, toRowIndex, toColIndex) {
		return true
	}
	// 右移動
	if fromColIndex < len(m[0])-1 && isValidRoute(m, visitMap, fromRowIndex, fromColIndex+1, toRowIndex, toColIndex) {
		return true
	}
	// 左移動
	if fromColIndex > 0 && isValidRoute(m, visitMap, fromRowIndex, fromColIndex-1, toRowIndex, toColIndex) {
		return true
	}
	return false
}

func findPos(m [][]string, s string) (int, int) {
	for rowIndex, row := range m {
		for colIndex, p := range row {
			if p == s {
				return rowIndex, colIndex
			}
		}
	}
	panic(s + " not found")
}

func main() {
	input := lib_MustNewInputFromReader(bufio.NewReader(io.Reader(os.Stdin)))
	fmt.Println(solve(input))
}
