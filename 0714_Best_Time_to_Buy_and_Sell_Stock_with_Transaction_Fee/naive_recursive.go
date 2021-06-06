package main

import (
	"fmt"
)

// 重用../0122_Best_Time_to_Buy_and_Sell_Stock_II/naive_recursive.go中的递归算法
// 增加新参数fee，在sell :=一行末尾加入 - fee即可
func maxProfit(prices []int, fee int) int {
	return maxProfitBis(prices, 10000, 0, fee)
}

func maxProfitBis(prices []int, min_buy int, i int, fee int) int {
	if i >= len(prices) {
		return 0
	}

	// 如果当前价格低于目前已知最低价，买入
	if prices[i] <= min_buy {
		return maxProfitBis(prices, prices[i], i+1, fee)
	}else{
		sell := prices[i] - min_buy + maxProfitBis(prices, 10000, i+1, fee) - fee
		notsell := maxProfitBis(prices, min_buy, i+1, fee)

		if sell > notsell {
			return sell
		}else{
			return notsell
		}
	}
}

func main() {
	prices := []int{1,3,2,8,4,9}
	fee := 2
	fmt.Println(maxProfit(prices, fee))

	prices = []int{1,3,7,5,10,3}
	fee = 3
	fmt.Println(maxProfit(prices, fee))
}
