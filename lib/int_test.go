package lib

import (
	"reflect"
	"testing"
)

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

func TestAAASliceToMap(t *testing.T) {
	type args struct {
		values []AAA
	}
	tests := []struct {
		name string
		args args
		want map[AAA]struct{}
	}{
		{
			args: args{values: []AAA{1, 2, 3}},
			want: map[AAA]struct{}{1: {}, 2: {}, 3: {}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AAASliceToMap(tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AAASliceToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
