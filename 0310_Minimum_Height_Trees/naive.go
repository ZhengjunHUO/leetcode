package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func findMinHeightTrees(n int, edges [][]int) []int {
	min := 20000
	rslt := []int{}
	
	// 制作邻接表
	adj := make(map[int][]int)
	for i := range edges {
		adj[edges[i][0]] = append(adj[edges[i][0]], edges[i][1])
		adj[edges[i][1]] = append(adj[edges[i][1]], edges[i][0])
	}
	
	// 遍历所有节点作为根节点
	for i := 0; i < n; i++ {
		height := -1
		visited := make([]bool, n)

		q := godtype.NewQueue()
		q.Push(i)
		visited[i] = true

		for !q.IsEmpty() {
			// 遍历一圈邻居
			s := q.Size()
			for j:=0; j<s; j++ {
				curr := q.Pop().(int)
				if v, ok := adj[curr]; ok {
					for k := range v {
						if ! visited[v[k]] {
							q.Push(v[k])
							visited[v[k]] = true
						}
					}
				}
			}
			// 遍历完成后高度加1
			height++
		}

		// 如果高度正好等于最小值，则添加到列表
		if height == min {
			rslt = append(rslt, i)
		}

		// 如果小于最小值，则替代列表
		if height < min {
			min = height
			rslt = []int{i}
		}
	}

	return rslt
}

func main() {
	ns := []int{4, 6, 1, 2}
	edges := [][][]int{[][]int{[]int{1,0},[]int{1,2},[]int{1,3}}, [][]int{[]int{3,0},[]int{3,1},[]int{3,2},[]int{3,4},[]int{5,4}}, [][]int{}, [][]int{[]int{0,1}}}
	for i := range ns {
		fmt.Println(findMinHeightTrees(ns[i], edges[i]))
	}
}
