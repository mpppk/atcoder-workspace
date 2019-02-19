package generic

import (
	"reflect"
	"testing"
)

func TestSumValue(t *testing.T) {
	type args struct {
		values []AAA
	}
	tests := []struct {
		name string
		args args
		want AAA
	}{
		{
			name: "can calc sum",
			args: args{
				values: []AAA{0, 1, 2},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumAAA(tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumAAA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterValue(t *testing.T) {
	type args struct {
		values []AAA
		f      func(v AAA) bool
	}
	tests := []struct {
		name          string
		args          args
		wantNewValues []AAA
	}{
		{
			name: "can filter",
			args: args{
				values: []AAA{1, 2, 3},
				f: func(v AAA) bool {
					return v == 2
				},
			},
			wantNewValues: []AAA{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewValues := FilterAAA(tt.args.values, tt.args.f); !reflect.DeepEqual(gotNewValues, tt.wantNewValues) {
				t.Errorf("FilterAAA() = %v, want %v", gotNewValues, tt.wantNewValues)
			}
		})
	}
}

func TestUniqValue(t *testing.T) {
	type args struct {
		values []AAA
	}
	tests := []struct {
		name          string
		args          args
		wantNewValues []AAA
	}{
		{
			name: "uniq",
			args: args{
				values: []AAA{1, 2, 2, 3, 1},
			},
			wantNewValues: []AAA{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewValues := UniqAAA(tt.args.values); !reflect.DeepEqual(gotNewValues, tt.wantNewValues) {
				t.Errorf("UniqAAA() = %v, want %v", gotNewValues, tt.wantNewValues)
			}
		})
	}
}

func TestSubtractValueBy(t *testing.T) {
	type args struct {
		values1 []AAA
		values2 []AAA
		f       func(v AAA) AAA
	}
	tests := []struct {
		name          string
		args          args
		wantNewValues []AAA
		wantErr       bool
	}{
		{
			name: "SubtractAAABy",
			args: args{
				values1: []AAA{4, 5, 6},
				values2: []AAA{3, 2, 1},
				f: func(v AAA) AAA {
					return v
				},
			},
			wantNewValues: []AAA{1, 3, 5},
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNewValues, err := SubtractAAABy(tt.args.values1, tt.args.values2, tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubtractAAABy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNewValues, tt.wantNewValues) {
				t.Errorf("SubtractAAABy() = %v, want %v", gotNewValues, tt.wantNewValues)
			}
		})
	}
}

func TestRDiffValueBy(t *testing.T) {
	type args struct {
		values []AAA
		f      func(v AAA) AAA
	}
	tests := []struct {
		name          string
		args          args
		wantNewValues []AAA
		wantErr       bool
	}{
		{
			name: "RDiffAAABy",
			args: args{
				values: []AAA{1, 2, 4},
				f: func(v AAA) AAA {
					return v
				},
			},
			wantNewValues: []AAA{1, 2},
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNewValues, err := RDiffAAABy(tt.args.values, tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("RDiffAAABy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNewValues, tt.wantNewValues) {
				t.Errorf("RDiffAAABy() = %v, want %v", gotNewValues, tt.wantNewValues)
			}
		})
	}
}

func TestStringToValueLine(t *testing.T) {
	type args struct {
		line []string
	}
	tests := []struct {
		name          string
		args          args
		wantValueLine []AAA
		wantErr       bool
	}{
		{
			name: "StringToAAALine",
			args: args{
				line: []string{"1", "2", "3"},
			},
			wantValueLine: []AAA{1, 2, 3},
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValueLine, err := StringToAAALine(tt.args.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringtoValueLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotValueLine, tt.wantValueLine) {
				t.Errorf("StringtoValueLine() = %v, want %v", gotValueLine, tt.wantValueLine)
			}
		})
	}
}