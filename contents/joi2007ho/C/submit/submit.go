package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type Bool2DList [][]bool

type Input struct {
	lines [][]string // Input は複数行からなる文字列から値をいい感じに取り出すためのメソッドを提供します.

}

func Area(x1, y1, x2, y2 int) int {
	return lib_PowInt(lib_AbsInt(x1-x2), 2) + lib_PowInt(lib_AbsInt(y1-y2), 2)
}
func (i *Input) GetFirstIntValue(rowIndex int) (int, error) {
	return i.GetIntValue(rowIndex, 0)
}
func (i *Input) GetIntLine(index int) ([]int, error) {
	if err := i.validateRowIndex(index); err != nil {
		return nil, err
	}
	newLine, err := lib_StringSliceToIntSlice(i.lines[index])
	if err != nil {
		return nil, fmt.Errorf("%dth index: %v", index, err)
	}
	return newLine, nil
}
func (i *Input) GetIntLinesFrom(fromIndex int) (newLines [][]int, err error) {
	for index := range i.lines {
		if index < fromIndex {
			continue
		}
		newLine, err := i.GetIntLine(index)
		if err != nil {
			return nil, err
		}
		newLines = append(newLines, newLine)
	}
	return
}
func (i *Input) GetIntValue(rowIndex, colIndex int) (int, error) {
	line, err := i.GetLine(rowIndex)
	if err != nil {
		return 0, err
	}
	if colIndex < 0 || colIndex >= len(line) {
		return 0, fmt.Errorf("Invalid col index: %v ", colIndex)
	}
	v, err := strconv.ParseInt(line[colIndex], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to convert string to int: %v, %v", line[colIndex], err)
	}
	return int(v), nil
}
func (i *Input) GetLine(index int) ([]string, error) {
	if err := i.validateRowIndex(index); err != nil {
		return nil, err
	}
	return i.lines[index], nil
}
func (i *Input) MustGetFirstIntValue(rowIndex int) int {
	_v0, _err := i.GetFirstIntValue(rowIndex)
	if _err != nil {
		panic(_err)
	}
	return _v0
}
func (i *Input) MustGetIntLinesFrom(fromIndex int) (newLines [][]int) {
	newLines, err := i.GetIntLinesFrom(fromIndex)
	if err != nil {
		panic(err)
	}
	return newLines
}
func (i *Input) validateRowIndex(index int) error {
	if index >= len(i.lines) {
		return errors.New(fmt.Sprintf("index(%d) is larger than lines(%d)", index, len(i.lines)))
	}
	if index < 0 {
		return errors.New(fmt.Sprintf("index is under zero: %d", index))
	}
	return nil
}
func lib_AbsInt(value int) int {
	return int(math.Abs(float64(value)))
}
func lib_CopyBool(values []bool) []bool {
	dst := make([]bool, len(values))
	copy(dst, values)
	return dst
}
func lib_MaxInt(values ...int) (max int, err error) {
	if len(values) == 0 {
		return 0, errors.New("empty slice is given")
	}
	max = values[0]
	for _, value := range values {
		if max < value {
			max = value
		}
	}
	return
}
func lib_MinInt(values ...int) (min int, err error) {
	if len(values) == 0 {
		return 0, errors.New("empty slice is given")
	}
	min = values[0]
	for _, value := range values {
		if min > value {
			min = value
		}
	}
	return
}
func lib_MustMaxInt(values ...int) (max int) {
	max, err := lib_MaxInt(values...)
	if err != nil {
		panic(err)
	}
	return max
}
func lib_MustMinInt(values ...int) (min int) {
	min, err := lib_MinInt(values...)
	if err != nil {
		panic(err)
	}
	return min
}
func lib_MustNewInputFromReader(reader *bufio.Reader) *Input {
	_v0, _err := lib_NewInputFromReader(reader)
	if _err != nil {
		panic(_err)
	}
	return _v0
}
func lib_New2DimBoolSliceWithInitialValue(row, col int, initialValue bool) [][]bool {
	ret := make([][]bool, row)
	values := lib_NewBoolSliceWithInitialValue(col, initialValue)
	for r := 0; r < row; r++ {
		ret[r] = lib_CopyBool(values)
	}
	return ret
}
func lib_NewBool2DList(row, col int, initialValue bool) Bool2DList {
	return lib_New2DimBoolSliceWithInitialValue(row, col, initialValue)
}
func lib_NewBoolSliceWithInitialValue(length int, initialValue bool) []bool {
	ret := make([]bool, length)
	for i := 0; i < length; i++ {
		ret[i] = initialValue
	}
	return ret
}
func lib_NewInputFromReader(reader *bufio.Reader) (*Input, error) {
	lines, err := lib_toLinesFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to create new Input from reader: %v", err)
	}
	return &Input{lines: lines}, nil
}
func lib_PowInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
func lib_StringSliceToIntSlice(line []string) (ValueLine []int, err error) {
	newLine, err := lib_toSpecificBitIntLine(line, 64)
	if err != nil {
		return nil, err
	}
	for _, v := range newLine {
		ValueLine = append(ValueLine, int(v))
	}
	return
}
func lib_TernaryOPInt(ok bool, v1, v2 int) int {
	if ok {
		return v1
	}
	return v2
}
func lib_TrimSpaceAndNewLineCodeAndTab(s string) string {
	return strings.TrimFunc(s, func(r rune) bool {
		return r == ' ' || r == '\r' || r == '\n' || r == '\t'
	})
}
func lib_readLineAsChunks(reader *bufio.Reader) (chunks []string, err error) {
	for {
		chunk, isPrefix, err := reader.ReadLine()
		if err != nil {
			return nil, err
		}
		chunks = append(chunks, string(chunk))
		if !isPrefix {
			return chunks, nil
		}
	}
}
func lib_toLinesFromReader(reader *bufio.Reader) (lines [][]string, err error) {
	for {
		chunks, err := lib_readLineAsChunks(reader)
		if err == io.EOF {
			return lines, nil
		}
		if err != nil {
			return nil, fmt.Errorf("failed to read line from reader: %v", err)
		}
		lineStr := lib_TrimSpaceAndNewLineCodeAndTab(strings.Join(chunks, ""))
		line := strings.Split(lineStr, " ")
		lines = append(lines, line)
	}
}
func lib_toSpecificBitIntLine(line []string, bitSize int) (intLine []int64, err error) {
	for j, v := range line {
		intV, err := strconv.ParseInt(v, 10, bitSize)
		if err != nil {
			return nil, fmt.Errorf(fmt.Sprintf("%dth value: %v", j, err.Error()))
		}
		intLine = append(intLine, intV)
	}
	return intLine, nil
}
func main() {
	input := lib_MustNewInputFromReader(bufio.NewReader(io.Reader(os.Stdin)))
	fmt.Println(solve(input))
}
func solve(input *Input) int {
	N := input.MustGetFirstIntValue(0)
	pillars := input.MustGetIntLinesFrom(1)
	list := lib_NewBool2DList(5000+5, 5000+5, false)
	for _, pillar := range pillars {
		x, y := pillar[0], pillar[1]
		list[x][y] = true
	}
	maxArea := 0
	for i := 0; i < N-1; i++ {
		for j := i; j < N; j++ {
			x1, y1, x2, y2 := pillars[i][0], pillars[i][1], pillars[j][0], pillars[j][1]
			diffX, diffY := x1-x2, y1-y2
			if lib_MustMinInt(x1-diffY, y1+diffX, x2-diffY, y2+diffX) >= 0 && lib_MustMaxInt(x1-diffY, y1+diffX, x2-diffY, y2+diffX) <= 5000 && list[x1-diffY][y1+diffX] && list[x2-diffY][y2+diffX] {
				area := Area(x1, y1, x2, y2)
				maxArea = lib_TernaryOPInt(maxArea < area, area, maxArea)
			}
			if lib_MustMinInt(x1+diffY, y1-diffX, x2+diffY, y2-diffX) >= 0 && lib_MustMaxInt(x1+diffY, y1-diffX, x2+diffY, y2-diffX) <= 5000 && list[x1+diffY][y1-diffX] && list[x2+diffY][y2-diffX] {
				area := Area(x1, y1, x2, y2)
				maxArea = lib_TernaryOPInt(maxArea < area, area, maxArea)
			}
		}
	}
	return maxArea
}
