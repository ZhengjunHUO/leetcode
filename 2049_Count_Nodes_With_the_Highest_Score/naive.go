package main

import (
	"fmt"
)

// 更新最大值，并统计出现次数
func updateMax(a int, max, count *int) {
	if a > (*max) {
		*max, *count = a, 1
	}else if a == (*max) {
		*count = (*count)+1
	}
}

// 递归调用，计算以当前节点为根的树的大小
func calculateSize(curr int, children map[int][]int, size *[]int) int {
	if (*size)[curr] != 0 {
		return (*size)[curr]
	}

	var ret int
	if v, ok := children[curr]; !ok {
		(*size)[curr] = 1
		return 1
	}else{
		ret = 1
		for j := range v {
			ret += calculateSize(v[j], children, size)
		}
	}

	(*size)[curr] = ret
	return ret
}

func countHighestScoreNodes(parents []int) int {
	// 存放每个节点的子节点
	children := make(map[int][]int)
	// 以每个节点为根的树的大小
	size := make([]int, len(parents))
	maxscore, count := 0, 0

	// 准备阶段: 计算子节点
	for i := range parents {
		if parents[i] != -1 {
			if _, ok := children[parents[i]]; ok {
				children[parents[i]] = append(children[parents[i]], i)
			}else{
				children[parents[i]] = []int{i}
			}
		}
	}

	// 准备阶段: 计算子树大小
	for i := range parents {
		if size[i] != 0 {
			continue
		}

		calculateSize(i, children, &size)
	}

	// 遍历节点计算最大值
	for i := range parents {
		if v, ok := children[i]; !ok {
			// 无子节点，即当前节点为叶子，移除后原树size减1
			updateMax(len(parents)-1, &maxscore, &count)
		}else{
			val := 1
			// 两个子树的大小的乘积
			for j := range v {
				val *= size[v[j]]
			}

			if parents[i] == -1 {
				// 当前节点是原树的根节点，直接返回上述乘积
				updateMax(val, &maxscore, &count)
			}else{
				// 当前节点是中间节点
				// 在上述乘积的基础上乘以原树size减去当前节点为根的子树的size
				val *= (size[parents[i]] - size[i])
				updateMax(val, &maxscore, &count)
			}
		}
	}

	return count
}

func main() {
	parents := [][]int{[]int{-1,2,0,2,0},[]int{-1,2,0}}
	for i := range parents {
		fmt.Println(countHighestScoreNodes(parents[i]))
	}
}
