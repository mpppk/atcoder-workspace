package generic

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

func TestMapType(t *testing.T) {
	type args struct {
		values []Type
		f      func(v Type) Type
	}
	tests := []struct {
		name          string
		args          args
		wantNewValues []Type
	}{
		{
			name: "can increment int",
			args: args{
				values: []Type{1, 2, 3},
				f: func(v Type) Type {
					intV, _ := v.(int)
					return intV + 1
				},
			},
			wantNewValues: []Type{2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewValues := MapType(tt.args.values, tt.args.f); !reflect.DeepEqual(gotNewValues, tt.wantNewValues) {
				t.Errorf("MapValue() = %v, want %v", gotNewValues, tt.wantNewValues)
			}
		})
	}
}

func TestMapTypeSlice(t *testing.T) {
	type args struct {
		values [][]Type
		f      func(v []Type) Type
	}
	tests := []struct {
		name          string
		args          args
		wantNewValues []Type
	}{
		{
			name: "mapTypeSlice",
			args: args{
				values: [][]Type{{1, 2, 3}, {4, 5, 6}},
				f: func(ts []Type) Type {
					t0, _ := ts[0].(int)
					t1, _ := ts[1].(int)
					t2, _ := ts[2].(int)
					return Type(t0 + t1 + t2)
				},
			},
			wantNewValues: []Type{6, 15},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewValues := MapTypeSlice(tt.args.values, tt.args.f); !reflect.DeepEqual(gotNewValues, tt.wantNewValues) {
				t.Errorf("MapTypeSlice() = %v, want %v", gotNewValues, tt.wantNewValues)
			}
		})
	}
}

func TestZipType(t *testing.T) {
	type args struct {
		valuesList [][]Type
	}
	tests := []struct {
		name              string
		args              args
		wantNewValuesList [][]Type
		wantErr           bool
	}{
		{
			name: "ZipType",
			args: args{
				valuesList: [][]Type{{1, 2, 3}, {4, 5, 6}},
			},
			wantNewValuesList: [][]Type{{1, 4}, {2, 5}, {3, 6}},
			wantErr:           false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNewValuesList, err := ZipType(tt.args.valuesList...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZipType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNewValuesList, tt.wantNewValuesList) {
				t.Errorf("ZipType() = %v, want %v", gotNewValuesList, tt.wantNewValuesList)
			}
		})
	}
}

func TestChunkByBits(t *testing.T) {
	type args struct {
		values []Type
		bits   []bool
	}
	tests := []struct {
		name          string
		args          args
		wantNewValues [][]Type
		wantErr       bool
	}{
		{
			name: "ChunkTypeByBits",
			args: args{
				values: []Type{1, 2, 3},
				bits:   []bool{true, false},
			},
			wantNewValues: [][]Type{{1}, {2, 3}},
			wantErr:       false,
		},
		{
			name: "ChunkByBits2",
			args: args{
				values: []Type{1, 2, 3},
				bits:   []bool{false, true},
			},
			wantNewValues: [][]Type{{1, 2}, {3}},
			wantErr:       false,
		},
		{
			name: "different length",
			args: args{
				values: []Type{1, 2, 3, 4},
				bits:   []bool{true, false},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNewValues, err := ChunkTypeByBits(tt.args.values, tt.args.bits)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChunkTypeByBits() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNewValues, tt.wantNewValues) {
				t.Errorf("ChunkTypeByBits() = %v, want %v", gotNewValues, tt.wantNewValues)
			}
		})
	}
}

func TestReduceType(t *testing.T) {
	type args struct {
		values  []Type
		f       func(acc, cur Type) Type
		initial Type
	}
	tests := []struct {
		name         string
		args         args
		wantNewValue Type
	}{
		{
			name: "ReduceType",
			args: args{
				values: []Type{1, 2, 3},
				f: func(acc, cur Type) Type {
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
			if gotNewValue := ReduceType(tt.args.values, tt.args.f, tt.args.initial); !reflect.DeepEqual(gotNewValue, tt.wantNewValue) {
				t.Errorf("ReduceType() = %v, want %v", gotNewValue, tt.wantNewValue)
			}
		})
	}
}

func TestReduceTypeSlice(t *testing.T) {
	type args struct {
		values  [][]Type
		f       func(acc Type, cur []Type) Type
		initial Type
	}
	tests := []struct {
		name         string
		args         args
		wantNewValue Type
	}{
		{
			name: "ReduceTypeSlice",
			args: args{
				values: [][]Type{{1, 2, 3}, {4, 5, 6}},
				f: func(acc Type, cur []Type) Type {
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
			if gotNewValue := ReduceTypeSlice(tt.args.values, tt.args.f, tt.args.initial); !reflect.DeepEqual(gotNewValue, tt.wantNewValue) {
				t.Errorf("ReduceTypeSlice() = %v, want %v", gotNewValue, tt.wantNewValue)
			}
		})
	}
}
