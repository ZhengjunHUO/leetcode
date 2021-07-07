package main

import (
	"fmt"
)

func backtrack(nums []int, rslt *[][]int, curr []int, dict map[int]int) {
	if len(curr) == len(nums) {
		valid := make([]int, len(nums), cap(nums))
		copy(valid, curr)
		*rslt = append(*rslt, valid)
		return
	}

	prev := -1
	for i:=0; i<len(nums); i++ {
		// 不同于0046_Permutations，因为有重复元素所以字典使用index而不是值作为key
		if _, ok := dict[i]; ok {
			continue
		}

		// 去重，同一层上如果遇到当前元素的值等于上一轮元素的值的情况直接跳过
		if prev >=0 && nums[prev] == nums[i] {
			continue
		}

		dict[i] = 1
		curr = append(curr, nums[i])
		prev = i

		backtrack(nums, rslt, curr, dict)

		delete(dict, i)
		curr = curr[0:len(curr)-1]
	}
}

func permuteUnique(nums []int) [][]int {
	var rslt [][]int
	dict := make(map[int]int)
	backtrack(nums, &rslt, []int{}, dict)
	return rslt
}

func main() {
	fmt.Println(permuteUnique([]int{1,1,2}))
	fmt.Println(permuteUnique([]int{1,2,3}))
}
