package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func lib_CopyBool(values []bool) []bool {
	dst := make([]bool, len(values))
	copy(dst, values)
	return dst
}
func lib_IntToBits(value int, minDigits int) (bits []bool) {
	bin := fmt.Sprintf("%b", int(value))
	digits := utf8.RuneCountInString(bin)
	for _, b := range bin {
		bits = append(bits, b == '1')
	}
	remainBitNums := lib_TernaryOPInt(minDigits > digits, minDigits-digits, 0)
	return append(lib_ReverseBool(bits), make([]bool, remainBitNums)...)
}
func lib_JoinInt(values []int, sep string) string {
	if len(values) == 0 {
		return ""
	}
	var sb strings.Builder
	for i := 0; i < len(values)-1; i++ {
		sb.WriteString(fmt.Sprint(values[i]))
		sb.WriteString(sep)
	}
	sb.WriteString(fmt.Sprint(values[len(values)-1]))
	return sb.String()
}
func lib_PowInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
func lib_ReverseBool(values []bool) []bool {
	newValues := lib_CopyBool(values)
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		newValues[i], newValues[j] = values[j], values[i]
	}
	return newValues
}
func lib_TernaryOPInt(ok bool, v1, v2 int) int {
	if ok {
		return v1
	}
	return v2
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
	a := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(solve(N, a))
}
func solve(N int, a []int) string {
	var bitsList [][]bool
	for _, ai := range a {
		bitsList = append(bitsList, lib_IntToBits(ai, 30))
	}
	var bits []bool
	for i := 0; i < 30; i++ {
		bitCount := 0
		for n := 0; n < N; n++ {
			if bitsList[n][i] {
				bitCount++
			}
		}
		bits = append(bits, bitCount%2 == 0)
	}
	var answers []int
	for n := 0; n < N; n++ {
		answer := 0
		for i := 0; i < 30; i++ {
			if bits[i] == bitsList[n][i] {
				answer += lib_PowInt(2, i)
			}
		}
		answers = append(answers, answer)
	}
	return lib_JoinInt(answers, " ")
}
