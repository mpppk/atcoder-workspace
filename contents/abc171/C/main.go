package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mpppk/atcoder-workspace/lib"
)

func solve(N int) (name string) {
	alphabets := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	getDigits := func(N int) int {
		max := 0
		for i := 1; i < 12; i++ {
			max += lib.PowInt(26, i)
			if N <= max {
				return i
			}
		}
		panic("unexpected value")
	}

	digits := getDigits(N)
	n := N
	for i := 1; i < digits; i++ {
		n -= lib.PowInt(26, i)
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

	//for i := 11; i >= 1; i-- {
	//	if N <= 26 {
	//		if N == 0 {
	//			name += alphabets[25]
	//		} else {
	//			name += alphabets[N-1]
	//		}
	//		return name
	//	}
	//	p := 0
	//	for j := i - 1; j >= 1; j-- {
	//		p += lib.PowInt(26, j)
	//	}
	//	fmt.Println(i, p)
	//	if orgN <= p {
	//		continue
	//	}
	//	n := N / p
	//	N = N % p
	//	if n == 0 {
	//		name += alphabets[25]
	//	} else {
	//		name += alphabets[n-1]
	//	}
	//}
	//for {
	//	n := N % 26
	//	if n == 0 {
	//		name = alphabets[25] + name
	//	} else {
	//		name = alphabets[n-1] + name
	//	}
	//	if N <= 26 {
	//		break
	//	}
	//	N = N / 26
	//}

	return name
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
