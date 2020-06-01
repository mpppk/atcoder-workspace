package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/mpppk/atcoder-workspace/lib"
)

const YES = "YES"
const NO = "NO"

func solve(N int, NG []int) string {
	sort.Ints(NG)
	if lib.HasInt(NG, N) {
		return NO
	}
	NG2 := lib.FilterInt(NG, func(v int) bool { return v < N })
	if len(NG2) == 3 && NG[0]+2 == NG[1]+1 && NG[1]+1 == NG[2] {
		return NO
	}

	cnt := 0
	for i := len(NG2) - 1; i >= 0; i-- {
		if (NG2[i]-N%3-cnt)%3 == 0 {
			cnt++
		}
	}
	if (300 - N) < cnt {
		return NO
	}

	return YES
}

func calcMinStep(N int) int {
	a := N % 3
	if a == 0 {
		return N / 3
	} else {
		return N/3 + 1
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
	NG := make([]int, 3)
	for i := 0; i < 3; i++ {
		scanner.Scan()
		NG[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(solve(N, NG))
}
