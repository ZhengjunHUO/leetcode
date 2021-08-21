package main

import (
	"fmt"
)

func splitMax(nums []int, cap int) int {
	rest, rslt := cap, 1
	
	for i := range nums {
		if nums[i] > rest {
			rslt++
			rest = cap - nums[i]
		}else{	
			rest -= nums[i]
		}
	}

	return rslt
}

// 把题目转换成二分搜寻max值，每次计算在给定的max下Array最多能分成几份
func splitArray(nums []int, m int) int {
	sum, max := 0, 0
	for i := range nums {
		if nums[i] > max {
			max = nums[i]
		}
		sum += nums[i]
	}

	l, r := max, sum + 1
	for l < r {
		mid := l + (r-l)/2
		if splitMax(nums, mid) <= m {
			r = mid
		}else{
			l = mid + 1
		}
	}

	return l
}

func main() {
	nums := [][]int{[]int{7,2,5,10,8}, []int{1,2,3,4,5}, []int{1,4,4}}
	ms := []int{2,2,3}

	for i := range nums {
		fmt.Println(splitArray(nums[i], ms[i]))
	}
}
