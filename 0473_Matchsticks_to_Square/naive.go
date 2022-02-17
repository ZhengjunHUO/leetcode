package main

import (
	"fmt"
)

// 和0698_Partition_to_K_Equal_Sum_Subsets的思路相同，为dfs加上bitmask
func backtrack(sticks *[]int, used *[]bool, sides, idx, currsum, target int) bool {
	// 已经凑齐四条边，返回true (因为target是根据数列和除以4得来，不可能有未使用的stick)
	if sides == 0 {
		return true
	}

	// 当前边已经凑齐，处理下一条边
	if currsum == target {
		return backtrack(sticks, used, sides-1, 0, 0, target)
	}

	// 当前边还没有凑齐，从idx起遍历数列
	for i:=idx; i<len(*sticks); i++ {
		// 跳过使用过的stick和加起来大于边长的stick
		if ((*used)[i] || currsum + (*sticks)[i] > target) {
			continue
		}

		// 标记为使用过
		(*used)[i] = true
		if backtrack(sticks, used, sides, i+1, currsum+(*sticks)[i], target) {
			return true
		}
		// 去掉标记
		(*used)[i] = false
	}

	return false
}

func makesquare(matchsticks []int) bool {
	var sum int
	for i := range matchsticks {
		sum += matchsticks[i]
	}

	if sum%4 != 0 {
		return false
	}

	used := make([]bool, len(matchsticks))

	return backtrack(&matchsticks, &used, 4, 0, 0, sum/4)
}

func main() {
	matchsticks := [][]int{[]int{1,1,2,2,2}, []int{3,3,3,3,4}}
	for i := range matchsticks {
		fmt.Println(makesquare(matchsticks[i]))
	}
}
