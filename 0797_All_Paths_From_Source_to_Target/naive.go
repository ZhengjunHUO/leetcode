package main

import (
	"fmt"
)

func allPathsSourceTarget(graph [][]int) [][]int {
	curr, rslt := []int{}, [][]int{}
	recursive(graph, curr, 0, &rslt)
	return rslt
}

func recursive(graph [][]int, curr []int, node int, rslt *[][]int) {
	curr = append(curr, node)

	if node == len(graph) - 1 {
		*rslt = append(*rslt, curr)
		curr = curr[:len(curr)-1]
		return 
	}

	for _, v := range graph[node] {
		recursive(graph, curr, v, rslt)
	}

	curr = curr[:len(curr)-1]
}


func main() {
	graphs := [][][]int{[][]int{[]int{1,2},[]int{3},[]int{3},[]int{}}, [][]int{[]int{4,3,1},[]int{3,2,4},[]int{3},[]int{4},[]int{}},[][]int{[]int{1},[]int{}}, [][]int{[]int{1,2,3},[]int{2},[]int{3},[]int{}}, [][]int{[]int{1,3},[]int{2},[]int{3},[]int{}}}
	for i := range graphs {
		fmt.Println(allPathsSourceTarget(graphs[i]))
	}
}
