package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var cache = lib_Int3ToIntCache{}
var companySalaryMap = lib_Int3ToIntCache{}

func solve(input *lib_Input) []int {
	N := input.MustGetFirstIntValue(0)
	M := input.MustGetIntValue(0, 1)
	companies := input.MustGetIntLineRange(1, N)
	applicants := input.MustGetIntLineRange(1+N, M)

	for _, company := range companies {
		companySalaryMap.Set(company[0], company[1], company[2], company[3])
	}

	results := lib_MapIntSliceToInt(applicants, func(applicant []int) int {
		return getMaxSalary(applicant[0], applicant[1], applicant[2])
	})
	return results
}

func getMaxSalary(a, b, c int) int {
	if a < 0 || b < 0 || c < 0 {
		return 0
	}
	if cachedResult, ok := cache.Get(a, b, c); ok {
		return cachedResult
	}

	salaries := []int{
		getMaxSalary(a-1, b, c),
		getMaxSalary(a, b-1, c),
		getMaxSalary(a, b, c-1),
	}
	if companySalary, ok := companySalaryMap[a][b][c]; ok {
		salaries = append(salaries, companySalary)
	}

	return cache.Set(a, b, c, lib_MustMaxInt(salaries))
}

func main() {
	input := lib_MustNewInputFromReader(bufio.NewReader(io.Reader(os.Stdin)))
	for _, result := range solve(input) {
		fmt.Println(result)
	}
}
