package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	var curr, idx, ret int
	accuSum := make([]int, fruits[len(fruits)-1][0]+1)

	// calculate the prefix sum
	for i := range fruits {
		for idx < fruits[i][0] {
			accuSum[idx] = curr
			idx++
		}

		curr += fruits[i][1]
	}
	accuSum[idx] = curr
	//fmt.Println(accuSum)

	for i:=0; i*2 < k; i++ {
		//fmt.Println("i: ", i)
		// go left for i step(can be 0), then go right
		l := max(startPos - i - 1, 0)
		r := min(startPos + (k - 2*i), len(accuSum)-1)
		//fmt.Printf("max(%d, %d[%d] - %d[%d])\n", ret, accuSum[r], r, accuSum[l], l)
		ret = max(ret, accuSum[r] - accuSum[l])

		// go right then left
		r = min(startPos + i, len(accuSum)-1)
		l = max(startPos - (k - 2*i) - 1, 0)
		//fmt.Printf("max(%d, %d[%d] - %d[%d])\n", ret, accuSum[r], r, accuSum[l], l)
		ret = max(ret, accuSum[r] - accuSum[l])
	}

	return ret
}

func main() {
	fruits := [][][]int{[][]int{[]int{2,8},[]int{6,3},[]int{8,6}}, [][]int{[]int{0,9},[]int{4,1},[]int{5,7},[]int{6,2},[]int{7,4},[]int{10,9}}, [][]int{[]int{0,3},[]int{6,4},[]int{8,5}}}
	startPos := []int{5, 5, 3}
	k := []int{4, 4, 2}

	for i := range fruits {
		fmt.Println(maxTotalFruits(fruits[i], startPos[i], k[i]))
	}
}
