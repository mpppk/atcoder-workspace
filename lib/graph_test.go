package lib

import (
	"reflect"
	"testing"
)

func TestNewGraph(t *testing.T) {
	type args struct {
		nodeNum  int
		edges    [][]int
		directed bool
	}
	tests := []struct {
		name    string
		args    args
		want    *Graph
		wantErr bool
	}{
		{
			name: "NewGraph",
			args: args{
				nodeNum:  2,
				edges:    [][]int{{0, 1}},
				directed: true,
			},
			want: &Graph{
				AdjacencyMatrix: [][]bool{{false, true}, {false, false}},
			},
			wantErr: false,
		},
		{
			name: "NewGraph",
			args: args{
				nodeNum:  2,
				edges:    [][]int{{0, 1}},
				directed: false,
			},
			want: &Graph{
				AdjacencyMatrix: [][]bool{{false, true}, {true, false}},
			},
			wantErr: false,
		},
		{
			name: "NewGraph",
			args: args{
				nodeNum:  0,
				edges:    nil,
				directed: false,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewGraph(tt.args.nodeNum, tt.args.edges, tt.args.directed)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewGraph() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_IsValidPath(t *testing.T) {
	type fields struct {
		AdjacencyMatrix [][]bool
	}
	type args struct {
		path []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Graph_IsValidPath",
			fields: fields{
				// 1 -> 2
				AdjacencyMatrix: [][]bool{{false, true}, {false, false}},
			},
			args: args{
				path: []int{0, 1},
			},
			want: true,
		},
		{
			name: "Graph_IsValidPath",
			fields: fields{
				// 1 -> 2
				AdjacencyMatrix: [][]bool{{false, true}, {false, false}},
			},
			args: args{
				path: []int{1, 0},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				AdjacencyMatrix: tt.fields.AdjacencyMatrix,
			}
			if got := g.IsValidPath(tt.args.path); got != tt.want {
				t.Errorf("Graph.IsValidPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
