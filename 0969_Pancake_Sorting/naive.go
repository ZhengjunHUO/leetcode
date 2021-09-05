package main

import (
	"fmt"
)

func reverse(arr []int, l, r int) {
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}

func sort(arr []int, n int) []int {
	if n == 1 {
		return []int{}
	}

	rslt := []int{}

	maxVal, maxIdx := 0, 0
	for i := 0; i < n; i++ {
		if arr[i] > maxVal {
			maxVal, maxIdx = arr[i], i
		}
	}

	reverse(arr, 0, maxIdx)	
	rslt = append(rslt, maxIdx+1)
	reverse(arr, 0, n-1)
	rslt = append(rslt, n)

	rslt = append(rslt, sort(arr, n-1)...)
	return rslt
}

func pancakeSort(arr []int) []int {
	return sort(arr, len(arr))
}

func main() {
	arrs := [][]int{[]int{3,2,4,1}, []int{1,2,3}}
	for i := range arrs {
		fmt.Println(pancakeSort(arrs[i]))
	}
}
