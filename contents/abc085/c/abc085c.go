package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(input *utl_Input) string {
	n := input.MustGetFirstIntValue(0)
	yenSum := input.MustGetLastIntValue(0)

	for manNum := 0; manNum <= n; manNum++ {
		for goNum := 0; goNum <= n-manNum; goNum++ {
			senNum := n - manNum - goNum
			if (manNum*10000 + goNum*5000 + senNum*1000) == yenSum {
				return fmt.Sprintf("%d %d %d", manNum, goNum, senNum)
			}
		}
	}
	return "-1 -1 -1"

}

func main() {
	input := utl_NewInput(bufio.NewScanner(os.Stdin))
	fmt.Println(solve(input))
}
