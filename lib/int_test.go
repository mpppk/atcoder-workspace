package lib

import (
	"reflect"
	"testing"
)

func TestValueToBits(t *testing.T) {
	type args struct {
		value     AAA
		minDigits int
	}
	tests := []struct {
		name     string
		args     args
		wantBits []bool
	}{
		{
			name: "AAAToBits",
			args: args{
				value:     0,
				minDigits: 1,
			},
			wantBits: []bool{false},
		},
		{
			name: "AAAToBits",
			args: args{
				value:     0,
				minDigits: 3,
			},
			wantBits: []bool{false, false, false},
		},
		{
			name: "ValueToBits2",
			args: args{
				value:     8,
				minDigits: 4,
			},
			wantBits: []bool{true, false, false, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotBits := AAAToBits(tt.args.value, tt.args.minDigits); !reflect.DeepEqual(gotBits, tt.wantBits) {
				t.Errorf("AAAToBits() = %v, want %v", gotBits, tt.wantBits)
			}
		})
	}
}

func TestPrimeFactors(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantPfs []int
	}{
		{
			args:    args{n: 1},
			wantPfs: nil,
		},
		{
			args:    args{n: 2},
			wantPfs: []int{2},
		},
		{
			args:    args{n: 3},
			wantPfs: []int{3},
		},
		{
			args:    args{n: 4},
			wantPfs: []int{2, 2},
		},
		{
			args:    args{n: 57},
			wantPfs: []int{3, 19},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPfs := PrimeFactorsAAA(tt.args.n); !reflect.DeepEqual(gotPfs, tt.wantPfs) {
				t.Errorf("PrimeFactors() = %v, want %v", gotPfs, tt.wantPfs)
			}
		})
	}
}
