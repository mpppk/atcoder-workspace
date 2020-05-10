package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/mpppk/atcoder-workspace/lib"
)

func solve(N int64, W int64, w []int64, v []int64) int64 {
	maxV := N * int64(1000)
	dp := lib.New2DimInt64SliceWithInitialValue(int(N)+1, int(maxV+1), math.MaxInt64/2)
	dp[0][0] = 0

	for i := 0; i < int(N); i++ {
		for j := int64(0); j <= maxV; j++ {
			if newV := j + v[i]; newV <= maxV {
				newW := dp[i][j] + w[i]
				dp[i+1][newV] = lib.MustMinInt64(
					dp[i+1][newV],
					newW,
				)
			}
			dp[i+1][j] = lib.MustMinInt64(
				dp[i+1][j],
				dp[i][j],
			)
		}
	}

	resV := int64(1)
	for i := int64(1); i <= maxV; i++ {
		if dp[N][i] <= W && resV < i {
			resV = i
		}
	}

	return resV
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
	var W int64
	scanner.Scan()
	W, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	w := make([]int64, N)
	v := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		w[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		v[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	fmt.Println(solve(N, W, w, v))
}
