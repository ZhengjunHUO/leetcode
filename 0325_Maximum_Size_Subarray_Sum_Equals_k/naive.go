package main

import (
	"fmt"
)

func maxSubArrayLen(nums []int, k int) int {
	cumsum := make(map[int]int)
	cumsum[0] = -1
	
	sum, size := 0,0

	for i,v := range nums {
		sum += v
		if val, ok := cumsum[sum-k]; ok {
			temp := i - val
			if temp > size {
				size = temp
			}
		}

		if _, ok := cumsum[sum]; !ok {
			cumsum[sum] = i	
		}
	}
	
	return size

}

func main() {
	nums := []int{1, -1, 5, -2, 3}
	k := 3
	fmt.Println(maxSubArrayLen(nums,k))

	nums = []int{-2, -1, 2, 1}
	k = 1
	fmt.Println(maxSubArrayLen(nums,k))
}
