package generic

import (
	"reflect"
	"testing"
)

func TestToDigitSliceValue(t *testing.T) {
	type args struct {
		n Value
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
			if gotDigits := ToDigitSliceValue(tt.args.n); !reflect.DeepEqual(gotDigits, tt.wantDigits) {
				t.Errorf("ToDigitSliceValue() = %v, want %v", gotDigits, tt.wantDigits)
			}
		})
	}
}

func TestGetEachDigitSumValue(t *testing.T) {
	type args struct {
		n Value
	}
	tests := []struct {
		name    string
		args    args
		wantSum Value
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
			if gotSum := GetEachDigitSumValue(tt.args.n); !reflect.DeepEqual(gotSum, tt.wantSum) {
				t.Errorf("GetEachDigitSumValue() = %v, want %v", gotSum, tt.wantSum)
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
		want int
	}{
		{
			name: "DigitsToValue",
			args: args{
				digits: []int8{1, 2, 3},
			},
			want: 123,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DigitsToValue(tt.args.digits); got != tt.want {
				t.Errorf("DigitsToValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
