package main

import (
	"fmt"
	"sort"
)

func maxEnvelopes(envelopes [][]int) int {
	// 按照宽度升序排列，对于同等宽度的按高度降序排列
	sort.SliceStable(envelopes, func(i, j int) bool {
		if envelopes[i][0] != envelopes[j][0] {
			return envelopes[i][0] < envelopes[j][0]
		}
		return envelopes[i][1] > envelopes[j][1]
	})

	// 只取高度作为数组，求其LIS
	list := make([]int, len(envelopes))
	for i := range list {
		list[i] = envelopes[i][1]
	}

	top, piles := make([]int, len(list)), 0
	for i := range list {
		l, r := 0, piles
		for l < r {
			m := l + (r-l)/2
			if list[i] <= top[m] {
				r = m
			}else{
				l = m + 1
			}
		}

		if l == piles {
			piles++
		}

		top[l] = list[i]
	}

	return piles
}

func main() {
	envs := [][][]int{[][]int{[]int{5,4},[]int{6,4},[]int{6,7},[]int{2,3}}, [][]int{[]int{1,1},[]int{1,1},[]int{1,1}}, [][]int{[]int{5,4}, []int{6,7}, []int{1,8}, []int{6,4}, []int{2,3}, []int{5,2}}}
	for i := range envs {
		fmt.Println(maxEnvelopes(envs[i]))
	}
}
