package main

import (
	"fmt"
	"github.com/ZhengjunHUO/gtoolkit"
)

func searchInsert(nums []int, target int) int {
	return gtoolkit.SearchLeftBnd(nums, target)
}

func main() {
	test := []int{5, 2, 7, 0}
	for _,v := range test {
		fmt.Println(searchInsert([]int{1,3,5,6}, v))
	}

	fmt.Println(searchInsert([]int{1}, 0))

	test = []int{8, 2}
	for _,v := range test {
		fmt.Println(searchInsert([]int{1,3,4,7,9}, v))
	}
}
