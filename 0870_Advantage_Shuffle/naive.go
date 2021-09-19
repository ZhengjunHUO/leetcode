package main

import (
	"fmt"
	"sort"
	"github.com/ZhengjunHUO/godtype"
)

func advantageCount(nums1 []int, nums2 []int) []int {
	rslt := make([]int, len(nums2))
	for i := range rslt {
		rslt[i] = i
	}

	pq := godtype.NewPQ(rslt, nums2, false)
	sort.Ints(nums1)
	small, big := 0, len(nums1)-1

	for pq.Size() > 0 {
		max2 := pq.PopWithPrio()
		if nums1[big] > max2[1].(int) {
			// 如果nums1当前最大值大于nums2当前最大值，则使用num1的当前值
			rslt[max2[0].(int)] = nums1[big]
			big--	
		}else{
			// 如果nums1当前最大值小于nums2当前最大值，则使用一张当前最小的牌与之对决
			// nums1当前最大值留到下一轮再比
			rslt[max2[0].(int)] = nums1[small]
			small++
		}
	}

	return rslt
}

func main() {
	nums1 := [][]int{[]int{2,7,11,15}, []int{12,24,8,32}}
	nums2 := [][]int{[]int{1,10,4,11}, []int{13,25,32,11}}

	for i := range nums1 {
		fmt.Println(advantageCount(nums1[i], nums2[i]))
	}
}
