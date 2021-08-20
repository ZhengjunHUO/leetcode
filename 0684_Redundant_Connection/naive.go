package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	uf := godtype.NewUF(n+1)

	for i := range edges {
		if uf.IsLinked(edges[i][0], edges[i][1]) {
			return edges[i]
		}
		
		uf.Union(edges[i][0], edges[i][1])	
	}

	return []int{}
}

func main() {
	edges := [][][]int{[][]int{[]int{1,2},[]int{1,3},[]int{2,3}}, [][]int{[]int{1,2},[]int{2,3},[]int{3,4},[]int{1,4},[]int{1,5}}}
	for i := range edges {
		fmt.Println(findRedundantConnection(edges[i]))
	}
}
