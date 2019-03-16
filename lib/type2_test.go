package lib

import (
	"reflect"
	"testing"
)

func TestMemoizeYYYToZZZ(t *testing.T) {
	funcCalledNum := 0
	type args struct {
		f func(v1 YYY) ZZZ
	}
	tests := []struct {
		name    string
		args    args
		testArg YYY
		want    ZZZ
	}{
		{
			name: "MemoizeYYY2ToZZZ",
			args: args{
				f: func(v1 YYY) ZZZ {
					funcCalledNum++
					v1Int := v1.(int)
					return ZZZ(v1Int + 1)
				},
			},
			testArg: 1,
			want:    2,
		},
	}
	for _, tt := range tests {
		funcCalledNum = 0
		t.Run(tt.name, func(t *testing.T) {
			generatedFunc := MemoizeYYYToZZZ(tt.args.f)
			if got := generatedFunc(tt.testArg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MemoizeYYY2ToZZZ() = %v, want %v", got, tt.want)
			}
			if got := generatedFunc(tt.testArg); !reflect.DeepEqual(got, tt.want) {
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

func TestMemoizeYYY5ToZZZ(t *testing.T) {
	funcCalledNum := 0
	type args struct {
		f func(v1, v2, v3, v4, v5 YYY) ZZZ
	}
	tests := []struct {
		name     string
		args     args
		testArgs [5]YYY
		want     ZZZ
	}{
		{
			name: "MemoizeYYY2ToZZZ",
			args: args{
				f: func(v1, v2, v3, v4, v5 YYY) ZZZ {
					funcCalledNum++
					v1Int := v1.(int)
					v2Int := v2.(int)
					v3Int := v3.(int)
					v4Int := v4.(int)
					v5Int := v5.(int)
					return ZZZ(v1Int + v2Int + v3Int + v4Int + v5Int)
				},
			},
			testArgs: [5]YYY{1, 2, 3, 4, 5},
			want:     15,
		},
	}
	for _, tt := range tests {
		funcCalledNum = 0
		t.Run(tt.name, func(t *testing.T) {
			generatedFunc := MemoizeYYY5ToZZZ(tt.args.f)
			if got := generatedFunc(tt.testArgs[0], tt.testArgs[1], tt.testArgs[2], tt.testArgs[3], tt.testArgs[4]); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MemoizeYYY2ToZZZ() = %v, want %v", got, tt.want)
			}
			if got := generatedFunc(tt.testArgs[0], tt.testArgs[1], tt.testArgs[2], tt.testArgs[3], tt.testArgs[4]); !reflect.DeepEqual(got, tt.want) {
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
