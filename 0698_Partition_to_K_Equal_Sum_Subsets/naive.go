package main

import (
	"fmt"
)

func recursive(nums []int, isUsed []bool, k, nextIdx, currVol, target int) bool {
	// 所有的subset都凑满了
	if k == 0 {
		return true
	}

	// 当前subset已满
	if currVol == target {
		return recursive(nums, isUsed, k-1, 0, 0, target)
	}

	for i := nextIdx; i<len(nums); i++ {
		// 跳过已使用过的元素和太大的元素
		if isUsed[i] || (currVol + nums[i] > target) {
			continue
		}

		isUsed[i] = true
		if recursive(nums, isUsed, k, i+1, currVol+nums[i], target) {
			return true
		}
		isUsed[i] = false
	}

	return false
}


func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}

	// 如果和除以k有余数则肯定不能分组
	if sum % k != 0 {
		return false
	}
	target := sum/k
	isUsed := make([]bool, len(nums))	

	return recursive(nums, isUsed, k, 0, 0, target)
}

func main() {
	nums := [][]int{[]int{4,3,2,3,5,2,1}, []int{1,2,3,4}}
	ks := []int{4, 3}

	for i := range nums {
		fmt.Println(canPartitionKSubsets(nums[i], ks[i]))
	}
}
