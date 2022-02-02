package main

import (
	"fmt"
)

func inSet(dict map[int]struct{}, num int) bool {
	if _, ok := dict[num]; ok {
		return true
	}

	return false
}

func findLonely(nums []int) []int {
	dict, rslt_set := make(map[int]struct{}), make(map[int]struct{})

	for i := range nums {
		if inSet(dict, nums[i]) ||  inSet(dict, nums[i]-1)|| inSet(dict, nums[i]+1) {
			delete(rslt_set, nums[i])
			delete(rslt_set, nums[i]-1)
			delete(rslt_set, nums[i]+1)
		}else{
			rslt_set[nums[i]] = struct{}{}
		}

		dict[nums[i]] = struct{}{}
	}

	ret := make([]int, 0, len(rslt_set))
	for k, _ := range rslt_set {
		ret = append(ret, k)
	}
	return ret
}

func main() {
	nums := [][]int{[]int{10,6,5,8}, []int{1,3,5,3}}
	for i := range nums {
		fmt.Println(findLonely(nums[i]))
	}
}
