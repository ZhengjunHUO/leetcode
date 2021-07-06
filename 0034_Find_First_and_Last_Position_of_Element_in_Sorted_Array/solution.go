package main

import (
	"fmt"
)

// 借用0035_Search_Insert_Position的解法寻找index最小的插入点
func searchInsert(nums []int, target int) int {
        l, r := 0, len(nums)

        for l < r {
                m := l + (r - l)/2
                if nums[m] >= target {
                        r = m
                }else{
                        l = m + 1
                }
        }

        return l
}

func searchRange(nums []int, target int) []int {
	lp, rp := searchInsert(nums, target), searchInsert(nums, target+1)-1    

	if lp < len(nums) && nums[lp] == target {
		return []int{lp, rp}	
	}else{
		return []int{-1, -1}
	}
}

func main() {
	fmt.Println(searchRange([]int{5,7,7,8,8,10}, 8))
	fmt.Println(searchRange([]int{5,7,7,8,8,10}, 6))
	fmt.Println(searchRange([]int{}, 0))
	fmt.Println(searchRange([]int{3,4,5,8,9}, 5))
}
