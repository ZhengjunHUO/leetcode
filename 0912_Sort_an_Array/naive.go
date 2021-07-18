package main

import (
	"fmt"
)

// 和0215_Kth_Largest_Element_in_an_Array/quickselect.go中的解法类似
func divide(nums []int, l int, r int) int {
	pivot := nums[l]	

	for l < r {
		for l < r && nums[r] >= pivot {
			r--
		}
		nums[l], nums[r] = nums[r], nums[l]

		for l < r && nums[l] <= pivot {
			l++
		}
		nums[l], nums[r] = nums[r], nums[l]
	}

	return l
}

func quicksort(nums []int, l int, r int) {
	if l >= r {
		return
	}

	// quickselect确定一个pivot的最终index
	m := divide(nums, l, r)
	// 在pivot左侧继续quicksort
	quicksort(nums, l, m-1)
	// 在pivot右侧继续quicksort
	quicksort(nums, m+1, r)
}

func sortArray(nums []int) []int {
	quicksort(nums, 0, len(nums)-1)	
	return nums	
}

func main() {
	fmt.Println(sortArray([]int{5,2,3,1}))
	fmt.Println(sortArray([]int{5,1,1,2,0,0}))
	fmt.Println(sortArray([]int{2,5,8,10,3,1,7}))
}
