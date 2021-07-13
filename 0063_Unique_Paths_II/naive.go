package main

import (
	"fmt"
)

/* 参考0062_Unique_Paths/naive.go中的DP思路
按照下面实现的算法如果3*4的矩阵在(1,2)有障碍物，会得到
_ 2 1 1
_ 1 0 1
_ 1 1 1
结果为2+1+1=4种路线
如果(1,0)也有障碍物则不能加上"+1+1"，即只有2种路线
*/

func uniquePathsWithObstacles(og [][]int) int {
	if len(og) == 1 || len(og[0]) == 1 {
		return 1
	}

	m, n := len(og), len(og[0])

	initValue := 1
	base := make([]int, m) 
	// 最右边的一列，如果下面的格子上有障碍物则其上方的格子都无法到达终点
	// 所以从下往上填充1，遇到障碍后填充0
	for k := m - 1; k >= 0; k-- {
		if og[k][n-1] == 1 {
			initValue = 0
		}
		base[k] = initValue
	}
	//fmt.Println("Init value: ", base)

	// 从右数第二列开始向左推进到左数第二列，如果有障碍的格子置0，反之则按0062中思路计算
	for i:=n-2; i>0; i-- {
		if og[m-1][i] == 1 {
			base[m-1] = 0
		}

		for j:=m-2; j>=0; j-- {
			if og[j][i] != 1 {
				base[j] = base[j] + base[j+1]
			}else{
				base[j] = 0
			}
		}	

		//fmt.Println(i, "th round: ", base)
	}

	// 左边的一列参与最后的累加计算，如果有障碍物则其及一下的格子就略过
	rslt := 0	
	for i,v := range base {
		if og[i][0] == 1 {
			break
		}
		rslt += v
	}

	return rslt
}

func main() {
	obstacleGrid1 := [][]int{[]int{0,0,0}, []int{0,1,0}, []int{0,0,0}}
	fmt.Println(uniquePathsWithObstacles(obstacleGrid1))
	obstacleGrid2 := [][]int{[]int{0,1}, []int{0,0}}
	fmt.Println(uniquePathsWithObstacles(obstacleGrid2))
	obstacleGrid3 := [][]int{[]int{0,0,0,0}, []int{0,0,1,0}, []int{0,0,0,0}}
	fmt.Println(uniquePathsWithObstacles(obstacleGrid3))
}
