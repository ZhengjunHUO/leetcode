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

// Learned from @heshan1234
func maximumInvitations(favorite []int) int {
	n := len(favorite)

	ingr, dp := make([]int, n), make([]int, n)
	visited := make([]bool, n)
	q := godtype.NewQueue()

	// calculate the "popularity" of each node
	for i := range favorite {
		ingr[favorite[i]]++
	}

	// find out the node no one "appreciate"
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

		ingr[dst]-=1
		if ingr[dst] == 0 {
			q.Push(dst)
			visited[dst] = true
		}
	}

	/* 
	   - if loop's size is 2, we can attach an acyclic chain at each end (at most 2),
	     and we can have serveral size==2 loops in result
	   - if loop's size > 2, we can only have at most one size>2 loop in result
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
