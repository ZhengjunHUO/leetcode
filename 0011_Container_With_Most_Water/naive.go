package main

import (
	"fmt"
)

func maxArea(height []int) int {
	lp, rp, max, current := 0, len(height)-1, 0, 0

	for lp < rp {
		if height[lp]<height[rp] {
			current = (rp - lp) * height[lp]
			lp += 1			
		}else{
			current = (rp - lp) * height[rp]
			rp -= 1
		}

		if current > max {
			max = current
		}
	}

	return max
    
}

func main() {
	height := []int{1,8,6,2,5,4,8,3,7}
	fmt.Println(maxArea(height))

	height = []int{1,1}
	fmt.Println(maxArea(height))

	height = []int{4,3,2,1,4}
	fmt.Println(maxArea(height))

	height = []int{1,2,1}
	fmt.Println(maxArea(height))
}
