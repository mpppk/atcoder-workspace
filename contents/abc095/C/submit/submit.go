package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func lib_MinInt(values ...int) (min int, err error) {
	if len(values) == 0 {
		return 0, errors.New("empty slice is given")
	}
	min = values[0]
	for _, value := range values {
		if min > value {
			min = value
		}
	}
	return
}
func lib_MustMinInt(values ...int) (min int) {
	min, err := lib_MinInt(values...)
	if err != nil {
		panic(err)
	}
	return min
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
	var A int
	scanner.Scan()
	A, _ = strconv.Atoi(scanner.Text())
	var B int
	scanner.Scan()
	B, _ = strconv.Atoi(scanner.Text())
	var C int
	scanner.Scan()
	C, _ = strconv.Atoi(scanner.Text())
	var X int
	scanner.Scan()
	X, _ = strconv.Atoi(scanner.Text())
	var Y int
	scanner.Scan()
	Y, _ = strconv.Atoi(scanner.Text())
	fmt.Println(solve(A, B, C, X, Y))
}
func solve(A int, B int, C int, X int, Y int) int {
	noMix := A*X + B*Y
	allMix := lib_TernaryOPInt(X > Y, C*2*X, C*2*Y)
	mix := lib_TernaryOPInt(X > Y, C*2*Y+A*(X-Y), C*2*X+B*(Y-X))
	return lib_MustMinInt(noMix, allMix, mix)
}
