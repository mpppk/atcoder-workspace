package generic

import (
	"reflect"
	"testing"
)

func TestZZZSliceToTypeSlice(t *testing.T) {
	type args struct {
		values []ZZZ
	}
	tests := []struct {
		name          string
		args          args
		wantNewValues []Type
	}{
		{
			name: "ZZZTypeSliceToTypeSlice",
			args: args{
				values: []ZZZ{1, 2, 3},
			},
			wantNewValues: []Type{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewValues := ZZZSliceToTypeSlice(tt.args.values); !reflect.DeepEqual(gotNewValues, tt.wantNewValues) {
				t.Errorf("ZZZSliceToTypeSlice() = %v, want %v", gotNewValues, tt.wantNewValues)
			}
		})
	}
}
