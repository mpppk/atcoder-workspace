package lib

import (
	"fmt"
	"math/big"
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

func Combination(n, r int) (int, error) {
	if n < r {
		return 0, fmt.Errorf("r(%d) is larger than n(%d)", r, n)
	}

	if n == r {
		return 1, nil
	}

	rFac, err := Factorial(r)
	if err != nil {
		return 0, fmt.Errorf("too large r: %s", err)
	}
	nFac, err := Factorial(n)
	if err != nil {
		return 0, fmt.Errorf("too large n: %s", err)
	}
	nrFac, err := Factorial(n - r)
	if err != nil {
		return 0, fmt.Errorf("too large n - r: %s", err)
	}
	return nFac / (rFac * nrFac), nil
}

func BigCombination(n, r int) (*big.Int, error) {
	if n < r {
		return nil, fmt.Errorf("r(%d) is larger than n(%d)", r, n)
	}

	if n == r {
		return big.NewInt(1), nil
	}

	rFac := BigFactorial(r)
	nFac := BigFactorial(n)
	nrFac := BigFactorial(n - r)
	return nFac.Div(nFac, rFac.Mul(rFac, nrFac)), nil
}

func ParallelBigCombination(n, r int) (*big.Int, error) {
	if n < r {
		return nil, fmt.Errorf("r(%d) is larger than n(%d)", r, n)
	}

	if n == r {
		return big.NewInt(1), nil
	}

	rChan := make(chan *big.Int)
	nChan := make(chan *big.Int)
	nrChan := make(chan *big.Int)
	go func(r int) {
		rChan <- BigFactorial(r)
	}(r)
	go func(n int) {
		nChan <- BigFactorial(n)
	}(n)
	go func(nr int) {
		nrChan <- BigFactorial(nr)
	}(n - r)

	rFac, nFac, nrFac := <-rChan, <-nChan, <-nrChan
	return nFac.Div(nFac, rFac.Mul(rFac, nrFac)), nil
}

func MemoizedCombination(n, r int) (int, error) {
	if n < r {
		return 0, fmt.Errorf("r(%d) is larger than n(%d)", r, n)
	}

	if n == r {
		return 1, nil
	}

	cache := map[int]int{}
	rFac, err := MemoizedFactorial(r, cache)
	if err != nil {
		return 0, fmt.Errorf("too large r: %s", err)
	}
	nFac, err := MemoizedFactorial(n, cache)
	if err != nil {
		return 0, fmt.Errorf("too large n: %s", err)
	}
	nrFac, err := MemoizedFactorial(n-r, cache)
	if err != nil {
		return 0, fmt.Errorf("too large n - r: %s", err)
	}
	return nFac / (rFac * nrFac), nil
}

func MemoizedBigCombination(n, r int) (*big.Int, error) {
	if n < r {
		return nil, fmt.Errorf("r(%d) is larger than n(%d)", r, n)
	}

	if n == r {
		return big.NewInt(1), nil
	}

	cache := map[int]*big.Int{}
	rFac := MemoizedBigFactorial(r, cache)
	nFac := MemoizedBigFactorial(n, cache)
	nrFac := MemoizedBigFactorial(n-r, cache)
	return nFac.Div(nFac, rFac.Mul(rFac, nrFac)), nil
}

func Factorial(n int) (f int, err error) {
	if n > 20 { // FIXME Consider 32bit architecture
		return 0, fmt.Errorf("too large n: %d", n)
	}

	f = 1
	for i := 2; i <= n; i++ {
		f = f * i
	}
	return
}

func BigFactorial(n int) *big.Int {
	result := big.NewInt(1)
	for i := 2; i <= n; i++ {
		result = result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}

func MemoizedFactorial(n int, cache map[int]int) (int, error) {
	if n > 20 { // FIXME Consider 32bit architecture
		return 0, fmt.Errorf("too large n: %d", n)
	}

	if cachedResult, ok := cache[n]; ok {
		return cachedResult, nil
	}

	if n == 1 {
		return 1, nil
	}

	beforeResult, err := MemoizedFactorial(n-1, cache)
	if err != nil {
		return 0, err
	}
	result := n * beforeResult
	cache[n] = result
	return result, nil
}

func MemoizedBigFactorial(n int, cache map[int]*big.Int) *big.Int {
	if cachedResult, ok := cache[n]; ok {
		return cachedResult
	}

	if n == 1 {
		return big.NewInt(1)
	}

	beforeResult := MemoizedBigFactorial(n-1, cache)
	bigN := big.NewInt(int64(n))
	result := bigN.Mul(bigN, beforeResult)
	cache[n] = result
	return result
}
