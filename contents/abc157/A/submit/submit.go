package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func lib_MaxInt(values []int) (max int, err error) {
	if len(values) == 0 {
		return 0, errors.New("empty slice is given")
	}
	max = values[0]
	for _, value := range values {
		if max < value {
			max = value
		}
	}
	return
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
	fmt.Println(solve(N))
}
func solve(N int64) string {
	fmt.Println(lib_MaxInt([]int{1, 2}))
	return ""
}
