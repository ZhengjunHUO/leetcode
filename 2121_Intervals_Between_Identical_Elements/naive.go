package main

import (
	"fmt"
)

func getDistances(arr []int) []int64 {
	ret := make([]int64, len(arr))
	dict := make(map[int][]int)

	for i := range arr {
		if _, ok := dict[arr[i]]; !ok {
			dict[arr[i]] = []int{i}
			continue
		}

		for j := range dict[arr[i]] {
			diff := int64(i - dict[arr[i]][j])
			ret[dict[arr[i]][j]] += diff
			ret[i] += diff
		}
		dict[arr[i]] = append(dict[arr[i]], i)
	}

	return ret
}

func main() {
	arrs := [][]int{[]int{2,1,3,1,2,3,3}, []int{10,5,10,10}}
	for i := range arrs {
		fmt.Println(getDistances(arrs[i]))
	}
}
