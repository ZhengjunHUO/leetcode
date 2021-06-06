package main

import (
	"testing"
)

func BenchmarkMoveZeroes(b *testing.B) {
	nums := []int{4,0,1,0,3,12,0,67,0,0,44,0,24,13,56,0,9,0,100}
	for N:=0;N<b.N;N++ {
		for i,p :=0,0 ; i<len(nums); i++ {
			if nums[i] != 0 {
				nums[p], nums[i] = nums[i], nums[p]
				p+=1
			}
		}
	}
}

func BenchmarkMoveZeroesCntl(b *testing.B) {
	nums := []int{4,0,1,0,3,12,0,67,0,0,44,0,24,13,56,0,9,0,100}
	for N:=0;N<b.N;N++ {
		for i,p :=0,0 ; i<len(nums); i++ {
			if nums[i] != 0 {
				if i==p {
					p+=1
				}else{
					nums[p], nums[i] = nums[i], nums[p]
					p+=1
				}
			}
		}
	}
}
