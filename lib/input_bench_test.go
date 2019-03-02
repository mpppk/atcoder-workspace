package lib

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"strings"
	"testing"
)

func generateLines(rowNum, colNum int) *bytes.Buffer {
	var sb strings.Builder
	for i := 0; i < rowNum; i++ {
		for j := 0; j < colNum; j++ {
			sb.WriteString("a ")
		}
		sb.WriteString("\n")
	}
	return bytes.NewBufferString(sb.String())
}

func pow10(n int) int {
	return int(math.Pow(10, float64(n)))
}

func BenchmarkFmtScan(b *testing.B) {
	cases := []struct {
		colNum int
	}{
		{pow10(4)},
		{pow10(5)},
	}

	for _, c := range cases {
		b.Run(fmt.Sprintf("col: %d", c.colNum), func(b *testing.B) {
			benchmarkFmtScan(b, c.colNum)
		})
	}
}

func benchmarkFmtScan(b *testing.B, col int) {
	buf := generateLines(b.N, col)
	b.ResetTimer()
	var a string
	for rowIndex := 0; rowIndex < b.N; rowIndex++ {
		for colIndex := 0; colIndex < col; colIndex++ {
			if _, err := fmt.Fscan(buf, &a); err != nil {
				panic(err)
			}
		}
	}
}

func BenchmarkBufioScanner(b *testing.B) {
	var sc = bufio.NewScanner(os.Stdin)

	var s, t string
	if sc.Scan() {
		s = sc.Text()
	}
	if sc.Scan() {
		t = sc.Text()
	}
	fmt.Println(s)
	fmt.Println(t)
}
