package extra0003

// 存储自fromDate那天结束后至今仓库中存放物品价值
type tuple struct {
	fromDate	int
	cumulVal	float32
}

func max(a, b float32) float32 {
	if a > b {
		return a
	}

	return b
}

/*
仓库每天会收到价值为value[i]的货物，每次去领取会产生cost的花费
仓库中的存货每天结束有risk的几率会被盗窃，求最值
*/
func MaxProfit(value []int, n, cost int, risk float32) float32 {
	var rslt float32

	// 第n天结束后仓库中存放物品价值的所有可能性
	// 不保存tuple{i, 0.0}的情况，即第i天前往领取了物品(故剩余价值为0)
	resid := make([][]tuple, n)
	for i:=0; i<n; i++ {
		resid[i] = []tuple{}
	}

	for i:=1; i<n; i++ {
		// case: 前一天领取过了(剩余价值为0)，当天积累的一天的价值(即当天到达的货物价值*损耗率)
		resid[i] = append(resid[i], tuple{i, float32(value[i-1])*(1-risk)})
		// case：前一天没有领（可能再之前领过或从来没有领过）的所有可能性，其剩余价值加上当天的价值后*损耗率
		for _, v := range resid[i-1] {
			resid[i] = append(resid[i], tuple{v.fromDate, (v.cumulVal+float32(value[i-1]))*(1-risk)})
		}
	}

	// 记录第n天去领可以获得的最大收益, 添加初值第0天dp[0]为0
	dp := make([]float32, n+1)
	// 从第1天开始计算
	for i:=1; i<=n; i++ {
		// 计算前一天仓库中所有可能剩余价值中的最大值
		// case: 前一天领取过了
		dp[i] = dp[i-1]
		// case: 前一天没有领取，查询resid表和已经计算过的dp表的部分来获得
		for _, tp := range resid[i-1] {
			val := dp[tp.fromDate - 1] + tp.cumulVal
			dp[i] = max(dp[i], val)
		}

		// 从获得的最值来计算当天去领可获得的最大收益
		dp[i] += float32(value[i-1] - cost)
		// 记录到目前为止的最大值(可能最后返回的最大收益不是dp最后一个元素)
		rslt = max(rslt, dp[i])
	}

	return rslt
}
