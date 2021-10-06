package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func validPath(n int, edges [][]int, start int, end int) bool {
	uf := godtype.NewUF(n)
	for i := range edges {
		uf.Union(edges[i][0], edges[i][1])
	}

	return uf.IsLinked(start, end)
}

func main() {
	ns := []int{3,6}
	edges := [][][]int{[][]int{[]int{0,1},[]int{1,2},[]int{2,0}}, [][]int{[]int{0,1},[]int{0,2},[]int{3,5},[]int{5,4},[]int{4,3}}}
	starts := []int{0,0}
	ends := []int{2,5}
	
	for i := range ns {
		fmt.Println(validPath(ns[i], edges[i], starts[i], ends[i]))
	}
}
