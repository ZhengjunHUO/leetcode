package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func kthSmallest(matrix [][]int, k int) int {
	rslt, pq := 0, godtype.NewPQ([][]int{}, []int{}, true)

	// 添加每一行的第一个元素到PQ, 由于行和列都是升序，所以如果矩阵超过了k行则超过部分不用考虑
	for i:=0; i<len(matrix) && i<k; i++ {
		pq.Push([]int{i, 0}, matrix[i][0])
	}
	
	for k > 0 && pq.Peek() != nil {
		curr := pq.Pop().([]int)	
		k--
		rslt = matrix[curr[0]][curr[1]]
		
		// 把Pop出的元素所在行的下一个元素(如果有)Push进PQ
		if curr[1] == len(matrix[curr[0]]) - 1 {
			continue
		}
		pq.Push([]int{curr[0], curr[1]+1}, matrix[curr[0]][curr[1]+1])
	}

	return rslt
}

func main() {
	fmt.Println(kthSmallest([][]int{[]int{1,5,9},[]int{10,11,13},[]int{12,13,15}}, 8))
	fmt.Println(kthSmallest([][]int{[]int{-5}}, 1))
}
