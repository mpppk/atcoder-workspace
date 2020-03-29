package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const ONE = "1"
const TWO = "2"

func solve(input *lib_Input) string {
	S := input.MustGetValue(0, 0)
	queries := input.MustGetStringLinesFrom(2)

	front := make([]string, 0, 2*100000)
	back := make([]string, 0, 2*100000)
	reverse := false
	for _, query := range queries {
		T := query[0]
		if T == ONE {
			reverse = !reverse
			continue
		}

		F, C := query[1], query[2]
		if F == ONE {
			if reverse {
				// back
				//S = S + C
				back = append(back, C)
			} else {
				// front
				//S = C + S
				front = append(front, C)
			}
		} else {
			if reverse {
				// front
				//S = C + S
				front = append(front, C)
			} else {
				// back
				//S = S + C
				back = append(back, C)
			}
		}
	}
	s := lib_ReverseStr(strings.Join(front, "")) + S + strings.Join(back, "")
	if reverse {
		return lib_ReverseStr(s)
	} else {
		return s
	}
}

func main() {
	input := lib_MustNewInputFromReader(bufio.NewReader(io.Reader(os.Stdin)))
	fmt.Println(solve(input))
}
