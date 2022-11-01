package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

const MAXINT = int(^uint(0) >> 1)

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func networkBecomesIdle(edges [][]int, patience []int) int {
	// 制作邻接表
	n := len(patience)
	neigh := make([][]int, n)
	for i := range neigh {
		neigh[i] = []int{}
	}

	for i := range edges {
		neigh[edges[i][0]] = append(neigh[edges[i][0]], edges[i][1])
		neigh[edges[i][1]] = append(neigh[edges[i][1]], edges[i][0])
	}

	// 避免重复访问
	visited := make([]bool, n)
	// 记录每个节点到服务器的最短距离
	distance := make([]int, n)
	for i:=1; i<n; i++ {
		distance[i] = MAXINT
	}

	// BFS计算每个节点到服务器的最短距离
	q := datastruct.NewQueue([]int{})
	q.Push(0)
	visited[0] = true

	for !q.IsEmpty() {
		curr := q.Pop()
		for _, v := range neigh[curr] {
			if !visited[v] {
				distance[v] = min(distance[v], distance[curr]+1)
				visited[v] = true
				q.Push(v)
			}
		}
	}

	ret := -1
	for i:=1; i<n; i++ {
		// 每个包往返需要的时间
		rtt := 2 * distance[i]
		/*
		  rtt-1: 可以重发包的最后时间点(因为rtt时收到第一个包的回复，便不再重发)
		  (rtt-1)/patience[i]: rtt期间发送的包的数量
		  ((rtt-1)/patience[i])* patience[i]: 最后一个包发出的时间点，再加上rtt可以得到结果
		*/
		ret = max(ret, ((rtt-1)/patience[i])* patience[i] + rtt)
	}

	// 根据题目要求需要加1
	return ret+1
}

func main() {
	edges := [][][]int{[][]int{[]int{0,1},[]int{1,2}}, [][]int{[]int{0,1},[]int{0,2},[]int{1,2}}}
	patience := [][]int{[]int{0,2,1}, []int{0,10,10}}

	for i := range edges {
		fmt.Println(networkBecomesIdle(edges[i], patience[i]))
	}
}
