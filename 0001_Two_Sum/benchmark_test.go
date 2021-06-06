package main

import (
	"testing"
)

func BenchmarkNaive(b *testing.B) {
	nums := []int{13,2,7,11,15}
	target := 22
	for N:=0; N<b.N; N++ {
		h := make(map[int]int)
        	for i := range nums {
			diff := target - nums[i]
			if _, ok := h[diff]; ok {
				continue
			}
				
			h[nums[i]] = i
		}
	}
}

func BenchmarkHash(b *testing.B) {
	nums := []int{13,2,7,11,15}
	target := 22
	for N:=0; N<b.N; N++ {
		base := make([]int, len(nums))   
		for i := range base {
			base[i] = target - nums[i]
		}

        	for j:= range base {
			for k := range nums {
				if base[j] == nums[k] && j !=k {
					continue
				}
			} 
		}
	}
}     

