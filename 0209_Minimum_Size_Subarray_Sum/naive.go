package main

import (
	"fmt"
)

func minSubArrayLen(target int, nums []int) int {
	cumsum := make(map[int]int)
	cumsum[0] = -1

	sum,len := 0,100000    

	for i,v := range nums {
		sum += v
		if sum >= target {
			for j:=0; j<=(sum-target);j++ {
				if val,ok := cumsum[j]; ok {
					if (i - val) < len {
						len = i - val
					}
				}	
			}
		}
		cumsum[sum] = i 
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
