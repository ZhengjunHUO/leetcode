package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

func findRelativeRanks(score []int) []string {
	n := len(score)
	rslt := make([]string, n)

	pq := datastruct.NewPQ([]int{}, []int{}, true)
	// 分数作为priority值，内容为该元素的index
	for i,v := range score{
		pq.Push(i, v)
	}

	for pq.Size() != 0 {
		idx := pq.Pop()
		switch n {
			case 1:
				rslt[idx] = "Gold Medal"
			case 2:
				rslt[idx] = "Silver Medal"
			case 3:
				rslt[idx] = "Bronze Medal"
			default:
				rslt[idx] = fmt.Sprint(n)
		}
		n--
	}

	return rslt
}

func main() {
	scores := [][]int{[]int{5,4,3,2,1}, []int{10,3,8,9,4}}
	for _, v := range scores {
		fmt.Println(findRelativeRanks(v))
	}
}
