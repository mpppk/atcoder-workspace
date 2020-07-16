package lib

import (
	"reflect"
	"testing"
)

func Test_bitGenerator_Next(t *testing.T) {
	tests := []struct {
		name      string
		minDigits int
		want      [][]bool
	}{
		{
			minDigits: 2,
			want: [][]bool{
				{false, false},
				{true, false},
				{false, true},
				{true, true},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBitGenerator(tt.minDigits)
			var got [][]bool
			for curB := b.Next(); curB != nil; curB = b.Next() {
				got = append(got, curB)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
