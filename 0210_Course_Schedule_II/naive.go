package main

import (
	"fmt"
)

// 基于0207的解法基础上，在后序遍历添加代码
func findOrder(numCourses int, prerequisites [][]int) []int {
	rslt := []int{}

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
		// 后序遍历，将当前节点加入结果
		rslt = append(rslt, curr)
	}

	// 可能有好几棵树，所以每个节点都要尝试bfs搜索
	for i := 0; i < numCourses; i++ {
		bfs(i)
	}

	// 如有循环依赖返回空slice
	if hasLoop {
		return []int{}
	}

	// 反转后序遍历所得的rslt并返回
	l, r := 0, len(rslt) - 1
	for l < r {
		rslt[l], rslt[r] = rslt[r], rslt[l]
		l++
		r--
	}
	return rslt
}

func main() {
	ns := []int{2,4,1}
	prerequis := [][][]int{[][]int{[]int{1,0}}, [][]int{[]int{1,0},[]int{2,0},[]int{3,1},[]int{3,2}}, [][]int{}}
	for i := range prerequis {
		fmt.Println(findOrder(ns[i], prerequis[i]))
	}
}
