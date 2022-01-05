package main

import (
	"fmt"
	"sort"
)

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	// 把三个列表统合成一个三元组{start, end, point}的列表，并按结束时间升序排序
	slice := make([][]int, 0, len(profit))
	for i := range profit {
		slice = append(slice, []int{startTime[i], endTime[i], profit[i]})
	}
	sort.SliceStable(slice, func(i, j int) bool {
		return slice[i][1] < slice[j][1]
	})
	//fmt.Println(slice)

	// 辅助表，分别存放取得的工作的结束时间和累计的利润
	dp_endTime := []int{0}
	dp_profit := []int{0}

	for i := range slice {
		/* debug
		fmt.Println("dp_endTime: ", dp_endTime)
		fmt.Println("dp_profit: ", dp_profit)
		fmt.Printf("sort.SearchInts(%d, %d+1)\n", dp_endTime, slice[i][0])
		*/

		// 寻找"小于当前工作开始时间"的最大结束时间对应的index
		idx := sort.SearchInts(dp_endTime, slice[i][0]+1) - 1
		//fmt.Println(idx)
		
		// 接下当前工作: 计算index对应的累积利润+当前工作的收入
		take := dp_profit[idx] + slice[i][2]
		// 如果大于不接工作（既得的当前最大利润）则更新两个辅助表
		if take > dp_profit[len(dp_profit)-1] {
			//fmt.Println("take: ", take)
			// 加入当前工作的结束时间
			dp_endTime = append(dp_endTime, slice[i][1])
			// 加入当前工作完成后累积利润
			dp_profit = append(dp_profit, take)
		}
	}

	return dp_profit[len(dp_profit)-1]
}

func main() {
	startTime := [][]int{[]int{1,2,3,3},[]int{1,2,3,4,6},[]int{1,1,1}}
	endTime := [][]int{[]int{3,4,5,6},[]int{3,5,10,6,9},[]int{2,3,4}}
	profit := [][]int{[]int{50,10,40,70},[]int{20,20,100,70,60},[]int{5,6,4}}
	for i := range profit {
		fmt.Println(jobScheduling(startTime[i],endTime[i],profit[i]))
	}
}
