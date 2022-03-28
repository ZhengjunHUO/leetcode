package main

import (
	"fmt"
)

func checkReplica(nums []int, dict map[int]int) {
	local := make(map[int]struct{})
	for i := range nums {
		if _, ok := local[nums[i]]; !ok {
			local[nums[i]] = struct{}{}
			if _, exist := dict[nums[i]]; exist {
				dict[nums[i]] += 1
			}else{
				dict[nums[i]] = 1
			}
		}
	}
}

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	ret := []int{}
	dict := make(map[int]int)

	for i := range nums1 {
		dict[nums1[i]] = 1
	}

	checkReplica(nums2, dict)
	checkReplica(nums3, dict)

	for k, v := range dict {
		if v > 1 {
			ret = append(ret, k)
		}
	}

	return ret
}

func main() {
	nums1 := [][]int{[]int{1,1,3,2}, []int{3,1}, []int{1,2,2}}
	nums2 := [][]int{[]int{2,3}, []int{2,3}, []int{4,3,3}}
	nums3 := [][]int{[]int{3}, []int{1,2}, []int{5}}
	for i := range nums1 {
		fmt.Println(twoOutOfThree(nums1[i], nums2[i], nums3[i]))
	}
}
