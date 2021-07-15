package main

import (
	"fmt"
)

func topKFrequent(nums []int, k int) []int {
	dict := make(map[int]int)
	for _, v := range nums {
		dict[v] += 1
	}

	bucket := make([][]int, len(nums))	

	for k, v := range dict {
		if bucket[v] == nil {
			bucket[v] = []int{k}
		}else{
			bucket[v] = append(bucket[v], k)
		}
	}

	rslt := make([]int, 0, k)
	loop: for i := len(bucket) - 1; i >= 0; i-- {
		if bucket[i] != nil {
			for j := 0; j < len(bucket[i]); j++ {
				rslt = append(rslt, bucket[i][j])
				if len(rslt) == k {
					break loop
				}
			}
		}
	}

	return rslt
}

func main() {
	fmt.Println(topKFrequent([]int{1,1,1,2,2,3}, 2))
	fmt.Println(topKFrequent([]int{1,1,2,2,3,3,4,4,4,4,5}, 2))
}
