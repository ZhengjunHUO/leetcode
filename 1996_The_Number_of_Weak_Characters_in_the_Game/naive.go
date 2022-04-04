package main

import (
	"fmt"
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func numberOfWeakCharacters(properties [][]int) int {
	sort.SliceStable(properties, func (i, j int) bool{
		if properties[i][0] == properties[j][0] {
			return properties[i][1] > properties[j][1]
		}

		return properties[i][0] < properties[j][0]
	})

	//fmt.Println(properties)
	var ret int
	maxVal := properties[len(properties)-1][1]
	for i := len(properties)-2; i>=0; i-- {
		if properties[i][1] < maxVal {
			//fmt.Printf("properties[%d][1] = %d < %d\n", i, properties[i][1], maxVal)
			ret++
		}
		maxVal = max(maxVal, properties[i][1])
	}

	return ret
}

func main() {
	properties := [][][]int{[][]int{[]int{5,5},[]int{6,3},[]int{3,6}}, [][]int{[]int{2,2},[]int{3,3}}, [][]int{[]int{1,5},[]int{10,4},[]int{4,3}}}
	for i := range properties {
		fmt.Println(numberOfWeakCharacters(properties[i]))
	}
}
