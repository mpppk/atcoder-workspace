package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func solve(input *lib_Input) string {
	m := input.MustReadAsStringGridFrom(1)
	startRowIndex, startColIndex := lib_FindPosFromStringGrid(m, "s")
	goalRowIndex, goalColIndex := lib_FindPosFromStringGrid(m, "g")

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

func main() {
	input := lib_MustNewInputFromReader(bufio.NewReader(io.Reader(os.Stdin)))
	fmt.Println(solve(input))
}
