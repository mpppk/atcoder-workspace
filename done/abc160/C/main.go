package main

import (
	"bufio"
	"fmt"
	"github.com/mpppk/atcoder/done/abc160/C"
	"os"
	"strconv"
)

func solve(K int64, N int64, A []int64) int64 {
	diff := C.lib_MustRDiffInt64(A)
	diff = append(diff, C.lib_AbsInt64(A[0]+(K-A[len(A)-1])))
	return K - C.lib_MustMaxInt64(diff)
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
	var N int64
	scanner.Scan()
	N, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	A := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		A[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	fmt.Println(solve(K, N, A))
}
