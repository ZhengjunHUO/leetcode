package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

// 存入邻接表中的元素：相邻节点的坐标 和 到达需要花费的气力
type elem struct {
	cord	[2]int
	effort	int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func minimumEffortPath(heights [][]int) int {
	m, n := len(heights), len(heights[0])

	// 方向矩阵: 上右下左
	direct := [][2]int{[2]int{0,-1}, [2]int{1,0}, [2]int{0,1}, [2]int{-1,0}}

	// 制作邻接表
	dict := make(map[[2]int][]elem)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := range direct {
				x, y := i + direct[k][0], j + direct[k][1]
				if x >= 0 && x < m && y >= 0 && y < n {
					dict[[2]int{i,j}] = append(dict[[2]int{i,j}], elem{[2]int{x,y}, abs(heights[i][j]-heights[x][y])})
				}
			}
		}
	}

	// 准备气力花费表
	costTo := make([]int, m*n)
	for i := range costTo {
		costTo[i] = 1000000
	}
	costTo[0] = 0

	// 优先队列 
	pq := godtype.NewPQ([][2]int{[2]int{0,0}}, []int{0}, true)
	for pq.Size() > 0 {
		curr := pq.PopWithPrio()
		currCord, costToCurr := curr[0].([2]int), curr[1].(int)
		currIdx := currCord[0]*n+currCord[1]

		if currIdx == len(costTo)-1 {
			return costToCurr
		}

		// 如到达当前节点所需气力大于记录中对应的值，则忽略这条分支
		if costToCurr > costTo[currIdx] {
			continue
		}

		if v, ok := dict[currCord]; ok {
			for i := range v {
				nextIdx := v[i].cord[0]*n+v[i].cord[1]

				// 到达邻居节点所花费的气力，需要考虑从开始到当前的值 和 当前到邻居节点所需的值
				costToNext := max(costTo[currIdx], v[i].effort)
				// 如果小于记录中的值，更新该值，并加入优先队列
				if costToNext < costTo[nextIdx] {
					costTo[nextIdx] = costToNext
					pq.Push(v[i].cord, costToNext)
				}
			}
		}
	}

	return costTo[len(costTo)-1]
}

func main() {
	hs := [][][]int{[][]int{[]int{1,2,2},[]int{3,8,2},[]int{5,3,5}}, [][]int{[]int{1,2,3},[]int{3,8,4},[]int{5,3,5}}, [][]int{[]int{1,2,1,1,1},[]int{1,2,1,2,1},[]int{1,2,1,2,1},[]int{1,2,1,2,1},[]int{1,1,1,2,1}}}
	for i := range hs {
		fmt.Println(minimumEffortPath(hs[i]))
	}
}
