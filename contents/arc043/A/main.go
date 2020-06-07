package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/mpppk/atcoder-workspace/lib"
)

func solve(N int, A int, B int, S []int) string {
	min, max := lib.MustMinInt(S...), lib.MustMaxInt(S...)
	if min == max {
		return "-1"
	}

	P := float64(B) / float64(max-min)
	PS := make([]float64, 0, len(S))
	for _, s := range S {
		PS = append(PS, P*float64(s))
	}
	mean := lib.MeanFloat64(PS)
	diff := mean - float64(A)

	return fmt.Sprintf("%v %v", P, -diff)
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
	var A int
	scanner.Scan()
	A, _ = strconv.Atoi(scanner.Text())
	var B int
	scanner.Scan()
	B, _ = strconv.Atoi(scanner.Text())
	S := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		S[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(solve(N, A, B, S))
}
