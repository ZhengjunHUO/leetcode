package main

import (
	"fmt"
	"sort"
)

// 排序后从大到小，以三颗糖为一组，花费为第一第二颗糖之和
func minimumCost(cost []int) int {
	sort.Ints(cost)
	var i,ret int

	for i = len(cost) - 1; i >= 2; i -= 3 {
		ret += cost[i] + cost[i-1]	
	}

	for ; i>=0; i-- {
		ret += cost[i]
	}

	return ret
}

func main() {
	costs := [][]int{[]int{1,2,3}, []int{6,5,7,9,2,2}, []int{5,5}}
	for i := range costs {
		fmt.Println(minimumCost(costs[i]))
	}
}
