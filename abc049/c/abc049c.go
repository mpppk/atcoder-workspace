package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func toRevStr(s string) (ns string) {
	return string(utl_ReverseRune([]rune(s)))
}

func solve(input *utl_Input) string {
	str := input.MustGetFirstValue(0)
	revStr := toRevStr(str)

	for {
		orgStrLen := len(revStr)
		if orgStrLen == 0 {
			return "YES"
		}
		newStr := strings.TrimPrefix(revStr, toRevStr("dreamer"))
		if len(newStr) != orgStrLen {
			revStr = newStr
			continue
		}
		newStr = strings.TrimPrefix(revStr, toRevStr("dream"))
		if len(newStr) != orgStrLen {
			revStr = newStr
			continue
		}
		newStr = strings.TrimPrefix(revStr, toRevStr("eraser"))
		if len(newStr) != orgStrLen {
			revStr = newStr
			continue
		}
		newStr = strings.TrimPrefix(revStr, toRevStr("erase"))
		if len(newStr) != orgStrLen {
			revStr = newStr
			continue
		}
		return "NO"
	}
}

func main() {
	input, err := utl_NewInputFromReader(bufio.NewReader(io.Reader(os.Stdin)))
	utl_PanicIfErrorExist(err)
	fmt.Println(solve(input))
}
