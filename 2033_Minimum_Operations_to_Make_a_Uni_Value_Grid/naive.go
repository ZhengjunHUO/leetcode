package main

import (
	"fmt"
	"sort"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func minOperations(grid [][]int, x int) int {
	// convert matrix to an array
	m, n := len(grid), len(grid[0])
	arr := make([]int, m*n)
	for i := range grid {
		for j := range grid[i] {
			arr[i*n+j] = grid[i][j]
		}
	}

	// sort the array
	sort.SliceStable(arr, func (i, j int) bool {
		return arr[i] < arr[j]
	})

	var ret int
	// set to store the remainders
	set := make(map[int]struct{})
	// get the median of the array
	med := arr[len(arr)/2]
	for i := range arr {
		set[arr[i]%x] = struct{}{}
		ret += abs(arr[i]-med)/x
	}

	// not possible
	if len(set) > 1 {
		return -1
	}

	return ret
}

func main() {
	grid := [][][]int{[][]int{[]int{2,4}, []int{6,8}}, [][]int{[]int{1,5}, []int{2,3}}, [][]int{[]int{1,2}, []int{3,4}}}
	x := []int{2,1,2}
	for i := range grid {
		fmt.Println(minOperations(grid[i], x[i]))
	}
}
