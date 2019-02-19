package generic

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
