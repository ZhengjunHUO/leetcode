package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	mono := godtype.NewStack()
	dict := make(map[int]int)

	for i := len(nums2) - 1; i >= 0; i-- {
		for !mono.IsEmpty() && nums2[i] > mono.Peek().(int) {
			mono.Pop()
		}

		if mono.IsEmpty() {
			dict[nums2[i]] = -1
		}else{
			dict[nums2[i]] = mono.Peek().(int)
		}

		mono.Push(nums2[i])
	}

	for i := range nums1 {
		nums1[i] = dict[nums1[i]]
	}

	return nums1
}

func main() {
	nums1 := [][]int{[]int{4,1,2}, []int{2,4}}
	nums2 := [][]int{[]int{1,3,4,2}, []int{1,2,3,4}}
	for i := range nums1 {
		fmt.Println(nextGreaterElement(nums1[i], nums2[i]))
	}
}
