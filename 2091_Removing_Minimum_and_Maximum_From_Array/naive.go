package main

import (
	"fmt"
)

func min(minVal, minIdx *int, x, idx int) {
	if x < *minVal {
		*minVal = x
		*minIdx = idx
	}
}

func max(maxVal, maxIdx *int, x, idx int) {
	if x > *maxVal {
		*maxVal = x
		*maxIdx = idx
	}
}

func minInt(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func minimumDeletions(nums []int) int {
	minVal, minIdx, maxVal, maxIdx := 10001, -1, -10001, -1

	// find the min/max value and its index
	for i,v := range nums {
		min(&minVal, &minIdx, v, i)
		max(&maxVal, &maxIdx, v, i)
	}

	// reuse variables, make sure minIdx is less or equal than maxIdx
	if minIdx > maxIdx {
		minIdx, maxIdx = maxIdx, minIdx
	}

	// compare "remove from two ends" and "remove from left"
	ret := minInt((minIdx + 1) + (len(nums)-maxIdx), maxIdx + 1)
	// compare "remove from right"
	ret = minInt(ret, len(nums)-minIdx)

	return ret
}

func main() {
	nums := [][]int{[]int{2,10,7,5,4,1,8,6}, []int{0,-4,19,1,8,-2,-3,5}, []int{101}}
	for i := range nums {
		fmt.Println(minimumDeletions(nums[i]))
	}
}
