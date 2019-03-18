package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

func solve(input *utl_Input) string {
	maxCoordNum := input.MustGetFirstIntValue(0)
	locations := input.MustGetIntLines(1, maxCoordNum+1)
	times := utl_MapIntSlice(locations, func(location []int) int {
		return location[0]
	})

	xys := utl_MapIntSlice(locations, func(location []int) int {
		return location[1] + location[2]
	})

	timeDiffs, err := utl_RDiffInt(times)
	utl_PanicIfErrorExist(err)
	timeDiffs = append([]int{times[0]}, timeDiffs...) // 0秒目から始まることを考慮

	xyDiffs, err := utl_RDiffInt(xys)
	utl_PanicIfErrorExist(err)
	xyDiffs = append([]int{xys[0]}, xyDiffs...)

	xyAbsDiffs := utl_MapInt(xyDiffs, func(v int) int {
		return int(math.Abs(float64(v)))
	})
	timeAndXyDiffs, err := utl_ZipInt(timeDiffs, xyAbsDiffs)
	utl_PanicIfErrorExist(err)

	moves := utl_MapIntSlice(timeAndXyDiffs, func(timeAndXy []int) int {
		return timeAndXy[0] - timeAndXy[1]
	})

	cantMoves := utl_FilterInt(moves, func(move int) bool {
		return move < 0
	})

	if len(cantMoves) > 0 {
		return "No"
	}

	cantMovesBecauseParity := utl_SomeIntSlice(timeAndXyDiffs, func(timeAndXy []int) bool {
		return timeAndXy[0]%2 != timeAndXy[1]%2
	})

	if cantMovesBecauseParity {
		return "No"
	}

	return "Yes"
}

func main() {
	input, err := utl_NewInputFromReader(bufio.NewReader(io.Reader(os.Stdin)))
	utl_PanicIfErrorExist(err)
	fmt.Println(solve(input))
}
