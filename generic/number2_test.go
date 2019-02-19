package generic

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
