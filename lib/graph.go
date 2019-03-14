package lib

import "fmt"

type Graph struct {
	AdjacencyMatrix [][]bool
}

func NewGraph(nodeNum int, edges [][]int, directed bool) (*Graph, error) {
	if nodeNum < 1 {
		return nil, fmt.Errorf("invalid nodeNum: %d", nodeNum)
	}

	var aMatrix [][]bool
	for i := 0; i < nodeNum; i++ {
		line := make([]bool, nodeNum)
		aMatrix = append(aMatrix, line)
	}

	for _, edge := range edges {
		aMatrix[edge[0]][edge[1]] = true
		if !directed {
			aMatrix[edge[1]][edge[0]] = true
		}
	}

	return &Graph{AdjacencyMatrix: aMatrix}, nil
}

func (g *Graph) IsValidPath(path []int) bool {
	for i := 1; i < len(path); i++ {
		if !g.AdjacencyMatrix[path[i-1]][path[i]] {
			return false
		}
	}
	return true
}
