package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/mpppk/atcoder-workspace/lib"
)

func solve(N int, W int, w []int, v []int) int {
	maxV := N * 1000
	list := lib.NewInt2DList(N+10, maxV+10, math.MaxInt64/2)
	list[0][0] = 0
	for i := 0; i < N; i++ {
		for curV := 0; curV <= maxV; curV++ {
			list.ChMin(i+1, curV, list[i][curV])
			if newW := list[i][curV] + w[i]; newW <= W {
				list.ChMin(i+1, curV+v[i], newW)
			}
		}
	}

	retV := 0
	for curV, w := range list[N] {
		if curV > retV && w <= W {
			retV = curV
		}
	}
	return retV
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
	var W int
	scanner.Scan()
	W, _ = strconv.Atoi(scanner.Text())
	w := make([]int, N)
	v := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		w[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		v[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(solve(N, W, w, v))
}
