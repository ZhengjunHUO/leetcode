package main

import (
	"fmt"
)

// 重复利用../0309_Best_Time_to_Buy_and_Sell_Stock_with_Cooldown/statemachine.go的逻辑
// 这里没有cooldown机制，所以移除sellPreprev，原先用到的地方替换为sellPrev
// fee相当于增加股票购买的成本，于是和buy有关的操作均需要考虑fee（卖操作不变）
func maxProfit(prices []int, fee int) int {
	// 最近一次购买操作后累计可以获得的最大利润，初始值为-prices[0]即购买第一个单位后还没来得及卖出导致利润为负
	buyCurr, buyPrev := -(prices[0]+fee), -(prices[0]+fee)

	// 最近一次卖出操作后累计可以获得的最大利润，初始值为0即尚未来得及出货
	sellCurr, sellPrev := 0, 0

	// 从prices[1]开始遍历，初始化变量时已经把prices[0]“买入”了
	for i:=1; i<len(prices); i++ {
		// 上上轮卖出后累计利润减去买入所需的成本 比 不买进股票 取最大值
		if tempifbuy := sellPrev - prices[i] - fee; tempifbuy > buyPrev {
			buyCurr = tempifbuy
		}else{
			buyCurr = buyPrev
		}
		
		// 上轮买进利润加上卖出股票所得 比 持股不卖 取最大值
		if tempifsell := buyPrev + prices[i]; tempifsell > sellPrev {
			sellCurr = tempifsell
		}else{
			sellCurr = sellPrev
		}
	
		buyPrev, sellPrev = buyCurr, sellCurr
	}

	return sellCurr
}

func main() {
	prices := []int{1,3,2,8,4,9}
	fee := 2
	fmt.Println(maxProfit(prices, fee))

	prices = []int{1,3,7,5,10,3}
	fee = 3
	fmt.Println(maxProfit(prices, fee))
}
