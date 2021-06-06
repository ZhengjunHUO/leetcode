package main

import (
	"fmt"
)

type triplet struct {
	freq	int
	first	int
}

func findShortestSubArray(nums []int) int {
	dict := make(map[int]*triplet)

	max_occr, min_len := 0, 100000 
	for i,val := range nums {
		if v, ok := dict[val]; ok {
			v.freq += 1
			if v.freq > max_occr {
				max_occr, min_len = v.freq, i - v.first + 1
			}else if v.freq == max_occr {
				if temp := i - v.first + 1; temp < min_len {
					min_len = temp
				}
			}
		}else{
			dict[val] = &triplet{
				freq: 1,
				first: i,
			}
		}	
	}

	return min_len
}

func main() {
	nums := []int{1,2,2,3,1}
	fmt.Println(findShortestSubArray(nums))

	nums = []int{1,2,2,3,1,4,2}
	fmt.Println(findShortestSubArray(nums))
}
