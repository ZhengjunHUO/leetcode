package main

import (
	"fmt"
)

func moveZeroes(nums []int)  {
	for i,p :=0,0 ; i<len(nums); i++ {
		if nums[i] != 0 {
			if i==p {
				p+=1
			}else{
				nums[p], nums[i] = nums[i], nums[p]
				p+=1
			}
		}
	}
}

func main() {
	nums := []int{0,1,0,3,12}
	fmt.Println("Before: ", nums)
	moveZeroes(nums)
	fmt.Println("After : ",nums)
	fmt.Println("======END======")

	nums = []int{4,0,1,0,3,12}
	fmt.Println("Before: ", nums)
	moveZeroes(nums)
	fmt.Println("After : ",nums)
	fmt.Println("======END======")
	
	nums = []int{4,1,3,12}
	fmt.Println("Before: ", nums)
	moveZeroes(nums)
	fmt.Println("After : ",nums)
	fmt.Println("======END======")
}
