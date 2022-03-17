package main

import (
	"fmt"
	"sort"
)

func findOriginalArray(changed []int) []int {
	// 确保长度为偶数
	if len(changed) % 2 != 0 {
		return []int{}
	}

	occu := make(map[int]int)
	nums, ret := []int{}, []int{}

	// 构建元素出现次数的hash表
	for i := range changed {
		occu[changed[i]]++
	}

	// 构建一个元素的集合(不重复)并递增排序
	for k,_ := range occu {
		nums = append(nums, k)
	}
	sort.SliceStable(nums, func (i, j int) bool {
		return nums[i] < nums[j]
	})

	// 从集合中最小的元素开始分析
	for i := range nums {
		// 如果当前数字的两倍出现次数少于当前数字，则不成立
		if occu[nums[i]] > occu[2*nums[i]] {
			return []int{}
		}

		// 将当前数字加入结果（出现多少次加入几次）
		for j:=0; j<occu[nums[i]]; j++ {
			ret = append(ret, nums[i])
			// 同时减少当前数字的两倍出现次数
			occu[2*nums[i]]--
		}
	}

	return ret
}

func main() {
	changed := [][]int{[]int{1,3,4,2,6,8}, []int{6,3,0,1}, []int{1}, []int{0,0,0,0,1,2,2,4}}
	for i := range changed {
		fmt.Println(findOriginalArray(changed[i]))
	}
}
