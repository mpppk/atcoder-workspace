package slice

import (
	"reflect"
	"testing"
)

func TestReverseType(t *testing.T) {
	type args struct {
		values []Type
	}
	tests := []struct {
		name string
		args args
		want []Type
	}{
		{
			name: "can reverse int slice",
			args: args{
				values: []Type{1, 2, 3},
			},
			want: []Type{3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseType(tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
