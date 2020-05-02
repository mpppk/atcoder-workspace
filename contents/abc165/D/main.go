package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(A int64, B int64, N int64) int64 {
	//max := int64(0)
	//for x := int64(1); x <= B; x++ {
	//	num := (A*x)/B - A*x/B
	//	if max < num {
	//		max = num
	//	}
	//}
	//mod := N % B
	//for i := 0; i < mod; i++ {
	//
	//}
	if N < (B - 1) {
		return (A*N)/B - A*(N/B)
	}
	return (A*(B-1))/B - A*((B-1)/B)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var A int64
	scanner.Scan()
	A, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var B int64
	scanner.Scan()
	B, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var N int64
	scanner.Scan()
	N, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Println(solve(A, B, N))
}
