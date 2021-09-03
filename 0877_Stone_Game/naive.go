package main

import (
	"fmt"
)

/*
把所有石头堆编号分为奇数堆和偶数堆，可以观察是奇数堆的值之和大还是偶数堆的大
由于只能从两边拿，先拿的人可以利用优势选择拿和比较大的堆
*/
func stoneGame(piles []int) bool {
	return true	
}

func main() {
	piles := [][]int{[]int{5,3,4,5}, []int{3,7,2,3}}
	for i := range piles {
		fmt.Println(stoneGame(piles[i]))
	}
}
