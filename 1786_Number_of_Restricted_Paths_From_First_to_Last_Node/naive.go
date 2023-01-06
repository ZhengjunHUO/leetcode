package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

const MAXINT = int(^uint(0) >> 1)

func countRestrictedPaths(n int, edges [][]int) int {
	// 邻接表，双向图
	adj := make(map[int][][2]int)
	for i := range edges {
		adj[edges[i][0]] = append(adj[edges[i][0]], [2]int{edges[i][1], edges[i][2]})
		adj[edges[i][1]] = append(adj[edges[i][1]], [2]int{edges[i][0], edges[i][2]})
	}

	// costTo表
	costTo := make([]int, n+1)
	for i := range costTo {
		costTo[i] = MAXINT
	}
	// 以n为起点
	costTo[n] = 0

	// 优先队列
	pq := datastruct.NewPQ([]int{n}, []int{0}, true)
	for pq.Size() > 0 {
                currNode, costToCurr := pq.PopWithPriority()

		if costToCurr > costTo[currNode] {
			continue
		}

		if v, ok := adj[currNode]; ok {
			for i := range v {
				costToNext := costToCurr + v[i][1]
				if costToNext < costTo[v[i][0]] {
					costTo[v[i][0]] = costToNext
					pq.Push(v[i][0], costToNext)
				}
			}
		}
	}

	// 存储所有路径
	rslt := []string{}
	var backtrack func(int, string)

	backtrack = func(currIdx int, currPath string) {
		if currIdx == n {
			rslt = append(rslt, currPath)
			return
		}

		for i := range adj[currIdx] {
			nextId := adj[currIdx][i][0]
			if costTo[currIdx] > costTo[nextId] {
				backtrack(nextId, currPath + fmt.Sprint(nextId))
			}
		}
	}

	backtrack(1, "1")

	return len(rslt)
}

func main() {
	ns := []int{5,7}
	edges := [][][]int{[][]int{[]int{1,2,3},[]int{1,3,3},[]int{2,3,1},[]int{1,4,2},[]int{5,2,2},[]int{3,5,1},[]int{5,4,10}}, [][]int{[]int{1,3,1},[]int{4,1,2},[]int{7,3,4},[]int{2,5,3},[]int{5,6,1},[]int{6,7,2},[]int{7,5,3},[]int{2,6,4}}}
	for i := range edges {
		fmt.Println(countRestrictedPaths(ns[i],edges[i]))
	}
}
