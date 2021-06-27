package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int)  {
	p1, p2, i := m-1, n-1, m+n-1

	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] < nums2[p2] {
			nums1[i] = nums2[p2]
			p2 -= 1
		}else{
			nums1[i] = nums1[p1]
			p1 -= 1
		}
		i -= 1
	}

	if p2 >= 0 {
		nums1[i] = nums2[p2]
		p2 -= 1
		i -= 1
	}
}

func main() {
	nums1 := []int{1,2,3,0,0,0}
	merge(nums1, 3, []int{2,5,6}, 3)
	fmt.Println(nums1)

	nums1 = []int{1}
	merge(nums1, 1, []int{}, 0)
	fmt.Println(nums1)

	nums1 = []int{0}
	merge(nums1, 0, []int{1}, 1)
	fmt.Println(nums1)
}
