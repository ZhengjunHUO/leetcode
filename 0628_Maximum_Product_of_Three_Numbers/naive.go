package main

import (
	"fmt"
)

func maximumProduct(nums []int) int {
	max1, max2, max3, min1, min2 := -1000, -1000, -1000, 1000, 1000

	for _,v := range(nums) {
		if v > max1 {
			max1, max2, max3 = v, max1, max2
		}else if v > max2 {
			max2, max3 = v, max2
		}else if v > max3 {
			max3 = v
		}

		if v < min1 {
			min1, min2 = v, min1
		}else if v < min2 {
			min2 = v 
		}
	}

	if prod1, prod2 := max1*max2*max3, max1*min1*min2; prod1 > prod2 {
		return prod1
	}else{
		return prod2
	}
}

func main() {
	nums := []int{1,2,3}
	fmt.Println(maximumProduct(nums))

	nums = []int{1,2,3,4}
	fmt.Println(maximumProduct(nums))

	nums = []int{-1,-2,-3}
	fmt.Println(maximumProduct(nums))
}
