package main

import (
	"bufio"
	"fmt"
	"github.com/mpppk/atcoder/done/abc161/D"
	"os"
	"strconv"
)

func solve(K int64) int64 {
	if K <= 12 {
		return K
	}

	n := int64(12)
	for curK := 13; curK <= int(K); curK++ {
		n = increment(n)
		//fmt.Println(curK, n)
	}
	return n
}

func increment(n int64) int64 {
	digits := D.lib_ToDigitSliceInt64(n)
	for i := len(digits) - 1; i >= 0; i-- {
		//fmt.Println("i", i)
		digits[i]++
		if digits[i] > 9 {
			continue
		}
		for j := i + 1; j <= len(digits)-1; j++ {
			//fmt.Println("j", j, digits[j-1]-1)
			digits[j] = D.lib_MustMaxInt8([]int8{digits[j-1] - 1, 0})
		}
		//fmt.Println("digits", lib_DigitsToInt64(digits))
		if i == 0 {
			return D.lib_DigitsToInt64(digits)
		}
		if D.lib_AbsInt8(digits[i]-digits[i-1]) < 2 {
			return D.lib_DigitsToInt64(digits)
		} else {
			continue
		}
	}

	newDigits := make([]int8, len(digits)+1)
	newDigits[0] = 1
	return D.lib_DigitsToInt64(newDigits)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var K int64
	scanner.Scan()
	K, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Println(solve(K))
}
