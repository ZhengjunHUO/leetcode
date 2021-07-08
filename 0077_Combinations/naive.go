package main

import (
	"fmt"
)

func backtrack(rslt *[][]int, n int, k int, curr []int, currIdx int) {
	if len(curr) == k {
		valid := make([]int, len(curr), cap(curr))
		copy(valid, curr)
		*rslt = append(*rslt, valid)
	}

	for i:=currIdx; i<=n; i++ {
		curr = append(curr, i)
		backtrack(rslt, n, k, curr, i+1)
		curr = curr[0:len(curr)-1]
	}
}

func combine(n int, k int) [][]int {
	var rslt [][]int
	backtrack(&rslt, n, k, []int{}, 1)
	return rslt 
}

func main() {
	fmt.Println(combine(4,2))
	fmt.Println(combine(1,1))
}
