package lib

import (
	"fmt"
	"strconv"
)

func toSpecificBitIntLine(line []string, bitSize int) (intLine []int64, err error) {
	for j, v := range line {
		intV, err := strconv.ParseInt(v, 10, bitSize)
		if err != nil {
			return nil, fmt.Errorf(fmt.Sprintf("%dth value: %v", j, err.Error()))
		}
		intLine = append(intLine, intV)
	}
	return intLine, nil
}

// BitEnumeration は指定されたビット数までの全組み合わせを返します.
func BitEnumeration(digits uint) (enums [][]bool) {
	if digits == 0 {
		return [][]bool{}
	}

	for i := 0; i < 1<<digits; i++ {
		e := []bool{}
		for d := uint(0); d < digits; d++ {
			e = append(e, i>>d&1 == 1)
		}
		enums = append(enums, e)
	}
	return
}

// Combination はnCr(n個の中からr個を選ぶ組み合わせの総数)を返します.
func Combination(n, r int64) (int64, error) {
	if n < r {
		return 0, fmt.Errorf("r(%d) is larger than n(%d)", r, n)
	}

	if n == r {
		return 1, nil
	}

	rRangeFac, err := RangeFactorial(n, r)
	if err != nil {
		return 0, fmt.Errorf("failed in Combination: %v", err)
	}
	rFac, err := Factorial(r)
	if err != nil {
		return 0, fmt.Errorf("failed in Combination: %v", err)
	}
	return rRangeFac / rFac, nil
}

// RangeFactorial は、n-num+1, n-num+2, ... , nを全てかけた値を返します.
func RangeFactorial(n, num int64) (f int64, err error) {
	f = 1
	for i := int64(0); i < num; i++ {
		f *= n - i
	}
	return
}

// Factorial はnの階乗を返します.
func Factorial(n int64) (f int64, err error) {
	if n > 20 { // FIXME Consider 32bit architecture
		return 0, fmt.Errorf("too large Factorical n: %d", n)
	}

	f = 1
	for i := int64(2); i <= n; i++ {
		f = f * i
	}
	return
}

// Gcd はユークリッドの互除法により2つ以上の数の最大公約数を求めます.
func Gcd(a, b int, values ...int) int {
	g := gcd(a, b)
	for _, v := range values {
		g = gcd(g, v)
	}
	return g
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// AdjacencyList は、x(x1...xn)とy(y1...yn)から、x1とy1、xnとynが相互接続しているグラフの隣接リストを返します
func AdjacencyList(x []int, y []int, length int) ([][]int, error) {
	if len(x) != len(y) {
		return nil, fmt.Errorf("x and y lengths are different. x:%d y:%d", len(x), len(y))
	}

	ret := make([][]int, length, length)
	for i := 0; i < len(x); i++ {
		ret[x[i]-1] = append(ret[x[i]-1], y[i])
		ret[y[i]-1] = append(ret[y[i]-1], x[i])
	}
	return ret, nil
}

// DirectedAdjacencyList は、x(x1...xn)とy(y1...yn)から、x1からy1、xnからynへ接続しているグラフの隣接リストを返します
func DirectedAdjacencyList(x []int, y []int, length int) ([][]int, error) {
	if len(x) != len(y) {
		return nil, fmt.Errorf("x and y lengths are different. x:%d y:%d", len(x), len(y))
	}

	ret := make([][]int, length, length)
	for i := 0; i < len(x); i++ {
		ret[x[i]-1] = append(ret[x[i]-1], y[i])
	}
	return ret, nil
}
