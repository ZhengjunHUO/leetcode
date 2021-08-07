package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		m := l + (r - l)/2
		if nums[m] >= target {
			r = m - 1
		}else{
			l = m + 1
		}
	}

	return l
}

func main() {
	test := []int{5, 2, 7, 0}
	for _,v := range test {
		fmt.Println(searchInsert([]int{1,3,5,6}, v))
	}

	fmt.Println(searchInsert([]int{1}, 0))

	test = []int{8, 2}
	for _,v := range test {
		fmt.Println(searchInsert([]int{1,3,4,7,9}, v))
	}
}
