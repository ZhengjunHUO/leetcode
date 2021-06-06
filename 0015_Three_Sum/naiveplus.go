package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return [][]int{}
	}

	rslt := [][]int{}
	dict := make(map[[3]int]int)
	key := [3]int{}

	for i := range nums {
		target := 0 - nums[i]
		h := make(map[int]int)
		for j := range nums {
			if j == i {
				continue
			}
			
			diff := target - nums[j]	
			if elem, ok := h[diff]; ok {
				triple := []int{nums[i],nums[elem],nums[j]}
				sort.Ints(triple)
				copy(key[:], triple)
				if _, ok := dict[key]; !ok {
					dict[key] = 1
					rslt = append(rslt, triple)
				} 
			}else {
				h[nums[j]] = j
			}			
		}
	}

	return rslt
}

func main() {
	fmt.Println(threeSum([]int{-1,0,1,2,-1,4}))
	fmt.Println(threeSum([]int{}))
	fmt.Println(threeSum([]int{0}))
}
