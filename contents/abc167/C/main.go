package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"

	"github.com/mpppk/atcoder-workspace/lib"
)

func solve(input *lib.Input) int {
	N, M, X := input.MustGetFromFirstToThirdIntValue(0)
	lines := input.MustGetIntLineRange(1, N)

	skills := make([]int, M)
	price := getMax(0, X, skills, lines)
	if price == math.MaxInt32 {
		return -1
	} else {
		return price
	}
}

// 現在のスキルと
func getMax(curPrice, X int, skills []int, lines [][]int) int {
	if len(lines) == 0 {
		if satisfySkill(skills, X) {
			return curPrice
		}
		return math.MaxInt32
	}

	// 買わない場合
	curPrice1 := getMax(curPrice, X, skills, lines[1:])

	// 買う場合
	C, A := lines[0][0], lines[0][1:]
	newSkills := make([]int, len(skills))
	copy(newSkills, skills)
	for i, a := range A {
		newSkills[i] += a
	}
	curPrice2 := getMax(curPrice+C, X, newSkills, lines[1:])

	if curPrice1 > curPrice2 {
		return curPrice2
	} else {
		return curPrice1
	}
}

func satisfySkill(skills []int, X int) bool {
	for _, skill := range skills {
		if skill < X {
			return false
		}
	}
	return true
}

func main() {
	input := lib.MustNewInputFromReader(bufio.NewReader(io.Reader(os.Stdin)))
	fmt.Println(solve(input))
}
