package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

func solve(input *utl_Input) int {
	n := input.MustGetFirstIntValue(0)
	int8Nums := utl_ToDigitSliceInt(n)
	nums := utl_Int8SliceToIntSlice(int8Nums)

	if len(nums) == 1 {
		return nums[0]
	}

	// +を挿入するかどうかのビット列を全組み合わせ作る
	// N桁の+の組み合わせは2^(N-1)通り
	allSum := 0
	for i := 0; i < int(math.Pow(2, float64(len(nums)-1))); i++ {
		bits := utl_IntToBits(i, len(nums)-1)
		chunks := utl_MustChunkIntByBits(nums, bits)
		sum := utl_ReduceIntSlice(chunks, func(acc int, cur []int) int {
			int8Cur := utl_IntSliceToInt8Slice(cur)
			joinedC := utl_DigitsToInt(int8Cur)
			return acc + joinedC
		}, 0)
		allSum += sum
	}
	return allSum
}

func main() {
	input := utl_MustNewInputFromReader(bufio.NewReader(io.Reader(os.Stdin)))
	fmt.Println(solve(input))
}
