package main

import (
	"fmt"
)

func frequencySort(nums []int) []int { 
	dict := make(map[int]int)
	for _,v := range nums {
		dict[v] += 1
	}

	bucket := make([][]int, len(nums))
	for k, v := range dict {
		if bucket[v] != nil {
			l := len(bucket[v])
			for i:=0; i<l; i++ {
				if k > bucket[v][i] {
					bucket[v] = append(bucket[v], 0)
					copy(bucket[v][i+1:], bucket[v][i:l])
					bucket[v][i] = k
					break
				}else{
					if i == l-1 {
						bucket[v] = append(bucket[v], k)
					}
				}
			}
		}else{
			bucket[v] = []int{k}
		}
	}

	rslt := make([]int, 0, len(nums))
	for i:=0; i<len(bucket); i++ {
		if bucket[i] != nil {
			for j := range bucket[i] {
				for k:=0; k<i; k++ {
					rslt = append(rslt, bucket[i][j])
				}
			}
		}
	}
	return rslt
}

func main() {
	ints := [][]int{[]int{1,1,2,2,2,3}, []int{2,3,1,3,2}, []int{-1,1,-6,4,5,-6,1,4,1}}
	for _,v := range ints {
		fmt.Println(frequencySort(v))
	}
}
