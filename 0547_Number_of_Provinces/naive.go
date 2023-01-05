package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/graph"
)

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	uf := graph.NewUF(n)

	for i := 0; i < n - 1 ; i++ {
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 1 {
				uf.Union(i, j)
			}
		}
	}

	return uf.Count()
}

func main() {
	ms := [][][]int{[][]int{[]int{1,1,0},[]int{1,1,0},[]int{0,0,1}}, [][]int{[]int{1,0,0},[]int{0,1,0},[]int{0,0,1}}}
	for i := range ms {
		fmt.Println(findCircleNum(ms[i]))
	}
}
