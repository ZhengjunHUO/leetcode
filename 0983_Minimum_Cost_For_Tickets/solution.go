package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

/*
func mincostTickets(days []int, costs []int) int {
	n := len(days)
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = 1001
	}
	dp [0] = 0
	validate := []int{1, 7, 30}

	for i:=1; i<=n; i++ {
		for j := range costs {
			if validate[j] <= days[i-1] {
				if days[i-1]-validate[j] > i - 1 {
					dp[i] = min(dp[i], dp[i-1]+costs[j])
				}else{
					dp[i] = min(dp[i], dp[days[i-1]-validate[j]]+costs[j])
				}
				fmt.Println(dp)
			}
		}
	}

	return dp[n]
}
*/
func mincostTickets(days []int, costs []int) int {
	// current cost
	cost := 0
	// element: [2]int{travalingDay, costSinceBegin}
	last7days, last30days := datastruct.NewQueue([][2]int{}), datastruct.NewQueue([][2]int{})

	for i := range days {
		// check the coverage/validity of the 7 days ticket, remove them if expired
		for (!last7days.IsEmpty() && days[i] >= (last7days.Peek())[0] + 7) {
			last7days.Pop()
		}
		// check the coverage/validity of the 30 days ticket, remove them if expired
		for (!last30days.IsEmpty() && days[i] >= (last30days.Peek())[0] + 30) {
			last30days.Pop()
		}
		// if we buy 7/30 days ticket today 
		last7days.Push([2]int{days[i], cost+costs[1]})
		last30days.Push([2]int{days[i], cost+costs[2]})

		// compare the different solutions and update the current cost
		cost = min(cost+costs[0], min((last7days.Peek())[1], (last30days.Peek())[1]))
	}

	return cost
}

func main() {
	days := [][]int{[]int{1,4,6,7,8,20}, []int{1,2,3,4,5,6,7,8,9,10,30,31}}
	costs := [][]int{[]int{2,7,15}, []int{2,7,15}}

	for i := range days {
		fmt.Println(mincostTickets(days[i], costs[i]))
	}
}
