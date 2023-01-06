package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

const MAXINT = int(^uint(0) >> 1)

func countPaths(n int, roads [][]int) int {
	// 邻接表
	adj := make(map[int][][2]int)
	for i := range roads {
		adj[roads[i][0]] = append(adj[roads[i][0]], [2]int{roads[i][1], roads[i][2]})
		adj[roads[i][1]] = append(adj[roads[i][1]], [2]int{roads[i][0], roads[i][2]})
	}

	// costTo表，每个节点对应一个二元组，在花费的基础上增加一个路径计数器
	costTo := make([][2]int, n)
	for i := range costTo {
		costTo[i] = [2]int{MAXINT, 0}
	}
	costTo[0] = [2]int{0,1}

	// 优先队列
	pq := datastruct.NewPQ([]int{0}, []int{0}, true)
	for pq.Size() > 0 {
                currNode, costToCurr := pq.PopWithPriority()

		if costToCurr > costTo[currNode][0] {
			continue
		}

		if v, ok := adj[currNode]; ok {
			for i := range v {
				costToNext := v[i][1] + costToCurr
				// 如果从起点经curr到next的花费和记录中相等，需要增加计数器的数字
				if costToNext == costTo[v[i][0]][0] {
					// 增加的量为当前节点在记录中的计数
					costTo[v[i][0]][1] += costTo[currNode][1]
					continue
				}

				// 找到一条捷径，更新记录中next的花费和计数器（当前节点在记录中的计数）
				if costToNext < costTo[v[i][0]][0] {
					costTo[v[i][0]][0], costTo[v[i][0]][1] = costToNext, costTo[currNode][1]
					pq.Push(v[i][0], costToNext)
				}
			}
		}
	}

	return costTo[n-1][1]
}

func main() {
	ns := []int{7,2}
	roads := [][][]int{[][]int{[]int{0,6,7},[]int{0,1,2},[]int{1,2,3},[]int{1,3,3},[]int{6,3,3},[]int{3,5,1},[]int{6,5,1},[]int{2,5,1},[]int{0,4,5},[]int{4,6,2}}, [][]int{[]int{1,0,10}}}
	for i := range roads {
		fmt.Println(countPaths(ns[i], roads[i]))
	}
}
