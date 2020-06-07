package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func lib_HasInt(values []int, v int) bool {
	for _, value := range values {
		if value == v {
			return true
		}
	}
	return false
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
	A := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		A[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(solve(N, A))
}
func solve(N int, A []int) int {
	if lib_HasInt(A, 0) {
		return 0
	}
	max := int(math.Pow10(18))
	ret := 1
	for _, a := range A {
		if a > max/ret {
			return -1
		}
		ret *= a
	}
	return ret
}
