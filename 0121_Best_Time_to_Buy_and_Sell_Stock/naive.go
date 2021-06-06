package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	min_buy, max_profit := 10000, 0

	for _,v := range prices {
		if v < min_buy {
			min_buy = v
		}
		if v > min_buy {
			if temp := v - min_buy; temp > max_profit {
				max_profit = temp
			}
		}
	}

	return max_profit
}

func main() {
	prices := []int{7,1,5,3,6,4}
	fmt.Println(maxProfit(prices))

	prices = []int{7,6,4,3,1}
	fmt.Println(maxProfit(prices))
}
