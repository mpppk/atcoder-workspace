package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/mpppk/atcoder-workspace/lib"
)

func solve(input *lib.Input) int {
	H, W, K := input.MustGetFromFirstToThirdIntValue(0)
	board := input.MustReadAsStringGridFrom(1)
	cnt := 0

	hGen := lib.NewBitGenerator(H)
	for hb := hGen.Next(); hb != nil; hb = hGen.Next() {
		wGen := lib.NewBitGenerator(W)
		for wb := wGen.Next(); wb != nil; wb = wGen.Next() {
			blackCnt := 0
			for i, row := range board {
				if hb[i] {
					continue
				}
				for j, color := range row {
					if wb[j] {
						continue
					}
					if color == "#" {
						blackCnt++
					}
				}
			}

			if blackCnt == K {
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	input := lib.MustNewInputFromReader(bufio.NewReader(io.Reader(os.Stdin)))
	fmt.Println(solve(input))
}
