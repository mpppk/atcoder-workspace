package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func lib_PowInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
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
func solve(N int) (name string) {
	alphabets := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	getDigits := func(N int) int {
		max := 0
		for i := 1; i < 12; i++ {
			max += lib_PowInt(26, i)
			if N <= max {
				return i
			}
		}
		panic("unexpected value")
	}
	digits := getDigits(N)
	n := N
	for i := 1; i < digits; i++ {
		n -= lib_PowInt(26, i)
	}
	if n == 0 {
		return strings.Repeat("z", digits)
	}
	n--
	for i := 0; i < digits; i++ {
		name = alphabets[n%26] + name
		n /= 26
	}
	return name
	return name
}
