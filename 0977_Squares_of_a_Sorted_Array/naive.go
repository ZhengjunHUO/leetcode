package main

import (
	"fmt"
)

func sortedSquares(nums []int) []int {
	n := len(nums)
	rslt := make([]int, n) 

	lp, rp, i := 0, n-1, n-1
	
	for lp <= rp {
		if lsq, rsq := nums[lp]*nums[lp], nums[rp]*nums[rp]; lsq < rsq {
			rslt[i] = rsq
			rp -= 1
		}else{
			rslt[i] = lsq
			lp += 1
		}
		i -= 1
	}

	return rslt
}

func main() {
	fmt.Println(sortedSquares([]int{-4,-1,0,3,10}))
	fmt.Println(sortedSquares([]int{-7,-3,2,3,11}))
}
