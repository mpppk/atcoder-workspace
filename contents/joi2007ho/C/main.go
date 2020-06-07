package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/mpppk/atcoder-workspace/lib"
)

type Pos struct {
	X int
	Y int
}

func Area(x1, y1, x2, y2 int) int {
	return lib.PowInt(lib.AbsInt(x1-x2), 2) + lib.PowInt(lib.AbsInt(y1-y2), 2)
}

func solve(input *lib.Input) int {
	N := input.MustGetFirstIntValue(0)
	pillars := input.MustGetIntLinesFrom(1)
	list := lib.NewBool2DList(5000+5, 5000+5, false)
	for _, pillar := range pillars {
		x, y := pillar[0], pillar[1]
		list[x][y] = true
	}

	maxArea := 0
	for i := 0; i < N-1; i++ {
		for j := i; j < N; j++ {
			x1, y1, x2, y2 := pillars[i][0], pillars[i][1], pillars[j][0], pillars[j][1]
			diffX, diffY := x1-x2, y1-y2
			if lib.MustMinInt(x1-diffY, y1+diffX, x2-diffY, y2+diffX) >= 0 &&
				lib.MustMaxInt(x1-diffY, y1+diffX, x2-diffY, y2+diffX) <= 5000 &&
				list[x1-diffY][y1+diffX] && list[x2-diffY][y2+diffX] {
				area := Area(x1, y1, x2, y2)
				maxArea = lib.TernaryOPInt(maxArea < area, area, maxArea)
			}
			if lib.MustMinInt(x1+diffY, y1-diffX, x2+diffY, y2-diffX) >= 0 &&
				lib.MustMaxInt(x1+diffY, y1-diffX, x2+diffY, y2-diffX) <= 5000 &&
				list[x1+diffY][y1-diffX] && list[x2+diffY][y2-diffX] {
				area := Area(x1, y1, x2, y2)
				maxArea = lib.TernaryOPInt(maxArea < area, area, maxArea)
			}
		}
	}
	return maxArea
}

func main() {
	input := lib.MustNewInputFromReader(bufio.NewReader(io.Reader(os.Stdin)))
	fmt.Println(solve(input))
}
