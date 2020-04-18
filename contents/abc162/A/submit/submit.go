package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const YES = "Yes"
const NO = "No"

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
func lib_ToDigitSliceInt64(n int64) (digits []int8) {
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
	var N int64
	scanner.Scan()
	N, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Println(solve(N))
}
func solve(N int64) string {
	digits := lib_ToDigitSliceInt64(N)
	if digits[0] == 7 || digits[1] == 7 || digits[2] == 7 {
		return YES
	}
	return NO
}
