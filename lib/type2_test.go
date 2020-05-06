package lib

import (
	"reflect"
	"testing"
)

func TestMemoizeYYY2ToZZZ(t *testing.T) {
	funcCalledNum := 0
	type args struct {
		f func(v1, v2 YYY) ZZZ
	}
	tests := []struct {
		name     string
		args     args
		testArgs [2]YYY
		want     ZZZ
	}{
		{
			name: "MemoizeYYY2ToZZZ",
			args: args{
				f: func(v1, v2 YYY) ZZZ {
					funcCalledNum++
					v1Int := v1.(int)
					v2Int := v2.(int)
					return ZZZ(v1Int + v2Int)
				},
			},
			testArgs: [2]YYY{1, 2},
			want:     3,
		},
	}
	for _, tt := range tests {
		funcCalledNum = 0
		t.Run(tt.name, func(t *testing.T) {
			generatedFunc := MemoizeYYY2ToZZZ(tt.args.f)
			if got := generatedFunc(tt.testArgs[0], tt.testArgs[1]); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MemoizeYYY2ToZZZ() = %v, want %v", got, tt.want)
			}
			if got := generatedFunc(tt.testArgs[0], tt.testArgs[1]); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MemoizeYYY2ToZZZ() = %v, want %v", got, tt.want)
			}
			if funcCalledNum != 1 {
				t.Errorf("funcCalledNum = %v, want %v", funcCalledNum, 1)
			}
		})
	}

	if funcCalledNum != 1 {

	}
}
