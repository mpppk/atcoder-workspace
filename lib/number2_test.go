package lib

import (
	"reflect"
	"testing"
)

func TestAAAliceToBBBBSlice(t *testing.T) {
	type args struct {
		values []AAA
	}
	tests := []struct {
		name          string
		args          args
		wantNewValues []BBB
	}{
		{
			name: "ZZZTypeSliceToTypeSlice",
			args: args{
				values: []AAA{1, 2, 3},
			},
			wantNewValues: []BBB{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewValues := AAASliceToBBBSlice(tt.args.values); !reflect.DeepEqual(gotNewValues, tt.wantNewValues) {
				t.Errorf("AAASliceToBBBSlice() = %v, want %v", gotNewValues, tt.wantNewValues)
			}
		})
	}
}

func TestMaxAAAByBBBSlice(t *testing.T) {
	type args struct {
		values [][]BBB
		f      func(vs []BBB) AAA
	}
	tests := []struct {
		name    string
		args    args
		wantMax AAA
		wantErr bool
	}{
		{
			name: "ZZZTypeSliceToTypeSlice",
			args: args{
				values: [][]BBB{{1, 2, 3}, {4, 5, 6}},
				f: func(vs []BBB) AAA {
					sum := 0
					for _, v := range vs {
						sum += int(v)
					}
					return AAA(sum)
				},
			},
			wantMax: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMax, err := MaxAAAByBBBSlice(tt.args.values, tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaxAAAByBBBSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotMax, tt.wantMax) {
				t.Errorf("MaxAAAByBBBSlice() = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}

func TestSumAAAByBBB(t *testing.T) {
	type args struct {
		values []AAA
		f      func(v AAA) BBB
	}
	tests := []struct {
		name string
		args args
		want BBB
	}{
		{
			name: "SumAAAByBBB",
			args: args{
				values: []AAA{1, 2, 3},
				f: func(a AAA) BBB {
					return BBB(int(a) * int(a))
				},
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumAAAByBBB(tt.args.values, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumAAAByBBB() = %v, want %v", got, tt.want)
			}
		})
	}
}
