package lib

import (
	"reflect"
	"testing"
)

func TestToDigitSliceValue(t *testing.T) {
	type args struct {
		n AAA
	}
	tests := []struct {
		name       string
		args       args
		wantDigits []int8
	}{
		{
			name: "can convert int to digit slice",
			args: args{
				n: 100,
			},
			wantDigits: []int8{1, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDigits := ToDigitSliceAAA(tt.args.n); !reflect.DeepEqual(gotDigits, tt.wantDigits) {
				t.Errorf("ToDigitSliceAAA() = %v, want %v", gotDigits, tt.wantDigits)
			}
		})
	}
}

func TestGetEachDigitSumValue(t *testing.T) {
	type args struct {
		n AAA
	}
	tests := []struct {
		name    string
		args    args
		wantSum AAA
	}{
		{
			name: "can sum int slice digits",
			args: args{
				n: 100,
			},
			wantSum: 1,
		},
		{
			name: "can sum int slice digits2",
			args: args{
				n: 0,
			},
			wantSum: 0,
		},
		{
			name: "can sum int slice digits3",
			args: args{
				n: 123,
			},
			wantSum: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := GetEachDigitSumAAA(tt.args.n); !reflect.DeepEqual(gotSum, tt.wantSum) {
				t.Errorf("GetEachDigitSumAAA() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}

func TestDigitsToValue(t *testing.T) {
	type args struct {
		digits []int8
	}
	tests := []struct {
		name string
		args args
		want AAA
	}{
		{
			name: "DigitsToAAA",
			args: args{
				digits: []int8{1, 2, 3},
			},
			want: AAA(123),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DigitsToAAA(tt.args.digits); got != tt.want {
				t.Errorf("DigitsToAAA() = %v, want %v", got, tt.want)
			}
		})
	}
}
