package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

/*
  参考0304_Range_Sum_Query_2D_Immutable和@votrubac的算法
  calcPrefixSum对每个坐标(i,j)计算其左上方的所有元素之和，注意返回的结果有padding
  rangeSum接受的对象为有padding的矩阵，计算由左上角和右下角确定的矩形区域中的元素和，注意接收的坐标为不考虑padding的坐标，需要转换
*/
func calcPrefixSum(grid *[][]int, sum *[][]int) {
	m, n := len(*grid), len((*grid)[0])
	*sum = make([][]int, m+1)
	for i := range *sum {
		(*sum)[i] = make([]int, n+1)
	}

	for i:=1; i<=m; i++ {
		for j:=1; j<=n; j++ {
			// i, j在grid中需要-1
			(*sum)[i][j] = (*sum)[i][j-1] + (*sum)[i-1][j] - (*sum)[i-1][j-1] + (*grid)[i-1][j-1]
		}
	}
}

func rangeSum(grid *[][]int, row1, col1, row2, col2 int) int {
	// grid为padding后的矩阵, 坐标需要+1来适应
	return (*grid)[row2+1][col2+1] - (*grid)[row2+1][col1] - (*grid)[row1][col2+1] + (*grid)[row1][col1]
}

func possibleToStamp(grid [][]int, stampHeight int, stampWidth int) bool {
	m, n := len(grid), len(grid[0])

	var sum [][]int
	// 对原始矩阵计算累计和矩阵，为下一步rangeSum作准备
	calcPrefixSum(&grid, &sum)
	//fmt.Println("sum:")
	//fmt.Println(sum)

	mark := make([][]int, m)
	for i := range mark {
		mark[i] = make([]int, n)
	}

	for i := stampHeight-1; i < m; i++ {
		for j := stampWidth-1; j < n; j++ {
			//fmt.Printf("rangeSum &sum, %d, %d, %d, %d\n", i-stampHeight+1, j-stampWidth+1, i, j)
			// 检查所有可以打戳的位置，借助上一步得到的prefix矩阵计算每个stamp的区域内是否有已标记的点
			// 如果没有已标记点则其值应该为0，用stamp区域右下角的坐标值取1来表示ok
			if rangeSum(&sum, i-stampHeight+1, j-stampWidth+1, i, j) == 0 {
				mark[i][j] = 1
			}
		}
	}
	//fmt.Println("mark:")
	//fmt.Println(mark)

	// 对能否打戳矩阵计算累计和矩阵，为下一步rangeSum作准备
	var markSum [][]int
	calcPrefixSum(&mark, &markSum)
	//fmt.Println("markSum:")
	//fmt.Println(markSum)

	for i := range grid {
		for j := range grid[i] {
			//fmt.Printf("rangeSum &markSum, %d, %d, %d, %d\n", i, j, min(i+stampHeight-1,m-1), min(j+stampWidth-1,n-1))
			// 如果原始矩阵(i,j)处本身为空值，且不能被任何stamp覆盖(rangSum为0)，则返回false
			if grid[i][j] == 0 && rangeSum(&markSum, i, j, min(i+stampHeight-1,m-1), min(j+stampWidth-1,n-1)) == 0 {
				return false
			}
		}
	}

	return true
}

func main() {
	grids := [][][]int{[][]int{[]int{1,0,0,0},[]int{1,0,0,0},[]int{1,0,0,0},[]int{1,0,0,0},[]int{1,0,0,0}}, [][]int{[]int{1,0,0,0},[]int{0,1,0,0},[]int{0,0,1,0},[]int{0,0,0,1}}}
	h := []int{4,2}
	w := []int{3,2}
	for i := range grids {
		fmt.Println(possibleToStamp(grids[i], h[i], w[i]))
	}
}
