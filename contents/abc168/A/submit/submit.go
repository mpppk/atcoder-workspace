package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func lib_CopyInt8(values []int8) []int8 {
	dst := make([]int8, len(values))
	copy(dst, values)
	return dst
}
func lib_ReverseInt8(values []int8) []int8 {
	newValues := lib_CopyInt8(values)
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		newValues[i], newValues[j] = values[j], values[i]
	}
	return newValues
}
func lib_ToDigitSliceInt(n int) (digits []int8) {
	nn := int64(n)
	for {
		if nn <= 0 {
			return lib_ReverseInt8(digits)
		}
		digit := int8(nn % 10)
		digits = append(digits, digit)
		nn /= 10
	}
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var N int
	scanner.Scan()
	N, _ = strconv.Atoi(scanner.Text())
	fmt.Println(solve(N))
}
func solve(N int) string {
	digits := lib_ToDigitSliceInt(N)
	switch digits[len(digits)-1] {
	case 2, 4, 5, 7, 9:
		return "hon"
	case 0, 1, 6, 8:
		return "pon"
	case 3:
		return "bon"
	}
	return ""
}
