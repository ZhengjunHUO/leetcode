package main

import (
	"fmt"
	"sort"
)

func backtrack(candidates []int, dupl map[string]int, rslt *[][]int, curr []int, rest int, curridx int) {
	switch {
	case rest < 0:
		return
	case rest == 0:
		key := fmt.Sprint(curr)
		if _, ok := dupl[key]; !ok {
			valid := make([]int, len(curr), cap(curr))
			copy(valid, curr)
			*rslt = append(*rslt, valid)
			dupl[key] = 1
		}
	default:
		for i:=curridx; i<len(candidates); i++ {
			curr = append(curr, candidates[i])
			backtrack(candidates, dupl, rslt, curr, rest - candidates[i], i+1)
			curr = curr[0:len(curr)-1]
		}
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var rslt [][]int
	dupl := make(map[string]int)
	backtrack(candidates, dupl, &rslt, []int{}, target, 0)
	return rslt
}

func main() {
	fmt.Println(combinationSum2([]int{10,1,2,7,6,1,5}, 8))
	fmt.Println(combinationSum2([]int{2,5,2,1,2}, 5))
}
