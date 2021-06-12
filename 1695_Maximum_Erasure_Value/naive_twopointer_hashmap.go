package main

import (
	"fmt"
)

// 在0003_Longest_Substring_Without_Repeating_Characters算法的基础上稍作修改即可
func maximumUniqueSubarray(nums []int) int {
	dict := make(map[int]int)	
	lp, rp, currVal, maxVal := 0, 0, 0 ,0
    
	for rp < len(nums) {
		curr := nums[rp]
		if _, ok := dict[curr]; ok {
			for lp < rp {
				currVal -= nums[lp]
				delete(dict, nums[lp])
				lp += 1 

				if nums[lp-1] == curr {
					break
				}
			}	
		}

		currVal += curr
		dict[curr] = rp
	
		if currVal > maxVal {
			maxVal = currVal
		}
		rp += 1
	}	

	return maxVal
}

func main() {
	fmt.Println(maximumUniqueSubarray([]int{4,2,4,5,6}))
	fmt.Println(maximumUniqueSubarray([]int{5,2,1,2,5,2,1,2,5}))
}
