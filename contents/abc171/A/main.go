package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"

	"github.com/mpppk/atcoder-workspace/lib"
)

func solve(input *lib.Input) string {
	a := []rune(input.MustGetFirstValue(0))[0]

	if unicode.IsUpper(a) {
		return "A"
	} else {
		return "a"
	}
}

func main() {
	input := lib.MustNewInputFromReader(bufio.NewReader(io.Reader(os.Stdin)))
	fmt.Println(solve(input))
}
