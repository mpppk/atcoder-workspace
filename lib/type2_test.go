package lib

import (
	"reflect"
	"testing"
	"unicode/utf8"
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

func TestYYYToZZZ2DMap_ChBy(t *testing.T) {
	type args struct {
		key1   YYY
		key2   YYY
		f      func(cur, new ZZZ) bool
		value  ZZZ
		values []ZZZ
	}
	tests := []struct {
		name                  string
		m                     YYYToZZZ2DMap
		args                  args
		wantReplaced          bool
		wantValueAlreadyExist bool
		wantValue             ZZZ
	}{
		{
			m: YYYToZZZ2DMap{
				0: {0: "a", 1: "b"},
				1: {0: "c", 1: "d"},
			},
			args: args{
				key1: 0,
				key2: 0,
				f: func(cur, new ZZZ) bool {
					return utf8.RuneCountInString(cur.(string)) < utf8.RuneCountInString(new.(string))
				},
				value: "ab",
			},
			wantReplaced:          true,
			wantValueAlreadyExist: true,
			wantValue:             "ab",
		},
		{
			m: YYYToZZZ2DMap{
				0: {0: "a", 1: "b"},
				1: {0: "c", 1: "d"},
			},
			args: args{
				key1: 0,
				key2: 0,
				f: func(cur, new ZZZ) bool {
					return utf8.RuneCountInString(cur.(string)) < utf8.RuneCountInString(new.(string))
				},
				value: "",
			},
			wantReplaced:          false,
			wantValueAlreadyExist: true,
			wantValue:             "a",
		},
		{
			m: YYYToZZZ2DMap{
				0: {0: "a", 1: "b"},
				1: {0: "c", 1: "d"},
			},
			args: args{
				key1: 2,
				key2: 1,
				f: func(cur, new ZZZ) bool {
					return utf8.RuneCountInString(cur.(string)) < utf8.RuneCountInString(new.(string))
				},
				value: "e",
			},
			wantReplaced:          true,
			wantValueAlreadyExist: false,
			wantValue:             "e",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotReplaced, gotValueAlreadyExist := tt.m.ChBy(tt.args.key1, tt.args.key2, tt.args.f, tt.args.value, tt.args.values...)
			if gotReplaced != tt.wantReplaced {
				t.Errorf("ChBy() gotReplaced = %v, want %v", gotReplaced, tt.wantReplaced)
			}
			if gotValueAlreadyExist != tt.wantValueAlreadyExist {
				t.Errorf("ChBy() gotValueAlreadyExist = %v, want %v", gotValueAlreadyExist, tt.wantValueAlreadyExist)
			}
			if gotValue := tt.m[tt.args.key1][tt.args.key2]; gotValue != tt.wantValue {
				t.Errorf("ChBy() got value = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}
