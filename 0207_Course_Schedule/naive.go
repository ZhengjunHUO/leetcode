package main

import (
	"fmt"
)

func canFinish(numCourses int, prerequisites [][]int) bool {
	// (1) 邻接表，构建DAG
	adj := make(map[int][]int)
	for i := range prerequisites {
		adj[prerequisites[i][1]] = append(adj[prerequisites[i][1]], prerequisites[i][0])
	}

	// (2) 遍历，探测是否有环，即循环依赖
	hasLoop := false
	visited, inCurrPath := make([]bool, numCourses), make([]bool, numCourses)

	// 递归调用
	var bfs func(int)
	bfs = func(curr int) {
		if inCurrPath[curr] {
			hasLoop = true
		}

		if visited[curr] || hasLoop {
			return
		}

		inCurrPath[curr], visited[curr] = true, true
		if v, ok := adj[curr]; ok {
			for i := range v {
				bfs(v[i])
			}
		}
		inCurrPath[curr] = false
	}

	// 可能有好几棵树，所以每个节点都要尝试bfs搜索
	for i := 0; i < numCourses; i++ {
		bfs(i)
	}

	// 如无循环依赖则可顺利完成课程
	return !hasLoop
}

func main() {
	ns := []int{2,2,4}
	prerequis := [][][]int{[][]int{[]int{1,0}}, [][]int{[]int{1,0},[]int{0,1}}, [][]int{[]int{1,0},[]int{2,0},[]int{3,1},[]int{3,2}}}
	for i := range prerequis {
		fmt.Println(canFinish(ns[i], prerequis[i]))
	}
}
