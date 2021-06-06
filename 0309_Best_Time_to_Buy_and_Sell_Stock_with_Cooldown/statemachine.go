package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	// 最近一次购买操作后累计可以获得的最大利润，初始值为-prices[0]即购买第一个单位后还没来得及卖出导致利润为负
	buyCurr, buyPrev := -prices[0], -prices[0]

	// 最近一次卖出操作后累计可以获得的最大利润，初始值为0即尚未来得及出货
	// 因为有cooldown机制所以需要额外的sellPreprev记录i-2时的数据
	sellCurr, sellPrev, sellPreprev := 0, 0, 0

	// 从prices[1]开始遍历，初始化变量时已经把prices[0]“买入”了
	for i:=1; i<len(prices); i++ {
		// 上上轮卖出后累计利润减去买入所需的成本 比 不买进股票 取最大值
		if tempifbuy := sellPreprev - prices[i]; tempifbuy > buyPrev {
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
	
		buyPrev, sellPrev, sellPreprev = buyCurr, sellCurr, sellPrev
	}

	return sellCurr
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
