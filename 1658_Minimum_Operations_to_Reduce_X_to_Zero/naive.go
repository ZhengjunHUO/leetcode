package main

import (
	"fmt"
)

func minOperations(nums []int, x int) int {
	n := len(nums)
	leftCumsum := make(map[int]int)    
	sum, ops := 0, 100000

	for i,v := range nums {
		sum += v
		if sum > x {
			break
		} 
		if sum == x {
			ops = i+1
			break
		}
		leftCumsum[sum] = i
	}

	sum = 0
	for i:=n-1; i>=0; i-- {
		sum += nums[i]
		if sum > x {
			break
		} 
		if val, ok := leftCumsum[x-sum]; ok {
			temp := (n - i) + (val + 1)
			if temp < ops {
				ops = temp
			}
		}
		if sum == x {
			if (n - i) < ops {
				ops = n - i 	
			} 
		}
	}

	if ops < 100000 {
		return ops
	}
	return -1
}

func main(){
	nums := []int{1,1,4,2,3}
	x := 5
	fmt.Println(minOperations(nums, x))

	nums = []int{5,6,7,8,9}
	x = 4
	fmt.Println(minOperations(nums, x))

	nums = []int{3,2,20,1,1,3}
	x = 10
	fmt.Println(minOperations(nums, x))
}
