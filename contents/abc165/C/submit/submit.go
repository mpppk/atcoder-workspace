package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type V struct {
	Value float64
	Index int
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
	var M int64
	scanner.Scan()
	M, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var Q int64
	scanner.Scan()
	Q, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	a := make([]int64, Q)
	b := make([]int64, Q)
	c := make([]int64, Q)
	d := make([]int64, Q)
	for i := int64(0); i < Q; i++ {
		scanner.Scan()
		a[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		b[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		c[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		d[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	fmt.Println(solve(N, M, Q, a, b, c, d))
}
func solve(N int64, M int64, Q int64, a []int64, b []int64, c []int64, d []int64) int64 {
	var values []*V
	for i := 0; i < int(Q); i++ {
		v := float64(d[i]) / float64(c[i])
		values = append(values, &V{Value: v, Index: i})
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i].Value > values[j].Value
	})
	cSum := int64(0)
	valueSum := int64(0)
	beforeValueSum := int64(0)
	for i := 0; i < len(values); i++ {
		v := values[i]
		cSum += c[v.Index]
		beforeValueSum = valueSum
		if cSum > N {
			return beforeValueSum
		}
		valueSum += d[v.Index]
	}
	return 0
}
