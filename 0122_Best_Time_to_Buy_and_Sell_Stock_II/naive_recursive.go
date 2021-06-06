package main 

import (
	"fmt"
)

func maxProfit(prices []int) int {
	return maxProfitBis(prices, 10000, 0)    
}

func maxProfitBis(prices []int, min_buy int, i int) int {
	if i >= len(prices) {
		return 0
	}

	// 如果当前价格低于目前已知最低价，买入
	if prices[i] <= min_buy {
		return maxProfitBis(prices, prices[i], i+1)
	}else{
		sell := prices[i] - min_buy + maxProfitBis(prices, 10000, i+1)
		notsell := maxProfitBis(prices, min_buy, i+1)

		if sell > notsell {
			return sell
		}else{
			return notsell
		}
	}
}

func main() {
	prices := []int{7,1,5,3,6,4}
	fmt.Println(maxProfit(prices))

	prices = []int{1,2,3,4,5}
	fmt.Println(maxProfit(prices))

	prices = []int{7,6,4,3,1}
	fmt.Println(maxProfit(prices))
}
