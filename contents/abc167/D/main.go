package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/mpppk/atcoder-workspace/lib"
)

func solve(input *lib.Input) int64 {
	_, K := input.MustGetFirstAndSecondInt64Value(0)
	A := input.MustGetInt64Line(1)

	route := map[int64]int64{}    // key: 町番号, value: 何番目に訪れた町か
	revRoute := map[int64]int64{} // key:何番目に訪れた町か, value:  町番号

	// 最初にループするまでの移動回数を数える
	curPos := int64(1) // 現在の町番号
	cnt := int64(0)
	loopNum := int64(0)
	for {
		route[curPos] = cnt
		revRoute[cnt] = curPos
		cnt++
		last := A[curPos-1]
		if loopStart, ok := route[last]; ok {
			loopNum = cnt - loopStart
			break
		}
		curPos = A[curPos-1]
	}

	if K < cnt {
		return revRoute[K]
	}

	index := (K - cnt) % loopNum
	return revRoute[(cnt-loopNum)+index]
}

func main() {
	input := lib.MustNewInputFromReader(bufio.NewReader(io.Reader(os.Stdin)))
	fmt.Println(solve(input))
}
