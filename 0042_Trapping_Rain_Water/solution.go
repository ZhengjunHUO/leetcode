package main

import (
	"fmt"
)

func trap(height []int) int {
	lp, rp, lmax, rmax, cum := 0, len(height)-1, 0, 0, 0
	
	for (lp < rp) {
		if height[lp] <= height[rp] {
			if height[lp] >= lmax {
				lmax = height[lp]
			}else{
				cum += (lmax - height[lp])
			}
			lp += 1
		}else{
			if height[rp] >= rmax {
				rmax = height[rp]	
			}else{
				cum += (rmax - height[rp])
			}
			rp -= 1
		}	

	}

	return cum
}

func main() {
	height := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	fmt.Println(trap(height))

	height = []int{4,2,0,3,2,5}
	fmt.Println(trap(height))
}
