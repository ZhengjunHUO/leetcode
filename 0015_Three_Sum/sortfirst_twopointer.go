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

	sort.Ints(nums)

	rslt := [][]int{}
	for i:=0; i<n-2; i++ {
		if nums[i] > 0 {
			break
		}

		if i==0 || nums[i] != nums[i-1] {
			target, lp, rp:= 0 - nums[i], i+1, n-1
			for lp < rp {
				if nums[lp] + nums[rp] == target {
					rslt = append(rslt, []int{nums[i],nums[lp],nums[rp]})
					for lp < rp && nums[lp] == nums[lp+1] {
						lp+=1
					}
					for lp < rp && nums[rp] == nums[rp-1] {
						rp-=1
					}
					lp+=1
					rp-=1
				} else if nums[lp] + nums[rp] > target {
					rp-=1
				} else {
					lp+=1
				}
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
