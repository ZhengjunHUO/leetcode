package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// 参考@heshan1234
func maximumInvitations(favorite []int) int {
	n := len(favorite)

	ingr, dp := make([]int, n), make([]int, n)
	visited := make([]bool, n)
	q := godtype.NewQueue()

	// calculate the "popularity" of each node
	// 计算每一个节点的入度，即有几个箭头指向它
	for i := range favorite {
		ingr[favorite[i]]++
	}

	// find out the node no one "appreciate"
	// 从入度为0的节点出发找无环链（加入一个队列并标记为已访问）
	for i := range ingr {
		if ingr[i] == 0 {
			q.Push(i)
			visited[i] = true
		}
	}

	// dp[i]: longest (non-loop) path ended at i
	for !q.IsEmpty() {
		src := q.Pop().(int)
		dst := favorite[src]
		dp[dst] = max(dp[dst], dp[src]+1)

		// 关键步骤，用来发现某些入度非0的节点其实是无环链的一部分
		ingr[dst]-=1
		if ingr[dst] == 0 {
			q.Push(dst)
			visited[dst] = true
		}
	}

	/* 
	   - if loop's size is 2, we can attach an acyclic chain at each end (at most 2),
	     and we can have serveral size==2 loops in result
	     尺寸等于2的环可以包容其他大小为2的环
	   - if loop's size > 2, we can only have at most one size>2 loop in result
	     尺寸大于2的环有排它性
	*/
	loopOfTwo, loop := 0, 0
	for i := range visited {
		if !visited[i] {
			size, curr := 0, i
			for !visited[curr] {
				visited[curr] = true
				size++
				curr = favorite[curr]
			}

			if size == 2 {
				loopOfTwo += 2 + dp[i] + dp[favorite[i]]
			}else{
				loop = max(loop, size)
			}
		}
	}

	return max(loopOfTwo, loop)
}

func main() {
	fs := [][]int{[]int{2,2,1,2}, []int{1,2,0}, []int{3,0,1,4,1}, []int{1,2,1,4,5,4,5}, []int{1,2,1,4,5,6,4,0,2}}
	for i := range fs {
		fmt.Println(maximumInvitations(fs[i]))
	}
}
