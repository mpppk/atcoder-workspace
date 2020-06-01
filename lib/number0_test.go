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

func TestCombination(t *testing.T) {
	type args struct {
		n int64
		r int64
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "Combination",
			args: args{
				n: 4,
				r: 2,
			},
			want:    6,
			wantErr: false,
		},
		{
			name: "Combination",
			args: args{
				n: 10,
				r: 5,
			},
			want:    252,
			wantErr: false,
		},
		{
			name: "Combination",
			args: args{
				n: 100,
				r: 2,
			},
			want:    4950,
			wantErr: false,
		},
		{
			name: "Combination",
			args: args{
				n: 2,
				r: 4,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Combination(tt.args.n, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Combination() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Combination() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCombination(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Combination(20, 19)
	}
}

func TestAdjacencyList(t *testing.T) {
	type args struct {
		x      []int
		y      []int
		length int
	}
	tests := []struct {
		name    string
		args    args
		want    [][]int
		wantErr bool
	}{
		{
			args: args{
				x:      []int{1, 2, 3},
				y:      []int{4, 4, 5},
				length: 5,
			},
			want: [][]int{
				{4}, {4}, {5},
				{1, 2}, {3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AdjacencyList(tt.args.x, tt.args.y, tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("AdjacencyList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AdjacencyList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDirectedAdjacencyList(t *testing.T) {
	type args struct {
		x      []int
		y      []int
		length int
	}
	tests := []struct {
		name    string
		args    args
		want    [][]int
		wantErr bool
	}{
		{
			args: args{
				x:      []int{1, 2, 3},
				y:      []int{4, 4, 5},
				length: 5,
			},
			want: [][]int{
				{4}, {4}, {5},
				nil, nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DirectedAdjacencyList(tt.args.x, tt.args.y, tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("DirectedAdjacencyList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DirectedAdjacencyList() got = %v, want %v", got, tt.want)
			}
		})
	}
}
