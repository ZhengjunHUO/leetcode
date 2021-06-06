package main

import (
	"sort"
	"testing"
)

func BenchmarkThreeSumNaive(b *testing.B) {
	nums := []int{-1,0,1,2,-1,4}
	n := len(nums)
	for N:=0; N<b.N; N++ {
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
	}
}

func BenchmarkThreeSumNaivePlus(b *testing.B) {
	nums := []int{-1,0,1,2,-1,4}
	for N:=0; N<b.N; N++ {
		rslt := [][]int{}
		dict := make(map[[3]int]int)
		key := [3]int{}
	
		for i := range nums {
			target := 0 - nums[i]
			h := make(map[int]int)
			for j := range nums {
				if j == i {
					continue
				}
				
				diff := target - nums[j]	
				if elem, ok := h[diff]; ok {
					triple := []int{nums[i],nums[elem],nums[j]}
					sort.Ints(triple)
					copy(key[:], triple)
					if _, ok := dict[key]; !ok {
						dict[key] = 1
						rslt = append(rslt, triple)
					} 
				}else {
					h[nums[j]] = j
				}			
			}
		}
	}
}

func BenchmarkSortfirst(b *testing.B) {
	nums := []int{-1,0,1,2,-1,4}
	n := len(nums)
	for N:=0; N<b.N; N++ {
		sort.Ints(nums)
		rslt := [][]int{}
		for i:=0; i<n-2; i++ {
			if nums[i] > 0 {
				break
			}
	
			if i==0 || nums[i] != nums[i-1] {
				target, lp, rp:= 0 - nums[i], i+1, n-1
				for lp < rp {
					if nums[lp] + nums[rp] == target {
						rslt = append(rslt, []int{nums[i],nums[lp],nums[rp]})
						for lp < rp && nums[lp] == nums[lp+1] {
							lp+=1
						}
						for lp < rp && nums[rp] == nums[rp-1] {
							rp-=1
						}
						lp+=1
						rp-=1
					} else if nums[lp] + nums[rp] > target {
						rp-=1
					} else {
						lp+=1
					}
				}
			}
		}
	}
}



