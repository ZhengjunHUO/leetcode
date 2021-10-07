package main

import (
	"fmt"
)

func findCenter(edges [][]int) int {
	if edges[0][0] == edges[1][0] || edges[0][0] == edges[1][1] {
		return edges[0][0]
	}

	return edges[0][1]
}

func main() {
	edges := [][][]int{[][]int{[]int{1,2},[]int{2,3},[]int{4,2}}, [][]int{[]int{1,2},[]int{5,1},[]int{1,3},[]int{1,4}}}
	for i := range edges {
		fmt.Println(findCenter(edges[i]))
	}
}
