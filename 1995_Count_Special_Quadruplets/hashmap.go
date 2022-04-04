package main

import (
	"fmt"
)

// @Kamikakushi's solution, find nums[a] + nums[b] = nums[d] - nums[c]
// scanning from the end to the begining, store nums[d] - nums[c] into a hashmap
// since a < b < c < d we can check if nums[a] + nums[b] in the hashmap, and in the same time enrich the hashmap
func countQuadruplets(nums []int) int {
	ret, n := 0, len(nums)

	dict := make(map[int]int)
	// put initial (nums[d] - nums[c]) in the map
	dict[nums[n-1]-nums[n-2]] = 1

	// b's range [1, n-3]
	for b:=n-3; b>=1; b-- {
		// a's range [0, b-1]
		for a:=b-1; a>=0; a-- {
			// if (nums[a]+nums[b]) is in the map, which meams == to some (nums[d] - nums[c])s
			ret += dict[nums[a]+nums[b]]
		}

		// calculate (nums[d] - nums[c]) to update the hashmap, b is used as 'c' here
		// d's range (b, n-1] 
		for d:=n-1; d>b; d-- {
			diff := nums[d] - nums[b]
			if _, ok := dict[diff]; ok {
				dict[diff] += 1
			}else{
				dict[diff] = 1
			}
		}
	}

	return ret
}

func main() {
	nums := [][]int{[]int{1,2,3,6}, []int{3,3,6,4,5}, []int{1,1,1,3,5}}
	for i := range nums {
		fmt.Println(countQuadruplets(nums[i]))
	}
}
