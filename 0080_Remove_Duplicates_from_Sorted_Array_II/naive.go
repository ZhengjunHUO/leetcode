package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	k := 0

	for _, v := range nums {
		if k < 2 || nums[k-2] < v {
			nums[k] = v
			k += 1
		}
	}    

	return k
}

func main() {
	nums1 := []int{1,1,1,2,2,3}
	fmt.Println(nums1[:removeDuplicates(nums1)])

	nums2 := []int{0,0,1,1,1,1,2,3,3}
	fmt.Println(nums2[:removeDuplicates(nums2)])
}
