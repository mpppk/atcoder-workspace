package slice

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

func TestMapValue(t *testing.T) {
	type args struct {
		values []Value
		f      func(v Value) Value
	}
	tests := []struct {
		name          string
		args          args
		wantNewValues []Value
	}{
		{
			name: "can increment int",
			args: args{
				values: []Value{1, 2, 3},
				f: func(v Value) Value {
					return v + 1
				},
			},
			wantNewValues: []Value{2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewValues := MapValue(tt.args.values, tt.args.f); !reflect.DeepEqual(gotNewValues, tt.wantNewValues) {
				t.Errorf("MapValue() = %v, want %v", gotNewValues, tt.wantNewValues)
			}
		})
	}
}

func TestReverseValue(t *testing.T) {
	type args struct {
		values []Value
	}
	tests := []struct {
		name string
		args args
		want []Value
	}{
		{
			name: "can reverse int slice",
			args: args{
				values: []Value{1, 2, 3},
			},
			want: []Value{3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseValue(tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
