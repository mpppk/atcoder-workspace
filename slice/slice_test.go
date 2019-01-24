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
