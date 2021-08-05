package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func nextGreaterElements(nums []int) []int {
	mono, n := godtype.NewStack(), len(nums)
	rslt := make([]int, n)
	
	for i := 2*n - 1; i >= 0 ; i-- {
		for !mono.IsEmpty() && nums[i%n] >= mono.Peek().(int) {
			mono.Pop()
		}

		if mono.IsEmpty() {
			rslt[i%n] = -1
		}else{
			rslt[i%n] = mono.Peek().(int)
		}

		mono.Push(nums[i%n])
	}	

	return rslt
}

func main() {
	nums := [][]int{[]int{1,2,1}, []int{1,2,3,4,3}}
	for i := range nums {
		fmt.Println(nextGreaterElements(nums[i]))
	}
}
