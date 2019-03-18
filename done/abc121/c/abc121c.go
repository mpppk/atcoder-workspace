package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

//type Drink struct {
//	price  int64
//	maxNum int64
//}
type Drink []int64
type Drinks []Drink

func (d Drinks) Len() int {
	return len(d)
}

func (d Drinks) Less(i, j int) bool {
	a1 := d[i][0]
	a2 := d[j][0]
	return a1 < a2
}

func (d Drinks) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func solve(input *lib_Input) int64 {
	M := input.MustGetInt64Value(0, 1)
	nums := input.MustGetInt64LinesFrom(1)
	var drinks []Drink
	for _, num := range nums {
		drinks = append(drinks, Drink(num))
	}
	sort.Sort(Drinks(drinks))
	//sort.Slice(drinks, func(i, j int) bool {
	//	a1 := drinks[i][0]
	//	a2 := drinks[j][0]
	//	return a1 < a2
	//})

	moneySum := int64(0)
	drinkCnt := int64(0)
	for _, drink := range drinks {
		a := drink[0]
		b := drink[1]
		for i := int64(0); i < b; i++ {
			moneySum += a
			drinkCnt++
			if drinkCnt == M {
				return moneySum
			}
		}
	}
	panic("no answer")
}

func main() {
	input := lib_MustNewInputFromReader(bufio.NewReader(io.Reader(os.Stdin)))
	fmt.Println(solve(input))
}
