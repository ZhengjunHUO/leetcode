package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func isBipartite(graph [][]int) bool {
	n := len(graph)
	uf := godtype.NewUF(n)

	for i := range graph {
		for j := range graph[i] {
			if uf.IsLinked(i, graph[i][j]) {
				return false
			}

			uf.Union(graph[i][0], graph[i][j])
		}
	}

	return true
}

func main() {
	gs := [][][]int{[][]int{[]int{1,2,3},[]int{0,2},[]int{0,1,3},[]int{0,2}}, [][]int{[]int{1,3},[]int{0,2},[]int{1,3},[]int{0,2}}}
	for i := range gs {
		fmt.Println(isBipartite(gs[i]))
	}
}
