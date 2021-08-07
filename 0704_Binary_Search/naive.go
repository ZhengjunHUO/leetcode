package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r - l)/2
		switch {
		case nums[mid] < target:
			l = mid + 1
		case nums[mid] > target:
			r = mid - 1
		case nums[mid] == target:	
			return mid
		}
	}

	return -1
}

func main() {
	nums := [][]int{[]int{-1,0,3,5,9,12}, []int{-1,0,3,5,9,12}}
	targets := []int{9, 2}
	for i := range nums {
		fmt.Println(search(nums[i], targets[i]))
	}
}
