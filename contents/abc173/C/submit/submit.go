package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	lines [][]string // Input は複数行からなる文字列から値をいい感じに取り出すためのメソッドを提供します.

}
type bitGenerator struct {
	cur       int
	max       int
	minDigits int
}

func (b *bitGenerator) Next() []bool {
	if b.cur == b.max {
		return nil
	}
	r := lib_IntToBits(b.cur, b.minDigits)
	b.cur++
	return r
}
func (i *Input) GetFirstAndSecondIntValue(rowIndex int) (int, int, error) {
	first, err := i.GetIntValue(rowIndex, 0)
	if err != nil {
		return first, 0, err
	}
	second, err := i.GetIntValue(rowIndex, 1)
	return first, second, err
}
func (i *Input) GetFromFirstToThirdIntValue(rowIndex int) (int, int, int, error) {
	first, second, err := i.GetFirstAndSecondIntValue(rowIndex)
	if err != nil {
		return first, second, 0, err
	}
	third, err := i.GetIntValue(rowIndex, 2)
	return first, second, third, err
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
func (i *Input) GetStringLinesFrom(fromIndex int) (newLines [][]string, err error) {
	for index := range i.lines {
		if index < fromIndex {
			continue
		}
		newLine, err := i.GetLine(index)
		if err != nil {
			return nil, err
		}
		newLines = append(newLines, newLine)
	}
	return
}
func (i *Input) MustGetFromFirstToThirdIntValue(rowIndex int) (int, int, int) {
	_v0, _v1, _v2, _err := i.GetFromFirstToThirdIntValue(rowIndex)
	if _err != nil {
		panic(_err)
	}
	return _v0, _v1, _v2
}
func (i *Input) MustReadAsStringGridFrom(fromIndex int) [][]string {
	_v0, _err := i.ReadAsStringGridFrom(fromIndex)
	if _err != nil {
		panic(_err)
	}
	return _v0
}
func (i *Input) ReadAsStringGridFrom(fromIndex int) ([][]string, error) {
	lines, err := i.GetStringLinesFrom(fromIndex)
	if err != nil {
		return nil, err
	}
	var m [][]string
	for _, line := range lines {
		if len(line) > 1 {
			return nil, fmt.Errorf("unexpected length line: %v", line)
		}
		var mLine []string
		for _, r := range line[0] {
			mLine = append(mLine, string(r))
		}
		m = append(m, mLine)
	}
	return m, nil
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
func lib_IntToBits(v int, minDigits int) (bts []bool) {
	minDigits = lib_MustMaxInt(minDigits, bits.Len64(uint64(v)))
	for b := 0; b < minDigits; b++ {
		bts = append(bts, (v>>b)&1 == 1)
	}
	return
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
func lib_MustMaxInt(values ...int) (max int) {
	max, err := lib_MaxInt(values...)
	if err != nil {
		panic(err)
	}
	return max
}
func lib_MustNewInputFromReader(reader *bufio.Reader) *Input {
	_v0, _err := lib_NewInputFromReader(reader)
	if _err != nil {
		panic(_err)
	}
	return _v0
}
func lib_NewBitGenerator(minDigits int) *bitGenerator {
	return &bitGenerator{cur: 0, max: (1 << minDigits) - 1, minDigits: minDigits}
}
func lib_NewInputFromReader(reader *bufio.Reader) (*Input, error) {
	lines, err := lib_toLinesFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to create new Input from reader: %v", err)
	}
	return &Input{lines: lines}, nil
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
func main() {
	input := lib_MustNewInputFromReader(bufio.NewReader(io.Reader(os.Stdin)))
	fmt.Println(solve(input))
}
func solve(input *Input) int {
	H, W, K := input.MustGetFromFirstToThirdIntValue(0)
	board := input.MustReadAsStringGridFrom(1)
	cnt := 0
	hGen := lib_NewBitGenerator(H)
	for hb := hGen.Next(); hb != nil; hb = hGen.Next() {
		wGen := lib_NewBitGenerator(W)
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
