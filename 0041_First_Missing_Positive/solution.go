package main

import (
	"fmt"
)

// 参考@asones的方法
func firstMissingPositive(nums []int) int {
	nums = append(nums, 0)
	n := len(nums)

	// 只保留1到len(原nums)之间的值
	for i, v := range nums {
		if v >= n || v < 0 {
			nums[i] = 0
		}
	}	

	/*
	让非零值"归位"，如有2就到nums[2]上做一个标记(加n)
	如果1缺失，那么nums[1]的值不会被标记到，即一定小于n
  	在下一轮遍历中会被if判断找出	 
	*/ 
	for i := range nums {
		nums[nums[i]%n]+=n
	}

	for i := range nums {
		if nums[i]/n == 0 {
			return i
		}
	}

	return n
}

func main() {
	nums := [][]int{[]int{1,2,0}, []int{3,4,-1,1}, []int{7,8,9,11,12}}

	for _,v := range nums {
		fmt.Println(firstMissingPositive(v))
	}
}
