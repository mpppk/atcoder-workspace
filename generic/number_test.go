package generic

import (
	"reflect"
	"testing"
)

func TestSumValue(t *testing.T) {
	type args struct {
		values []Value
	}
	tests := []struct {
		name string
		args args
		want Value
	}{
		{
			name: "can calc sum",
			args: args{
				values: []Value{0, 1, 2},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumValue(tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterValue(t *testing.T) {
	type args struct {
		values []Value
		f      func(v Value) bool
	}
	tests := []struct {
		name          string
		args          args
		wantNewValues []Value
	}{
		{
			name: "can filter",
			args: args{
				values: []Value{1, 2, 3},
				f: func(v Value) bool {
					return v == 2
				},
			},
			wantNewValues: []Value{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewValues := FilterValue(tt.args.values, tt.args.f); !reflect.DeepEqual(gotNewValues, tt.wantNewValues) {
				t.Errorf("FilterValue() = %v, want %v", gotNewValues, tt.wantNewValues)
			}
		})
	}
}

func TestUniqValue(t *testing.T) {
	type args struct {
		values []Value
	}
	tests := []struct {
		name          string
		args          args
		wantNewValues []Value
	}{
		{
			name: "uniq",
			args: args{
				values: []Value{1, 2, 2, 3, 1},
			},
			wantNewValues: []Value{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewValues := UniqValue(tt.args.values); !reflect.DeepEqual(gotNewValues, tt.wantNewValues) {
				t.Errorf("UniqValue() = %v, want %v", gotNewValues, tt.wantNewValues)
			}
		})
	}
}

func TestSubtractValueBy(t *testing.T) {
	type args struct {
		values1 []Value
		values2 []Value
		f       func(v Value) Value
	}
	tests := []struct {
		name          string
		args          args
		wantNewValues []Value
		wantErr       bool
	}{
		{
			name: "SubtractValueBy",
			args: args{
				values1: []Value{4, 5, 6},
				values2: []Value{3, 2, 1},
				f: func(v Value) Value {
					return v
				},
			},
			wantNewValues: []Value{1, 3, 5},
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNewValues, err := SubtractValueBy(tt.args.values1, tt.args.values2, tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubtractValueBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNewValues, tt.wantNewValues) {
				t.Errorf("SubtractValueBy() = %v, want %v", gotNewValues, tt.wantNewValues)
			}
		})
	}
}

func TestRDiffValueBy(t *testing.T) {
	type args struct {
		values []Value
		f      func(v Value) Value
	}
	tests := []struct {
		name          string
		args          args
		wantNewValues []Value
		wantErr       bool
	}{
		{
			name: "RDiffValueBy",
			args: args{
				values: []Value{1, 2, 4},
				f: func(v Value) Value {
					return v
				},
			},
			wantNewValues: []Value{1, 2},
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNewValues, err := RDiffValueBy(tt.args.values, tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("RDiffValueBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNewValues, tt.wantNewValues) {
				t.Errorf("RDiffValueBy() = %v, want %v", gotNewValues, tt.wantNewValues)
			}
		})
	}
}
