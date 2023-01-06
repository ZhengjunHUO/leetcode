package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

// 和0787类似，使用dijkstra算法
func minCost(maxTime int, edges [][]int, passingFees []int) int {
	// 邻接表
	adj := make(map[int][][2]int)
	for i := range edges {
		adj[edges[i][0]] = append(adj[edges[i][0]], [2]int{edges[i][1], edges[i][2]})
		adj[edges[i][1]] = append(adj[edges[i][1]], [2]int{edges[i][0], edges[i][2]})
	}

	// 优先队列，value为[目标节点，剩余时间]，priority为累计花费，优先pop出花费最小的方案
	n := len(adj)
	pq := datastruct.NewPQ([][2]int{[2]int{0, maxTime}}, []int{passingFees[0]}, true)

	for pq.Size() > 0 {
		curr, prio := pq.PopWithPriority()
		currId, restTime, currFees := curr[0], curr[1], prio

		// 跳过超时的路线
		if restTime >= 0 {
			// 在规定时间内达到了目标，返回最小费用（优先队列确保）
			if currId == n - 1 {
				return currFees
			}

			if v, ok := adj[currId]; ok {
				for j := range v {
					// 到下一个城市，剩余时间减少，花费增加
					pq.Push([2]int{v[j][0], restTime-v[j][1]}, currFees + passingFees[v[j][0]])
				}
			}
		}
	}

	return -1
}

func main() {
	edges := [][]int{[]int{0,1,10},[]int{1,2,10},[]int{2,5,10},[]int{0,3,1},[]int{3,4,10},[]int{4,5,15}}
	pf := []int{5,1,2,20,20,3}
	maxT := []int{30,29,25}

	for i := range maxT {
		fmt.Println(minCost(maxT[i], edges, pf))
	}
}
