package main 

import (
	"fmt"
)

// 取了../0122_Best_Time_to_Buy_and_Sell_Stock_II/naive_recursive.go中的递归方法
// 仅把sell := ... 中的maxProfitBis(prices, 10000, i+1)改为i+2（考虑售出后的一天冷静期）

func maxProfit(prices []int) int {
	return maxProfitBis(prices, 1000, 0)    
}

func maxProfitBis(prices []int, min_buy int, i int) int {
	if i >= len(prices) {
		return 0
	}

	// 如果当前价格低于目前已知最低价，买入
	if prices[i] <= min_buy {
		return maxProfitBis(prices, prices[i], i+1)
	}else{
		sell := prices[i] - min_buy + maxProfitBis(prices, 10000, i+2)
		notsell := maxProfitBis(prices, min_buy, i+1)

		if sell > notsell {
			return sell
		}else{
			return notsell
		}
	}
}

func main(){
	prices := []int{1,2,3,0,2}
	fmt.Println(maxProfit(prices))

	prices = []int{1}
	fmt.Println(maxProfit(prices))

	// need to return 6
	prices = []int{1,2,3,0,2,5}
	fmt.Println(maxProfit(prices))
}
