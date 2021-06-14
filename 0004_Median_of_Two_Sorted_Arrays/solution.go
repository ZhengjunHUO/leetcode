package main

import (
	"fmt"
)

/*
@MissMary的阐述很到位
1）中位数的要点一是其左边的元素数量要等于右边的数量
所以n1数列中取了index i后，n2数列的index j也会定下来：j = (l1 + l2) / 2 - i

2）要点二是其左边的元素得小于右边的元素，每次检查某个i的时候会把两个数列分成四份
    左边两段 n1[0] ~ n1[i-1], n2[0] ~ n2[j-1] 和
    右边两段 n1[i] ~ n1[l1-1], n2[j] ~ n2[l2-1]
由于n1, n2本身是排好序的，所以只需要检查是否满足n1[i] >= n2[j-1]和n2[j] >= n1[i-1]即可
如果n1[i] < n2[j-1]说明当前i值取得太小，需要向右取更大的i
如果n1[i-1] > n2[j]说明当前i值太大，需要向左取更大的i

3）对于i的搜索因为已经排好序了可以使用二叉树进行（时间复杂度为O(logn)）；另外需要把短的那条数列作为n1来检索i，防止j的计算超出n2的range
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var n1, n2 []int
	var l1, l2 int
	var isOdd bool

	if len(nums1) > len(nums2) {
		n1, n2, l1, l2 = nums2, nums1, len(nums2), len(nums1)   
	}else{
		n1, n2, l1, l2 = nums1, nums2, len(nums1), len(nums2)
	}

	if (l1 + l2) % 2 == 1 {
		isOdd = true
	}

	var lmax, rmin int
	lp, rp, i, j := 0, l1, 0, 0
	half := (l1 + l2 + 1) / 2

	for lp <= rp {
		i = (lp + rp) / 2
		j = half - i

		if i < l1 && n1[i] < n2[j-1] {
			lp = i + 1
			continue
		}
		
		if i > 0 && n2[j] < n1[i-1] {
			rp = i - 1
			continue
		}

		break
	}

	if i == 0 {
		lmax = n2[j-1]
	}else if j == 0 {
		lmax = n1[i-1]
	}else{
		if n1[i-1] > n2[j-1] {
			lmax = n1[i-1]
		}else{
			lmax = n2[j-1]
		}
	}

	if isOdd {
		return float64(lmax)
	}	

	if i == l1 {
		rmin = n2[j]
	}else if j == l2 {
		rmin = n1[i]
	}else{
		if n1[i] > n2[j] {
			rmin = n2[j]
		}else{
			rmin = n1[i]
		}
	}

	return float64(lmax + rmin) / 2
	
}

func main() {
	fmt.Printf("%.5f\n", findMedianSortedArrays([]int{1,3}, []int{2}))
	fmt.Printf("%.5f\n", findMedianSortedArrays([]int{1,2}, []int{3,4}))
	fmt.Printf("%.5f\n", findMedianSortedArrays([]int{0,0}, []int{0,0}))
	fmt.Printf("%.5f\n", findMedianSortedArrays([]int{}, []int{1}))
	fmt.Printf("%.5f\n", findMedianSortedArrays([]int{2}, []int{}))
}
