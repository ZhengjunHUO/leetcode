package main 

import (
	"fmt"
)

func maxProfit(prices []int) int {
	prev, big1, big2 := 0, 0, 0

	for i:=1; i<len(prices); i++ {
		if prices[i] > prices[i-1] {
			if prev > 0 {
				prev += (prices[i] - prices[i-1])
			}else{
				prev = (prices[i] - prices[i-1])
			}
		}else{
			if prev > 0 {
				if prev > big1 {
					big1, big2 = prev, big1
				}else if prev > big2 {
					big2 = prev
				}
				prev = 0
			}
		}
	}
	
	if prev > 0 {
		if prev > big1 {
			big1, big2 = prev, big1
		}else if prev > big2 {
			big2 = prev
		}
	}

	return big1 + big2
}

func main() {
	prices := []int{3,3,5,0,0,3,1,4}
	fmt.Println(maxProfit(prices))

	prices = []int{1,2,3,4,5}
	fmt.Println(maxProfit(prices))

	prices = []int{7,6,4,3,1}
	fmt.Println(maxProfit(prices))

	prices = []int{1}
	fmt.Println(maxProfit(prices))
}
