package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	
	for l <= r {
		m := l + ( r - l )/2
		switch {
		case target == nums[m]:
			return m
		case target == nums[l]:
			return l
		case target == nums[r]:
			return r
		case target < nums[m] && target < nums[l]:
			l = m + 1
		case target < nums[m] && target > nums[l]:
			r = m - 1
		case target > nums[m] && target > nums[r]:
			return -1
		case target > nums[m] && target < nums[r]:
			l = m + 1
		}
	}

	return -1
}

func main() {
	fmt.Println(search([]int{4,5,6,7,0,1,2}, 0))
	fmt.Println(search([]int{4,5,6,7,0,1,2}, 3))
	fmt.Println(search([]int{1}, 0))
	fmt.Println(search([]int{13, 15, 22, 30, 47, 54, 1, 3, 7}, 15))
}
