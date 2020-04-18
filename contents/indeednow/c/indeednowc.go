package main

import (
	"bufio"
	"fmt"
	"github.com/mpppk/atcoder/done/indeednow/c"
	"io"
	"os"
)

func solve(input *c.lib_Input) []int {
	var cache = c.lib_Int3ToIntCache{}
	var companySalaryMap = c.lib_Int3ToIntCache{}
	N := input.MustGetFirstIntValue(0)
	M := input.MustGetIntValue(0, 1)
	companies := input.MustGetIntLineRange(1, N)
	applicants := input.MustGetIntLineRange(1+N, M)

	for _, company := range companies {
		a, b, c, w := company[0], company[1], company[2], company[3]
		currentW, ok := companySalaryMap.Get(a, b, c)
		newW := c.lib_TernaryOPInt(ok && currentW > w, currentW, w)
		companySalaryMap.Set(a, b, c, newW)
	}

	results := c.lib_MapIntSliceToInt(applicants, func(applicant []int) int {
		return getMaxSalary(applicant[0], applicant[1], applicant[2], cache, companySalaryMap)
	})
	return results
}

func getMaxSalary(a, b, c int, cache, companySalaryMap c.lib_Int3ToIntCache) int {
	if a < 0 || b < 0 || c < 0 {
		return 0
	}
	if cachedResult, ok := cache.Get(a, b, c); ok {
		return cachedResult
	}

	salaries := []int{
		getMaxSalary(a-1, b, c, cache, companySalaryMap),
		getMaxSalary(a, b-1, c, cache, companySalaryMap),
		getMaxSalary(a, b, c-1, cache, companySalaryMap),
	}
	if companySalary, ok := companySalaryMap[a][b][c]; ok {
		salaries = append(salaries, companySalary)
	}

	return cache.Set(a, b, c, c.lib_MustMaxInt(salaries))
}

func main() {
	input := c.lib_MustNewInputFromReader(bufio.NewReader(io.Reader(os.Stdin)))
	for _, result := range solve(input) {
		fmt.Println(result)
	}
}
