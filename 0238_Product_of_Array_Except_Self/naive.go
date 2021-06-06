package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	n := len(nums)
	rslt := make([]int, n)    
	rslt[0] = 1

	for i:=0; i<n-1; i++ {
		rslt[i+1] = rslt[i]*nums[i]	
	}

	prod := 1
	for i:=n-1; i>0; i-- {
		prod *= nums[i]
		rslt[i-1] *= prod
	}

	return rslt 
}

func main() {
	nums := []int{1,2,3,4}
	fmt.Println(productExceptSelf(nums))

	nums = []int{-1,1,0,-3,3}
	fmt.Println(productExceptSelf(nums))

	nums = []int{2,3,4,5,6,7}
	fmt.Println(productExceptSelf(nums))
}
