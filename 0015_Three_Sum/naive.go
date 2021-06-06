package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return [][]int{}
	}

	rslt := [][]int{}
	dict := make(map[[3]int]int)
	key := [3]int{}
	for i:=0; i<n-2; i++ {
		for j:=i+1; j<n-1; j++ {
			for k:=j+1; k<n; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					triple := []int{nums[i],nums[j],nums[k]}
					sort.Ints(triple)
					copy(key[:], triple)
					if _, ok := dict[key]; !ok {
						dict[key] = 1
						rslt = append(rslt, triple)
					} 
				}
			}
		}
	}

	return rslt

}

func main() {
	fmt.Println(threeSum([]int{-1,0,1,2,-1,4}))
	fmt.Println(threeSum([]int{}))
	fmt.Println(threeSum([]int{0}))
}
