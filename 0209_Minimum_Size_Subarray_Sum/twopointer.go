package main

import (
	"fmt"
)

func minSubArrayLen(target int, nums []int) int {
	lp,sum,len := 0,0,100000    

	for i,v := range nums {
		sum += v
		
		for sum >= target {
			temp := i - lp + 1 
			if temp < len {
				len = temp
			}

			sum -= nums[lp]
			lp += 1
		}
	}

	if len < 100000 {
		return len
	}
	return 0
}

func main() {
	nums := []int{2,3,1,2,4,3}
	target := 7
	fmt.Println(minSubArrayLen(target, nums))

	nums = []int{1,4,4}
	target = 4
	fmt.Println(minSubArrayLen(target, nums))

	nums = []int{1,1,1,1,1,1,1,1}
	target = 11
	fmt.Println(minSubArrayLen(target, nums))
}
