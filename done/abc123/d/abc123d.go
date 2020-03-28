package main

import (
	"bufio"
	"fmt"
	"github.com/mpppk/atcoder/done/abc123/d"
	"io"
	"math"
	"os"
	"sort"
)

type Cake struct {
	price  int64
	candle int8
}

type Cakes []Cake

func (c Cakes) Len() int {
	return len(c)
}

func (c Cakes) Less(i, j int) bool {
	return c[i].price < c[j].price
}

func (c Cakes) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func solve(input *d.lib_Input) []int64 {
	K := input.MustGetIntLine(0)[3]
	cake1Prices := input.MustGetInt64Line(1)
	cake2Prices := input.MustGetInt64Line(2)
	cake3Prices := input.MustGetInt64Line(3)

	sort.Slice(cake1Prices, func(i, j int) bool {
		return cake1Prices[i] > cake1Prices[j]
	})
	sort.Slice(cake2Prices, func(i, j int) bool {
		return cake2Prices[i] > cake2Prices[j]
	})
	sort.Slice(cake3Prices, func(i, j int) bool {
		return cake3Prices[i] > cake3Prices[j]
	})

	cake1s := []int64{cake1Prices[0]}
	cake2s := []int64{cake2Prices[0]}
	cake3s := []int64{cake3Prices[0]}
	baseCake1Price := cake1Prices[0]
	baseCake2Price := cake2Prices[0]
	baseCake3Price := cake3Prices[0]
	cake1Prices = cake1Prices[1:]
	cake2Prices = cake2Prices[1:]
	cake3Prices = cake3Prices[1:]
	var otherCakesA, otherCakesB []int64
	sumPrices := []int64{cake1s[0] + cake2s[0] + cake3s[0]}

	for {
		c1, c2, c3 := int64(math.MaxInt64), int64(math.MaxInt64), int64(math.MaxInt64)
		if len(cake1Prices) > 0 {
			c1 = int64(math.Abs(float64(cake1Prices[0] - baseCake1Price)))
			//c1 = int64(math.Abs(float64(cake1Prices[0] - cake1s[len(cake1s)-1])))
		}
		if len(cake2Prices) > 0 {
			c2 = int64(math.Abs(float64(cake2Prices[0] - baseCake2Price)))
			//c2 = int64(math.Abs(float64(cake2Prices[0] - cake2s[len(cake2s)-1])))
		}
		if len(cake3Prices) > 0 {
			c3 = int64(math.Abs(float64(cake3Prices[0] - baseCake3Price)))
			//c3 = int64(math.Abs(float64(cake3Prices[0] - cake3s[len(cake3s)-1])))
		}

		var c1Flag, c2Flag, c3Flag bool
		if c1 <= c2 && c1 <= c3 && len(cake1Prices) > 0 {
			c1Flag = true
			newCakePrice := cake1Prices[0]
			otherCakesA, otherCakesB = cake2s, cake3s
			for _, otherCakeA := range otherCakesA {
				for _, otherCakeB := range otherCakesB {
					sumPrice := newCakePrice + otherCakeA + otherCakeB
					sumPrices = append(sumPrices, sumPrice)
				}
			}
		}
		if c2 <= c1 && c2 <= c3 && len(cake2Prices) > 0 {
			c2Flag = true
			newCakePrice := cake2Prices[0]
			otherCakesA, otherCakesB = cake1s, cake3s
			for _, otherCakeA := range otherCakesA {
				for _, otherCakeB := range otherCakesB {
					sumPrice := newCakePrice + otherCakeA + otherCakeB
					sumPrices = append(sumPrices, sumPrice)
				}
			}
		}
		if c3 <= c1 && c3 <= c2 && len(cake3Prices) > 0 {
			c3Flag = true
			newCakePrice := cake3Prices[0]
			otherCakesA, otherCakesB = cake1s, cake2s
			for _, otherCakeA := range otherCakesA {
				for _, otherCakeB := range otherCakesB {
					sumPrice := newCakePrice + otherCakeA + otherCakeB
					sumPrices = append(sumPrices, sumPrice)
				}
			}
		}

		if c1Flag {
			cake1s = append(cake1s, cake1Prices[0])
			cake1Prices = cake1Prices[1:]
		}
		if c2Flag {
			cake2s = append(cake2s, cake2Prices[0])
			cake2Prices = cake2Prices[1:]
		}
		if c3Flag {
			cake3s = append(cake3s, cake3Prices[0])
			cake3Prices = cake3Prices[1:]
		}

		if len(sumPrices) >= K {
			sort.Slice(sumPrices, func(i, j int) bool {
				return sumPrices[i] > sumPrices[j]
			})
			return sumPrices[:K]
		}
	}
}

func main() {
	input := d.lib_MustNewInputFromReader(bufio.NewReader(io.Reader(os.Stdin)))
	fmt.Println(solve(input))
}
