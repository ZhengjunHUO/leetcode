package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	lp, rp := 0, len(numbers)-1
	for (lp < rp) {
		sum:=numbers[lp]+numbers[rp]
		switch {
		case sum==target:
			return []int{lp+1,rp+1}
		case sum<target:
			lp += 1
		case sum>target:
			rp -= 1
		}
	}

	fmt.Println("Error in problem")
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2,7,11,15},9))
	fmt.Println(twoSum([]int{2,3,4},6))
	fmt.Println(twoSum([]int{-1,0},-1))
	fmt.Println(twoSum([]int{1,4,7,9,15,18,22},25))
	fmt.Println(twoSum([]int{1,4,7,9,15,18,22},11))
}
