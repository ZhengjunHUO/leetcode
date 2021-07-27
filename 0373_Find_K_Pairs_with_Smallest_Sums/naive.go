package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	rslt := [][]int{}
	n1, n2 := len(nums1), len(nums2)

	if n1 < 1 || n2 < 1 || k < 1 {
		return rslt
	}

	pq := godtype.NewPQ([][]int{}, []int{}, true)
	// 基础为nums1[0]和nums2[i]的组合
	for i := range nums2 {
		if i == k {
			break
		}
		pq.Push([]int{0, i}, nums1[0]+nums2[i])
	}

	for k > 0 && pq.Peek() != nil {
		curr := pq.Pop().([]int)
		rslt = append(rslt, []int{nums1[curr[0]], nums2[curr[1]]})
		k--

		// 防止push(i+1,j)超出nums1范围
		if curr[0] == n1 - 1 {
			continue
		}
		/* 
			在(i,j)后下一个可能的值是(i+1,j)或(i,j+1)
			但pq中肯定已经包含(i,x)且nums2[x]<=nums2[j+1]
			所以只需要加入(i+1,j)即可	
		*/
		pq.Push([]int{curr[0]+1, curr[1]}, nums1[curr[0]+1]+nums2[curr[1]])
	}

	return rslt
}

func main() {
	nums1 := [][]int{[]int{1,7,11}, []int{1,1,2}, []int{1,2}, []int{1,7,11,16}}
	nums2 := [][]int{[]int{2,4,6}, []int{1,2,3}, []int{3}, []int{2,9,10,15}}
	k := []int{3, 2, 3, 5}

	for i := range nums1 {
		fmt.Println(kSmallestPairs(nums1[i], nums2[i], k[i]))
	}
}
