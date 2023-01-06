package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

func minStoneSum(piles []int, k int) int {
	var ret int
	pq := datastruct.NewPQ(piles, piles, false)
	for k > 0 {
		curr := pq.Pop()
		reduced := curr - curr/2
		pq.Push(reduced, reduced)
		k--
	}

	for pq.Size() > 0 {
		ret += pq.Pop()
	}

	return ret
}

func main() {
	piles := [][]int{[]int{5,4,9}, []int{4,3,6,7}}
	k := []int{2, 3}
	for i := range piles {
		fmt.Println(minStoneSum(piles[i], k[i]))
	}
}
