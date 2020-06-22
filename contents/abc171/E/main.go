package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/mpppk/atcoder-workspace/lib"
)

func solve(N int, a []int) string {
	var bitsList [][]bool
	for _, ai := range a {
		bitsList = append(bitsList, lib.IntToBits(ai, 30))
	}

	var bits []bool
	for i := 0; i < 30; i++ {
		bitCount := 0
		for n := 0; n < N; n++ {
			if bitsList[n][i] {
				bitCount++
			}
		}
		bits = append(bits, bitCount%2 == 0)
	}

	var answers []int
	for n := 0; n < N; n++ {
		answer := 0
		for i := 0; i < 30; i++ {
			if bits[i] == bitsList[n][i] {
				answer += lib.PowInt(2, i)
			}
		}
		answers = append(answers, answer)
	}

	return lib.JoinInt(answers, " ")
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
	a := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(solve(N, a))
}
