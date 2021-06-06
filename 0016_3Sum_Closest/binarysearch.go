package main

import (
	"fmt"
	"sort"
	"math"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n,diff := len(nums), 13000

	for i:=0; i<n-2; i++ {
		for j:=i+1; j<n-1; j++ {
			rest := target - nums[i] - nums[j]
			idx := sort.SearchInts(nums[j+1:], rest) + j+1
			// bigger
			if idx<n && math.Abs(float64(rest - nums[idx])) < math.Abs(float64(diff)) {
				diff = rest - nums[idx] 
			}
			// smaller
			if (idx - 1) > j && math.Abs(float64(rest - nums[idx-1])) < math.Abs(float64(diff)) {
				diff = rest - nums[idx-1] 
			} 	

		}
	}	
	return target - diff
}

func main() {
	nums := []int{-1,2,1,-4}
	target := 1
	fmt.Println(threeSumClosest(nums, target))

	nums = []int{-10,-7,2,8,13,22,36}
	target = 9
	fmt.Println(threeSumClosest(nums, target))
}
