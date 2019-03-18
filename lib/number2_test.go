package lib

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

func TestMaxAAAByBBBSlice(t *testing.T) {
	type args struct {
		values [][]BBB
		f      func(vs []BBB) AAA
	}
	tests := []struct {
		name    string
		args    args
		wantMax AAA
		wantErr bool
	}{
		{
			name: "ZZZTypeSliceToTypeSlice",
			args: args{
				values: [][]BBB{{1, 2, 3}, {4, 5, 6}},
				f: func(vs []BBB) AAA {
					sum := 0
					for _, v := range vs {
						sum += int(v)
					}
					return AAA(sum)
				},
			},
			wantMax: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMax, err := MaxAAAByBBBSlice(tt.args.values, tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaxAAAByBBBSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotMax, tt.wantMax) {
				t.Errorf("MaxAAAByBBBSlice() = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}

func TestSumAAAByBBB(t *testing.T) {
	type args struct {
		values []AAA
		f      func(v AAA) BBB
	}
	tests := []struct {
		name string
		args args
		want BBB
	}{
		{
			name: "SumAAAByBBB",
			args: args{
				values: []AAA{1, 2, 3},
				f: func(a AAA) BBB {
					return BBB(int(a) * int(a))
				},
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumAAAByBBB(tt.args.values, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumAAAByBBB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBBB3ToAAACache_Has(t *testing.T) {
	c := BBB3ToAAACache{
		0: map[BBB]map[BBB]AAA{
			0: {
				0: 1,
			},
		},
	}
	type args struct {
		k1 BBB
		k2 BBB
		k3 BBB
	}
	tests := []struct {
		name string
		c    BBB3ToAAACache
		args args
		want bool
	}{
		{
			name: "BBB3ToAAACache_Has",
			c:    c,
			args: args{
				k1: 0,
				k2: 0,
				k3: 0,
			},
			want: true,
		},
		{
			name: "BBB3ToAAACache_Has",
			c:    c,
			args: args{
				k1: 0,
				k2: 0,
				k3: 1,
			},
			want: false,
		},
		{
			name: "BBB3ToAAACache_Has",
			c:    c,
			args: args{
				k1: 0,
				k2: 1,
				k3: 0,
			},
			want: false,
		},
		{
			name: "BBB3ToAAACache_Has",
			c:    c,
			args: args{
				k1: 1,
				k2: 0,
				k3: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Has(tt.args.k1, tt.args.k2, tt.args.k3); got != tt.want {
				t.Errorf("BBB3ToAAACache.Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBBB3ToAAACache_Get(t *testing.T) {
	c := BBB3ToAAACache{
		0: map[BBB]map[BBB]AAA{
			0: {
				0: 1,
			},
		},
	}
	type args struct {
		k1 BBB
		k2 BBB
		k3 BBB
	}
	tests := []struct {
		name  string
		c     BBB3ToAAACache
		args  args
		want  AAA
		want1 bool
	}{
		{
			name: "BBB3ToAAACache_Get",
			c:    c,
			args: args{
				k1: 0,
				k2: 0,
				k3: 0,
			},
			want:  1,
			want1: true,
		},
		{
			name: "BBB3ToAAACache_Get",
			c:    c,
			args: args{
				k1: 0,
				k2: 0,
				k3: 1,
			},
			want:  0,
			want1: false,
		},
		{
			name: "BBB3ToAAACache_Get",
			c:    c,
			args: args{
				k1: 0,
				k2: 1,
				k3: 0,
			},
			want:  0,
			want1: false,
		},
		{
			name: "BBB3ToAAACache_Get",
			c:    c,
			args: args{
				k1: 1,
				k2: 0,
				k3: 0,
			},
			want:  0,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.c.Get(tt.args.k1, tt.args.k2, tt.args.k3)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BBB3ToAAACache.Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("BBB3ToAAACache.Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBBB3ToAAACache_Set(t *testing.T) {
	type args struct {
		k1 BBB
		k2 BBB
		k3 BBB
		v  AAA
	}
	tests := []struct {
		name string
		c    BBB3ToAAACache
		args args
		want AAA
	}{
		{
			name: "BBB3ToAAACache_Set",
			c:    BBB3ToAAACache{},
			args: args{
				k1: 0,
				k2: 0,
				k3: 0,
				v:  1,
			},
			want: 1,
		},
		{
			name: "BBB3ToAAACache_Set",
			c:    BBB3ToAAACache{},
			args: args{
				k1: 0,
				k2: 1,
				k3: 0,
				v:  1,
			},
			want: 1,
		},
		{
			name: "BBB3ToAAACache_Set",
			c:    BBB3ToAAACache{},
			args: args{
				k1: 0,
				k2: 1,
				k3: 0,
				v:  1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Set(tt.args.k1, tt.args.k2, tt.args.k3, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BBB3ToAAACache.Set() = %v, want %v", got, tt.want)
			}
			if got := tt.c[tt.args.k1][tt.args.k2][tt.args.k3]; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BBB3ToAAACache.Set() = %v, want %v", got, tt.want)
			}
		})
	}
}
