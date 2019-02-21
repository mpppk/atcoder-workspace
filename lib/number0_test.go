package lib

import (
	"reflect"
	"testing"
)

func TestBitEnumeration(t *testing.T) {
	type args struct {
		digits uint
	}
	tests := []struct {
		name      string
		args      args
		wantEnums [][]bool
	}{
		{
			name: "BitEnumeration",
			args: args{
				digits: 0,
			},
			wantEnums: [][]bool{},
		},
		{
			name: "BitEnumeration",
			args: args{
				digits: 1,
			},
			wantEnums: [][]bool{{false}, {true}},
		},
		{
			name: "BitEnumeration",
			args: args{
				digits: 2,
			},
			wantEnums: [][]bool{{false, false}, {true, false}, {false, true}, {true, true}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotEnums := BitEnumeration(tt.args.digits); !reflect.DeepEqual(gotEnums, tt.wantEnums) {
				t.Errorf("BitEnumeration() = %v, want %v", gotEnums, tt.wantEnums)
			}
		})
	}
}
