package main

import (
	"fmt"
)

func maximalNetworkRank(n int, roads [][]int) int {
	egress := make([]int, n)
	conn := make(map[[2]int]int)

	for i := range roads {
		egress[roads[i][0]] += 1
		egress[roads[i][1]] += 1
		conn[[2]int{roads[i][0], roads[i][1]}], conn[[2]int{roads[i][1], roads[i][0]}] = 1, 1
	}

	max := 0
	// 暴力破解
	for i := range egress {
		for j:=i+1; j<n; j++ {
			// 需要减去i,j之间（如果有）共有的边
			if rank := egress[i] + egress[j] - conn[[2]int{i, j}]; rank > max {
				max = rank
			}
		}
	}

	return max
}

func main() {
	roads := [][][]int{[][]int{[]int{0,1},[]int{0,3},[]int{1,2},[]int{1,3}}, [][]int{[]int{0,1},[]int{0,3},[]int{1,2},[]int{1,3},[]int{2,3},[]int{2,4}}, [][]int{[]int{0,1},[]int{1,2},[]int{2,3},[]int{2,4},[]int{5,6},[]int{5,7}}}
	ns := []int{4,5,8}

	for i := range ns {
		fmt.Println(maximalNetworkRank(ns[i], roads[i]))
	}
}
