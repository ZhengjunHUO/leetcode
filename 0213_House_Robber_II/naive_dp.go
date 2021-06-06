package main

import (
	"fmt"
)

func rob(nums []int) int {
	n := len(nums) 

	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		}else{
			return nums[1]
		}
	}

	//把0和n-1拆开，分别求两个长度为n-1的子串的值
	rslt1 := robbis(nums, 0, n-2)
	rslt2 := robbis(nums, 1, n-1)

	if rslt1 > rslt2 {
		return rslt1
	}else{
		return rslt2
	}
}

func robbis(nums []int, lp int, rp int) int {
	preprev, prev := 0, nums[lp]
	for i:=lp+1; i<=rp; i++ {
		if temp:=preprev + nums[i]; temp > prev {
			preprev, prev = prev, temp
		}else{
			preprev, prev = prev, prev
		}
	}
	return prev
}

func main() {
	nums := []int{2,3,2}
	fmt.Println(rob(nums))

	nums = []int{1,2,3,1}
	fmt.Println(rob(nums))

	nums = []int{0}
	fmt.Println(rob(nums))

	nums = []int{3,2,1,4}
	fmt.Println(rob(nums))
}
