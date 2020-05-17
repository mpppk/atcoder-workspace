package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func solve(A int, B int, H int, M int) float64 {
	//shortDegree := 360.0 / 12.0 * float64(H)
	shortDegree := 360.0/12.0*float64(H) + 30.0/60.0*float64(M)
	longDegree := 360.0 / 60.0 * float64(M)

	smallDegree := shortDegree
	bigDegree := longDegree
	if smallDegree > bigDegree {
		smallDegree, bigDegree = bigDegree, smallDegree
	}
	diff := bigDegree - smallDegree
	if diff > 180 {
		diff = 360 - diff
	}

	//diff := lib.MustMinFloat64(math.Abs(shortDegree-longDegree), math.Abs(longDegree-shortDegree))
	//diff := lib.MustMinFloat64(math.Abs(shortDegree-longDegree), math.Abs(longDegree-shortDegree))
	if diff > 180 {
		diff -= 180
	}
	ret := math.Pow(float64(A), float64(2)) + math.Pow(float64(B), float64(2)) - 2.0*float64(A)*float64(B)*math.Cos(diff*math.Pi/180)
	return math.Sqrt(ret)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var A int
	scanner.Scan()
	A, _ = strconv.Atoi(scanner.Text())
	var B int
	scanner.Scan()
	B, _ = strconv.Atoi(scanner.Text())
	var H int
	scanner.Scan()
	H, _ = strconv.Atoi(scanner.Text())
	var M int
	scanner.Scan()
	M, _ = strconv.Atoi(scanner.Text())
	fmt.Println(solve(A, B, H, M))
}
