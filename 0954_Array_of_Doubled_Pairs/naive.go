package main

import (
	"fmt"
	"sort"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

// 借鉴2007的解法
func canReorderDoubled(arr []int) bool {
	// 确保长度为偶数
	if len(arr) % 2 != 0 {
		return false
	}

	occu := make(map[int]int)
	nums := []int{}

	// 构建元素出现次数的hash表
	for i := range arr {
		occu[arr[i]]++
	}

	// 构建一个元素的集合(不重复)并递增排序
	for k,_ := range occu {
		nums = append(nums, k)
	}
	sort.SliceStable(nums, func (i, j int) bool {
		return abs(nums[i]) < abs(nums[j])
	})

	// 从集合中最小的元素开始分析
	for i := range nums {
		// 如果当前数字的两倍出现次数少于当前数字，则不成立
		if occu[nums[i]] > occu[2*nums[i]] {
			return false
		}

		for j:=0; j<occu[nums[i]]; j++ {
			// 减少当前数字的两倍出现次数
			occu[2*nums[i]]--
		}
	}

	return true
}

func main() {
	arrs := [][]int{[]int{3,1,3,6}, []int{2,1,2,6}, []int{4,-2,2,-4}}
	for i := range arrs {
		fmt.Println(canReorderDoubled(arrs[i]))
	}
}
