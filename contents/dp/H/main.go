package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/mpppk/atcoder-workspace/lib"
)

const MOD = 1000000007

//func solve(H int, W int, a [][]string) int {
func solve(input *lib.Input) int {
	H, W := input.MustGetFirstAndSecondIntValue(0)
	lines := input.MustGetStringLinesFrom(1)
	var a [][]rune
	for _, line := range lines {
		a = append(a, []rune(line[0]))
	}

	m := lib.NewInt2DList(H+5, W+5, 0)
	m[0][0] = 1
	for h := 0; h < H; h++ {
		for w := 0; w < W; w++ {
			if h < H-1 && a[h+1][w] != '#' {
				m[h+1][w] += m[h][w]
				if m[h+1][w] >= MOD {
					m[h+1][w] %= MOD
				}
			}
			if w < W-1 && a[h][w+1] != '#' {
				m[h][w+1] += m[h][w]
				if m[h][w+1] >= MOD {
					m[h][w+1] %= MOD
				}
			}
		}
	}
	//fmt.Println(m)
	return m[H-1][W-1]

	//list := list2.New()
	//list.PushBack([]int{0, 0})
	//for list.Len() > 0 {
	//	elm := list.Front()
	//	pos := elm.Value.([]int)
	//	list.Remove(elm)
	//	x, y := pos[0], pos[1]
	//
	//	if x == H-1 && y == W-1 {
	//		return m[x][y]
	//	}
	//
	//	if x < H-1 && []rune(a[x+1][0])[y] != '#' {
	//		m[x+1][y] = (m[x+1][y] + 1) % MOD
	//		list.PushBack([]int{x + 1, y})
	//	}
	//	if y < W-1 && []rune(a[x][0])[y+1] != '#' {
	//		m[x][y+1] = (m[x][y+1] + 1) % MOD
	//		list.PushBack([]int{x, y + 1})
	//	}
	//}
	return 0
}

func main() {
	input := lib.MustNewInputFromReader(bufio.NewReader(io.Reader(os.Stdin)))
	fmt.Println(solve(input))
}

//func main() {
//	scanner := bufio.NewScanner(os.Stdin)
//	const initialBufSize = 4096
//	const maxBufSize = 1000000
//	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
//	scanner.Split(bufio.ScanWords)
//	var H int
//	scanner.Scan()
//	H, _ = strconv.Atoi(scanner.Text())
//	var W int
//	scanner.Scan()
//	W, _ = strconv.Atoi(scanner.Text())
//	a := make([][]string, H)
//	for i := 0; i < H; i++ {
//		a[i] = make([]string, 1)
//	}
//	for i := 0; i < H; i++ {
//		for j := 0; j < 1; j++ {
//			scanner.Scan()
//			a[i][j] = scanner.Text()
//		}
//	}
//	fmt.Println(solve(H, W, a))
//}
