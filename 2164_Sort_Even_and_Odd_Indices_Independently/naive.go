package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func sortEvenOdd(nums []int) []int {
	maxodd, mineven := godtype.NewPQ([]int{}, []int{}, false), godtype.NewPQ([]int{}, []int{}, true)

	for i := range nums {
		if i%2 == 1 {
			maxodd.Push(nums[i], nums[i])
		}else{
			mineven.Push(nums[i], nums[i])
		}
	}

	for i := range nums {
		if i%2 == 1 {
			nums[i] = maxodd.Pop().(int)
		}else{
			nums[i] = mineven.Pop().(int)
		}
	}

	return nums
}

func main() {
	nums := [][]int{[]int{4,1,2,3}, []int{2,1}, []int{4,3,2,7,8,5,3,1}}
	for i := range nums {
		fmt.Println(sortEvenOdd(nums[i]))
	}
}
