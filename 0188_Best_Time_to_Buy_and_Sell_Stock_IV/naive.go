package main

import (
	"fmt"
	"sort"
)

func maxProfit(k int, prices []int) int {
	prev := 0
	maxk := make([]int, k)    

	for i:=1; i<len(prices); i++ {
		if prices[i] > prices[i-1] {
			if prev > 0 {
				prev += (prices[i] - prices[i-1])
			}else{
				prev = (prices[i] - prices[i-1])
			}
		}else{
			if prev > 0 {
//				fmt.Println("Find a value: ", prev)
				updateMax(&maxk, prev)
				prev = 0
			}
		}
	}
	
	if prev > 0 {
//		fmt.Println("Find a value: ", prev)
		updateMax(&maxk, prev)
	}

	sum := 0
//	fmt.Println("The result is: ", maxk)
	for _,v := range maxk {
		sum += v
	}

	return sum
}

func updateMax(x *[]int, y int) {
	idx := sort.SearchInts(*x, y)
//	fmt.Println("insert at: ", idx)
	switch idx {
	case 0:
		// Do nothing
	case 1:
		(*x)[0] = y
	default:
		copy((*x)[0:idx-1], (*x)[1:idx])
		(*x)[idx-1] = y
	}
}


func main() {
	k := 2
	prices := []int{2,4,1}
	fmt.Println(maxProfit(k, prices))

	prices = []int{3,2,6,5,0,3}
	fmt.Println(maxProfit(k, prices))

	prices = []int{3,3,5,0,0,3,1,4}
	fmt.Println(maxProfit(k, prices))

	prices = []int{1,2,3,4,5}
	fmt.Println(maxProfit(k, prices))

	prices = []int{7,6,4,3,1}
	fmt.Println(maxProfit(k, prices))
}


