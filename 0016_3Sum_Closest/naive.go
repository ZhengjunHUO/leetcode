package main

import (
	"fmt"
	"sort"
	"math"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)	
	n,diff := len(nums), 13000

	for i:=0; i<n-2 && diff!=0; i++ {
		lp, rp := i+1, n-1
		for (lp<rp) {
			sum := nums[i]+nums[lp]+nums[rp]
			tmp := target - sum

			if math.Abs(float64(tmp)) < math.Abs(float64(diff)) {
				diff = tmp

			}

			if sum > target {
				rp -= 1
			}else{
				lp += 1
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
