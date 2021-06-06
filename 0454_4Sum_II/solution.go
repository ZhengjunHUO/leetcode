package main

import (
	"fmt"
)

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	dict := make(map[int]int)
	rslt := 0
	
	for _,i := range nums1 {
		for _,j := range nums2 {
			dict[i+j] += 1	
		}
	}

	fmt.Println(dict)

	for _,x := range nums3 {
		for _,y := range nums4 {
			rslt += dict[-x-y]
		}
	}

	return rslt
    
}

func main() {
	nums1 := []int{1,2}
	nums2 := []int{-2,-1}
	nums3 := []int{-1,2}
	nums4 := []int{0,2}

	fmt.Println(fourSumCount(nums1, nums2, nums3, nums4))
}
