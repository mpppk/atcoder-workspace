package lib

import (
	"fmt"
)

func AAAToBits(value AAA, minDigits int) (bits []bool) {
	bin := fmt.Sprintf("%b", int(value))
	digits := 0
	for _, b := range bin {
		digits++
		if b == '0' {
			bits = append(bits, false)
		} else if b == '1' {
			bits = append(bits, true)
		} else {
			panic("invalid bit:" + string(b) + ", " + string('0'))
		}
	}

	for minDigits > digits {
		bits = append([]bool{false}, bits...)
		digits++
	}

	return
}
