package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(A_x int, A_y int, B_x int, B_y int, N int, X []int, Y []int) string {
	return ""
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var A_x int
	scanner.Scan()
	A_x, _ = strconv.Atoi(scanner.Text())
	var A_y int
	scanner.Scan()
	A_y, _ = strconv.Atoi(scanner.Text())
	var B_x int
	scanner.Scan()
	B_x, _ = strconv.Atoi(scanner.Text())
	var B_y int
	scanner.Scan()
	B_y, _ = strconv.Atoi(scanner.Text())
	var N int
	scanner.Scan()
	N, _ = strconv.Atoi(scanner.Text())
	X := make([]int, N)
	Y := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		X[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		Y[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(solve(A_x, A_y, B_x, B_y, N, X, Y))
}
