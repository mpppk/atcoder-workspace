package lib

import (
	"fmt"
	"math/big"
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
		n int
		r int
	}
	tests := []struct {
		name    string
		args    args
		want    int
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
				n: 2,
				r: 4,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Combination",
			args: args{
				n: 21,
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

func TestMemoizedCombination(t *testing.T) {
	type args struct {
		n int
		r int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "MemoizedCombination",
			args: args{
				n: 2,
				r: 4,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "MemoizedCombination",
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
				n: 21,
				r: 4,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MemoizedCombination(tt.args.n, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("MemoizedCombination() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MemoizedCombination() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBigFactorial(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			name: "BigFactorial",
			args: args{
				n: 4,
			},
			want: big.NewInt(24),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BigFactorial(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BigFactorial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBigCombination(t *testing.T) {
	type args struct {
		n int
		r int
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		{
			name: "MemoizedCombination",
			args: args{
				n: 2,
				r: 4,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "MemoizedCombination",
			args: args{
				n: 4,
				r: 2,
			},
			want:    big.NewInt(6),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BigCombination(tt.args.n, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("BigCombination() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BigCombination() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParallelBigCombination(t *testing.T) {
	type args struct {
		n int
		r int
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		{
			name: "ParallelBigCombination",
			args: args{
				n: 2,
				r: 4,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ParallelBigCombination",
			args: args{
				n: 4,
				r: 2,
			},
			want:    big.NewInt(6),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParallelBigCombination(tt.args.n, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParallelBigCombination() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParallelBigCombination() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemoizedBigCombination(t *testing.T) {
	type args struct {
		n int
		r int
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		{
			name: "MemoizedCombination",
			args: args{
				n: 2,
				r: 4,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "MemoizedCombination",
			args: args{
				n: 4,
				r: 2,
			},
			want:    big.NewInt(6),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MemoizedBigCombination(tt.args.n, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("MemoizedBigCombination() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MemoizedBigCombination() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCombination(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Combination(20, 19)
	}
}

func BenchmarkMemoizedCombination(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MemoizedCombination(20, 19)
	}
}

func BenchmarkBigCombination(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BigCombination(100000, 99999)
	}
}

func BenchmarkMemoizedBigCombination(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MemoizedBigCombination(100000, 99999)
	}
}

func BenchmarkParallelBigCombination(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println(ParallelBigCombination(100000, 99999))
	}
}
