package lib

import (
	"reflect"
	"testing"
)

func TestReverseType(t *testing.T) {
	type args struct {
		values []ZZZ
	}
	tests := []struct {
		name string
		args args
		want []ZZZ
	}{
		{
			name: "can reverse int slice",
			args: args{
				values: []ZZZ{1, 2, 3},
			},
			want: []ZZZ{3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseZZZ(tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapType(t *testing.T) {
	type args struct {
		values []ZZZ
		f      func(v ZZZ) ZZZ
	}
	tests := []struct {
		name          string
		args          args
		wantNewValues []ZZZ
	}{
		{
			name: "can increment int",
			args: args{
				values: []ZZZ{1, 2, 3},
				f: func(v ZZZ) ZZZ {
					intV, _ := v.(int)
					return intV + 1
				},
			},
			wantNewValues: []ZZZ{2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewValues := MapZZZ(tt.args.values, tt.args.f); !reflect.DeepEqual(gotNewValues, tt.wantNewValues) {
				t.Errorf("MapValue() = %v, want %v", gotNewValues, tt.wantNewValues)
			}
		})
	}
}

func TestZipType(t *testing.T) {
	type args struct {
		valuesList [][]ZZZ
	}
	tests := []struct {
		name              string
		args              args
		wantNewValuesList [][]ZZZ
		wantErr           bool
	}{
		{
			name: "ZipZZZ",
			args: args{
				valuesList: [][]ZZZ{{1, 2, 3}, {4, 5, 6}},
			},
			wantNewValuesList: [][]ZZZ{{1, 4}, {2, 5}, {3, 6}},
			wantErr:           false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNewValuesList, err := ZipZZZ(tt.args.valuesList...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZipZZZ() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNewValuesList, tt.wantNewValuesList) {
				t.Errorf("ZipZZZ() = %v, want %v", gotNewValuesList, tt.wantNewValuesList)
			}
		})
	}
}

func TestChunkByBits(t *testing.T) {
	type args struct {
		values []ZZZ
		bits   []bool
	}
	tests := []struct {
		name          string
		args          args
		wantNewValues [][]ZZZ
		wantErr       bool
	}{
		{
			name: "ChunkZZZByBits",
			args: args{
				values: []ZZZ{1, 2, 3},
				bits:   []bool{true, false},
			},
			wantNewValues: [][]ZZZ{{1}, {2, 3}},
			wantErr:       false,
		},
		{
			name: "ChunkByBits2",
			args: args{
				values: []ZZZ{1, 2, 3},
				bits:   []bool{false, true},
			},
			wantNewValues: [][]ZZZ{{1, 2}, {3}},
			wantErr:       false,
		},
		{
			name: "different length",
			args: args{
				values: []ZZZ{1, 2, 3, 4},
				bits:   []bool{true, false},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNewValues, err := ChunkZZZByBits(tt.args.values, tt.args.bits)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChunkZZZByBits() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNewValues, tt.wantNewValues) {
				t.Errorf("ChunkZZZByBits() = %v, want %v", gotNewValues, tt.wantNewValues)
			}
		})
	}
}

func TestReduceType(t *testing.T) {
	type args struct {
		values  []ZZZ
		f       func(acc, cur ZZZ) ZZZ
		initial ZZZ
	}
	tests := []struct {
		name         string
		args         args
		wantNewValue ZZZ
	}{
		{
			name: "ReduceZZZ",
			args: args{
				values: []ZZZ{1, 2, 3},
				f: func(acc, cur ZZZ) ZZZ {
					a := acc.(int)
					c := cur.(int)
					return a + c
				},
				initial: 0,
			},
			wantNewValue: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewValue := ReduceZZZ(tt.args.values, tt.args.f, tt.args.initial); !reflect.DeepEqual(gotNewValue, tt.wantNewValue) {
				t.Errorf("ReduceZZZ() = %v, want %v", gotNewValue, tt.wantNewValue)
			}
		})
	}
}

func TestReduceTypeSlice(t *testing.T) {
	type args struct {
		values  [][]ZZZ
		f       func(acc ZZZ, cur []ZZZ) ZZZ
		initial ZZZ
	}
	tests := []struct {
		name         string
		args         args
		wantNewValue ZZZ
	}{
		{
			name: "ReduceZZZSlice",
			args: args{
				values: [][]ZZZ{{1, 2, 3}, {4, 5, 6}},
				f: func(acc ZZZ, cur []ZZZ) ZZZ {
					a := acc.(int)
					for _, c := range cur {
						cInt := c.(int)
						a += cInt
					}
					return a
				},
				initial: 0,
			},
			wantNewValue: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewValue := ReduceZZZSlice(tt.args.values, tt.args.f, tt.args.initial); !reflect.DeepEqual(gotNewValue, tt.wantNewValue) {
				t.Errorf("ReduceZZZSlice() = %v, want %v", gotNewValue, tt.wantNewValue)
			}
		})
	}
}

func TestUnsetZZZ(t *testing.T) {
	type args struct {
		values []ZZZ
		i      int
	}
	tests := []struct {
		name    string
		args    args
		want    []ZZZ
		wantErr bool
	}{
		{
			name: "UnsetZZZ",
			args: args{
				values: []ZZZ{1, 2, 3, 4},
				i:      4,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "UnsetZZZ",
			args: args{
				values: []ZZZ{1, 2, 3, 4},
				i:      -1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "UnsetZZZ",
			args: args{
				values: []ZZZ{1, 2, 3, 4},
				i:      0,
			},
			want:    []ZZZ{2, 3, 4},
			wantErr: false,
		},
		{
			name: "UnsetZZZ",
			args: args{
				values: []ZZZ{1, 2, 3, 4},
				i:      1,
			},
			want:    []ZZZ{1, 3, 4},
			wantErr: false,
		},
		{
			name: "UnsetZZZ",
			args: args{
				values: []ZZZ{1, 2, 3, 4},
				i:      2,
			},
			want:    []ZZZ{1, 2, 4},
			wantErr: false,
		},
		{
			name: "UnsetZZZ",
			args: args{
				values: []ZZZ{1, 2, 3, 4},
				i:      3,
			},
			want:    []ZZZ{1, 2, 3},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnsetZZZ(tt.args.values, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnsetZZZ() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnsetZZZ() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAAACombination(t *testing.T) {
	type args struct {
		values []ZZZ
		r      int
	}
	tests := []struct {
		name             string
		args             args
		wantCombinations [][]ZZZ
		wantErr          bool
	}{
		{
			name: "ZZZCombination",
			args: args{
				values: []ZZZ{1, 2, 3},
				r:      2,
			},
			wantCombinations: [][]ZZZ{{2, 1}, {3, 1}, {3, 2}}, // FIXME ignore sequence
			wantErr:          false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCombinations, err := ZZZCombination(tt.args.values, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZZZCombination() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCombinations, tt.wantCombinations) {
				t.Errorf("ZZZCombination() = %v, want %v", gotCombinations, tt.wantCombinations)
			}
		})
	}
}

func TestZZZPermutation(t *testing.T) {
	type args struct {
		values []ZZZ
		r      int
	}
	tests := []struct {
		name             string
		args             args
		wantPermutations [][]ZZZ
	}{
		{
			name: "ZZZPermutation",
			args: args{
				values: []ZZZ{1, 2, 3},
				r:      3,
			},
			wantPermutations: [][]ZZZ{{3, 2, 1}, {2, 3, 1}, {3, 1, 2}, {1, 3, 2}, {2, 1, 3}, {1, 2, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPermutations, err := ZZZPermutation(tt.args.values, tt.args.r)
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(gotPermutations, tt.wantPermutations) {
				t.Errorf("ZZZPermutation() = %v, want %v", gotPermutations, tt.wantPermutations)
			}
		})
	}
}

func TestMapZZZSlice(t *testing.T) {
	type args struct {
		values [][]ZZZ
		f      func(v []ZZZ) []ZZZ
	}
	tests := []struct {
		name          string
		args          args
		wantNewValues [][]ZZZ
	}{
		{
			name: "MapZZZSlice",
			args: args{
				values: [][]ZZZ{{1, 2, 3}, {4, 5, 6}},
				f: func(v []ZZZ) []ZZZ {
					return append(v, 0)
				},
			},
			wantNewValues: [][]ZZZ{{1, 2, 3, 0}, {4, 5, 6, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewValues := MapZZZSlice(tt.args.values, tt.args.f); !reflect.DeepEqual(gotNewValues, tt.wantNewValues) {
				t.Errorf("MapZZZSlice() = %v, want %v", gotNewValues, tt.wantNewValues)
			}
		})
	}
}

func TestTernaryOPZZZ(t *testing.T) {
	type args struct {
		ok bool
		v1 ZZZ
		v2 ZZZ
	}
	tests := []struct {
		name string
		args args
		want ZZZ
	}{
		{
			name: "TernaryOPZZZ",
			args: args{
				ok: true,
				v1: 1,
				v2: 2,
			},
			want: 1,
		},
		{
			name: "TernaryOPZZZ",
			args: args{
				ok: false,
				v1: 1,
				v2: 2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TernaryOPZZZ(tt.args.ok, tt.args.v1, tt.args.v2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TernaryOPZZZ() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUnionFindZZZ(t *testing.T) {
	type args struct {
		values []ZZZ
	}
	tests := []struct {
		name string
		args args
		want *UnionFindZZZ
	}{
		{
			name: "NewUnionFindZZZ",
			args: args{
				values: []ZZZ{1, 2, 3},
			},
			want: &UnionFindZZZ{
				nodes: map[ZZZ]ZZZ{
					1: 1,
					2: 2,
					3: 3,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUnionFindZZZ(tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUnionFindZZZ() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnionFindZZZ_GetRoot(t *testing.T) {
	type fields struct {
		nodes map[ZZZ]ZZZ
	}
	type args struct {
		value ZZZ
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ZZZ
		want1  int
	}{
		{
			name: "UnionFindZZZ_GetRoot",
			fields: fields{
				nodes: map[ZZZ]ZZZ{
					1: 2,
					2: 3,
					3: 3,
				},
			},
			args: args{
				value: 1,
			},
			want:  3,
			want1: 2,
		},
		{
			name: "UnionFindZZZ_GetRoot",
			fields: fields{
				nodes: map[ZZZ]ZZZ{
					1: 2,
					2: 3,
					3: 3,
				},
			},
			args: args{
				value: 2,
			},
			want:  3,
			want1: 1,
		},
		{
			name: "UnionFindZZZ_GetRoot",
			fields: fields{
				nodes: map[ZZZ]ZZZ{
					1: 2,
					2: 3,
					3: 3,
				},
			},
			args: args{
				value: 3,
			},
			want:  3,
			want1: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UnionFindZZZ{
				nodes: tt.fields.nodes,
			}
			got, got1 := u.GetRoot(tt.args.value)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnionFindZZZ.GetRoot() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("UnionFindZZZ.GetRoot() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestUnionFindZZZ_Unite(t *testing.T) {
	type fields struct {
		nodes map[ZZZ]ZZZ
	}
	type args struct {
		v1 ZZZ
		v2 ZZZ
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ZZZ
		want1  bool
	}{
		{
			name: "UnionFindZZZ_Unite",
			fields: fields{
				nodes: map[ZZZ]ZZZ{
					1: 1,
					2: 2,
					3: 3,
				},
			},
			args: args{
				v1: 1,
				v2: 2,
			},
			want:  1,
			want1: true,
		},
		{
			name: "UnionFindZZZ_Unite",
			fields: fields{
				nodes: map[ZZZ]ZZZ{
					1: 2,
					2: 2,
					3: 3,
				},
			},
			args: args{
				v1: 1,
				v2: 2,
			},
			want:  2,
			want1: false,
		},
		{
			name: "UnionFindZZZ_Unite",
			fields: fields{
				nodes: map[ZZZ]ZZZ{
					1: 2,
					2: 2,
					3: 3,
				},
			},
			args: args{
				v1: 3,
				v2: 1,
			},
			want:  2,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UnionFindZZZ{
				nodes: tt.fields.nodes,
			}
			got, got1 := u.Unite(tt.args.v1, tt.args.v2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnionFindZZZ.Unite() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("UnionFindZZZ.Unite() got1 = %v, want %v", got1, tt.want1)
			}

			fromRoot, _ := u.GetRoot(tt.args.v1)
			toRoot, _ := u.GetRoot(tt.args.v2)
			if fromRoot != toRoot {
				t.Errorf("UnionFindZZZ.GetRoot() v1 root = %v, v2 root %v", fromRoot, toRoot)
			}
		})
	}
}

func TestUnionFindZZZ_IsSameGroup(t *testing.T) {
	type fields struct {
		nodes map[ZZZ]ZZZ
	}
	type args struct {
		v1 ZZZ
		v2 ZZZ
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "UnionFindZZZ_IsSameGroup",
			fields: fields{
				nodes: map[ZZZ]ZZZ{
					1: 2,
					2: 2,
					3: 3,
				},
			},
			args: args{
				v1: 1,
				v2: 2,
			},
			want: true,
		},
		{
			name: "UnionFindZZZ_IsSameGroup",
			fields: fields{
				nodes: map[ZZZ]ZZZ{
					1: 2,
					2: 2,
					3: 3,
				},
			},
			args: args{
				v1: 1,
				v2: 3,
			},
			want: false,
		},
		{
			name: "UnionFindZZZ_IsSameGroup",
			fields: fields{
				nodes: map[ZZZ]ZZZ{
					1: 2,
					2: 2,
					3: 3,
				},
			},
			args: args{
				v1: 2,
				v2: 3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UnionFindZZZ{
				nodes: tt.fields.nodes,
			}
			if got := u.IsSameGroup(tt.args.v1, tt.args.v2); got != tt.want {
				t.Errorf("UnionFindZZZ.IsSameGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCopyZZZSlice(t *testing.T) {
	type args struct {
		values [][]ZZZ
	}
	tests := []struct {
		name string
		args args
		want [][]ZZZ
	}{
		{
			name: "CopyZZZSlice",
			args: args{
				values: [][]ZZZ{{1, 2, 3}, {4, 5, 6}},
			},
			want: [][]ZZZ{{1, 2, 3}, {4, 5, 6}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CopyZZZSlice(tt.args.values)
			tt.args.values[0][0] = 999
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CopyZZZSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew2DimZZZSlice(t *testing.T) {
	type args struct {
		row int
		col int
	}
	tests := []struct {
		name string
		args args
		want [][]ZZZ
	}{
		{
			args: args{row: 2, col: 2},
			want: [][]ZZZ{{nil, nil}, {nil, nil}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New2DimZZZSlice(tt.args.row, tt.args.col); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New2DimZZZSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew2DimZZZSliceWithInitialValue(t *testing.T) {
	type args struct {
		row int
		col int
	}
	tests := []struct {
		name string
		args args
		want [][]ZZZ
	}{
		{
			args: args{row: 2, col: 2},
			want: [][]ZZZ{{1, 1}, {1, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New2DimZZZSliceWithInitialValue(tt.args.row, tt.args.col, 1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New2DimZZZSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewZZZSliceWithInitialValue(t *testing.T) {
	type args struct {
		length       int
		initialValue ZZZ
	}
	tests := []struct {
		name string
		args args
		want []ZZZ
	}{
		{
			args: args{
				length:       3,
				initialValue: 1,
			},
			want: []ZZZ{1, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewZZZSliceWithInitialValue(tt.args.length, tt.args.initialValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewZZZSliceWithInitialValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
