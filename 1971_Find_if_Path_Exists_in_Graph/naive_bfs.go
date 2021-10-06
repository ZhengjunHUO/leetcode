package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func validPath(n int, edges [][]int, start int, end int) bool {
	// 制作邻接表
	adj := make(map[int][]int)
	for i := range edges {
		adj[edges[i][0]] = append(adj[edges[i][0]], edges[i][1])
		adj[edges[i][1]] = append(adj[edges[i][1]], edges[i][0])
	}

	visited := make([]bool, n)
	q := godtype.NewQueue()
	q.Push(start)
	visited[start] = true 
	
	for !q.IsEmpty() {
		s := q.Size()
		for i:=0; i<s; i++ {
			curr := q.Pop().(int)
			if curr == end {
				return true
			}
			if v, ok := adj[curr]; ok {
				for j := range v {
					if !visited[v[j]] {
						q.Push(v[j])
						visited[v[j]] = true
					}
				}
			}
		}
	}

	return false
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
