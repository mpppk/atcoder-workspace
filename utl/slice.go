package utl

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "Value=int,int16,int32,int64,float32,float64"

import "github.com/cheekybits/genny/generic"

type Value generic.Number

func SumValue(values []Value) Value {
	var sum Value = 0
	for _, value := range values {
		sum += value
	}
	return sum
}
