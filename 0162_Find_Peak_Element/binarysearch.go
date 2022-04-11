package main

import (
	"fmt"
)

/*
  learn from @cosmin79's solution
  idea is to keep [l, r] greater than nums[l-1] and nums[r+1] ([l, r] is the peak)
  nums[-1] = nums[n] = -∞ so initially the [0, n-1] is a peak
  then do the binary search, compare the nums[m] and nums[m+1] and pick the "bigger" side's half 
  until len([l, r]) is 2, pick up the bigger one
*/
func findPeakElement(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}

	l, r := 0, n-1
	for r - l > 1 {
		m := l + (r-l)/2
		if (nums[m] < nums[m+1]) {
			l = m + 1
		}else{
			r = m
		}
	}

	if l == n-1 || nums[l] > nums[l+1] {
		return l
	}

	return r
}

func main() {
	nums := [][]int{[]int{1,2,3,1}, []int{1,2,1,3,5,6,4}}
	for i := range nums {
		fmt.Println(findPeakElement(nums[i]))
	}
}
