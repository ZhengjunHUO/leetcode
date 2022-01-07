package main

import (
	"fmt"
	"sort"
)

// 思路同1235_Maximum_Profit_in_Job_Scheduling/dp.go
func maxTaxiEarnings(n int, rides [][]int) int64 {
	// 把路费加入到小费中
	for i := range rides {
		rides[i][2] += rides[i][1] - rides[i][0]
	}

	// 按终点升序排序
	sort.SliceStable(rides, func(i, j int) bool {
		if rides[i][1] == rides[j][1] {
			return rides[i][0] < rides[j][0]
		}
		return rides[i][1] < rides[j][1]
	})

	// 动态编程辅助表，记录下车位置和累积盈利
	dp_end := []int{0}
	dp_earn := []int{0}

	for i := range rides {
		// 用bs寻找"小于当前上车点"的最大历史下车点
		idx := sort.SearchInts(dp_end, rides[i][0]+1) - 1
		// 如果"在该点选择载客获得的利润"大于"不载客当前最大累积利润"则更新dp表
		pick := dp_earn[idx] + rides[i][2]
		if pick > dp_earn[len(dp_earn)-1] {
			//fmt.Printf("[DEBUG] pick up client {%d, %d, %d}\n", rides[i][0], rides[i][1], rides[i][2])
			dp_end = append(dp_end, rides[i][1])
			dp_earn = append(dp_earn, pick)
			//fmt.Println("[DEBUG] dp_end: ", dp_end)
			//fmt.Println("[DEBUG] dp_earn: ", dp_earn)
		}
	}

	return int64(dp_earn[len(dp_earn)-1])
}

func main() {
	rides := [][][]int{[][]int{[]int{2,5,4},[]int{1,5,1}}, [][]int{[]int{1,6,1},[]int{3,10,2},[]int{10,12,3},[]int{11,12,2},[]int{12,15,2},[]int{13,18,1}}}
	n := []int{5, 20}

	for i := range n {
		fmt.Println(maxTaxiEarnings(n[i], rides[i]))
	}
}
