package main

import (
	"fmt"
	"sort"
)

func threeSumSmaller(nums []int, target int) int {
	sort.Ints(nums)
	n, rslt := len(nums), 0

	for i:=0; i<n-2; i++ {
		lp, rp, diff := i+1, n-1, target - nums[i]
		for (lp < rp) {
			if sum := nums[lp] + nums[rp]; sum < diff {
				rslt += (rp - lp)
				lp += 1
			}else{
				rp -= 1	
			}
		} 		
	}
	
	return rslt
}

func main() {
	nums := []int{-2, 0, 1, 3}
	target := 2
	fmt.Println(threeSumSmaller(nums, target))
	
	nums = []int{-1, 0, 1, 2, 3}
	target = 3
	fmt.Println(threeSumSmaller(nums, target))
}
