package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/mpppk/atcoder-workspace/lib"
)

func solve(X int, N int, p []int) int {
	counts := lib.CountInt(p)
	for i := 0; ; i++ {
		if _, ok := counts[X-i]; !ok {
			return X - i
		}
		if _, ok := counts[X+i]; !ok {
			return X + i
		}
	}

	//min := p[0]
	//minAbs := lib.AbsInt(p[0] - X)
	//for _, pi := range p {
	//	if _, ok := counts[pi]; ok {
	//		continue
	//	}
	//	mAbs := lib.AbsInt(pi - X)
	//	if minAbs > mAbs {
	//		min = pi
	//		minAbs = mAbs
	//		continue
	//	}
	//	if minAbs == mAbs && min == pi {
	//		min = pi
	//	}
	//}
	//return min
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var X int
	scanner.Scan()
	X, _ = strconv.Atoi(scanner.Text())
	var N int
	scanner.Scan()
	N, _ = strconv.Atoi(scanner.Text())
	p := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		p[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(solve(X, N, p))
}
