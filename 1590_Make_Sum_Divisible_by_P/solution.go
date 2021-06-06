package main

import (
	"fmt"
)

func minSubarray(nums []int, p int) int {
	sum := 0
	for _,v := range nums {
		sum += v
	}
	sum %= p
	if sum == 0 {
		return 0
	}

	rem := p - sum

	dict := make(map[int]int)    
	dict[0] = -1
	cumsum,size := 0, len(nums)

	for i,v := range nums {
		cumsum += v
		cumsum %= p
		
		if v, ok := dict[(rem+cumsum)%p]; ok {
			if (i - v) < size {
				size = i - v 
			}
		}	
		dict[cumsum] = i
	}
	
	if size<len(nums) {
		return size
	}
	return -1
}

func main() {
	nums := []int{3,1,4,2}
	p := 6
	fmt.Println(minSubarray(nums, p))

	nums = []int{6,3,5,2}
	p = 9
	fmt.Println(minSubarray(nums, p))

	nums = []int{1,2,3}
	p = 3
	fmt.Println(minSubarray(nums, p))

	nums = []int{1,2,3}
	p = 7
	fmt.Println(minSubarray(nums, p))

	nums = []int{1000000000,1000000000,1000000000}
	p = 3
	fmt.Println(minSubarray(nums, p))
}
