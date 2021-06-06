package main

import (
	"fmt"
)

func maxOperations(nums []int, k int) int {
	dict := make(map[int]int)
	rslt := 0

	for _,v := range nums {
		if val, ok := dict[k-v]; ok {
			if val > 1 {
				dict[k-v] = val - 1 
			}else{
				delete(dict, k-v)
			}
			rslt += 1 
		}else{
			dict[v] += 1
		}
	}    

	return rslt
}


func main() {
	nums1 := []int{1,2,3,4}
	target1 := 5

	nums2 := []int{3,1,3,4,3}
	target2 := 6	

	nums3 := []int{8,9,10,11,12,13,14,15,12}
	target3 := 24	

	fmt.Println(maxOperations(nums1,target1))
	fmt.Println(maxOperations(nums2,target2))
	fmt.Println(maxOperations(nums3,target3))

}
