package main

import (
	"fmt"
)

func backtrack(candidates []int, rslt *[][]int, curr []int, rest int, curridx int) {
	switch {
	case rest == 0:
		valid := make([]int, len(curr), cap(curr))
		copy(valid, curr)
		*rslt = append(*rslt, valid)
	case rest < 0:
		return
	default:
		for i:=curridx; i<len(candidates); i++ {
			curr = append(curr, candidates[i])
			backtrack(candidates, rslt, curr, rest - candidates[i], i)
			curr = curr[0:len(curr)-1]
		}
	}
}

func combinationSum(candidates []int, target int) [][]int {
	var rslt [][]int
	backtrack(candidates, &rslt, []int{}, target, 0)
	return rslt
}


func main() {
	fmt.Println(combinationSum([]int{2,3,6,7}, 7))
	fmt.Println(combinationSum([]int{2,3,5}, 8))
	fmt.Println(combinationSum([]int{2}, 1))
	fmt.Println(combinationSum([]int{1}, 1))
	fmt.Println(combinationSum([]int{1}, 2))
}
