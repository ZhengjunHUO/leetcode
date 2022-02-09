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
  参考@hxu10的O(len(s))方法
  - n和startPos只是用来计算出上下左右最远可移动的距离，然后就可以忘记
  - 构建一个虚拟的坐标系，出发点为(0,0), 从后向前倒推执行指令s，计算出新的坐标
    (同时(0,0)也是第一个表示出界的终止点，其横纵坐标分别存放于hash表中)
  - 对于这个新的坐标，用开始得到的四个最大距离进行计算，比较计算所得的横/纵坐标是否存在于hash表中
      如果不存在则表示不会触发停止条件，可以走完
      如果存在则表示最远只能走到那个位置
  - 每一轮结束将新坐标加入hash表，成为新的终止点
  * 每一个坐标对本轮来说是用于计算的位置坐标，对于后面的轮来说是一个终止坐标
  * 从正向考虑，当前坐标和四个最值构成了一个矩形面积，如果有终止点正好坐落于边界上，说明按照指令集行动会在该处出界
*/
func executeInstructions(n int, startPos []int, s string) []int {
	// 方向映射表
	direct := map[byte][2]int{'L': [2]int{0, -1}, 'U': [2]int{-1, 0}, 'R': [2]int{0, 1}, 'D': [2]int{1, 0}}
	// 从起始位置出发上下左右最远可以走的距离
	lmax, umax, rmax, dmax := startPos[1] + 1, startPos[0] + 1, n - startPos[1], n - startPos[0]

	fmt.Printf("lmax: %d, umax: %d, rmax: %d, dmax: %d\n\n", lmax, umax, rmax, dmax)
	pathLen := len(s)

	// 设停止点为(0,0)
	row, col := 0, 0
	rowNext, colNext := map[int]int{0: pathLen}, map[int]int{0: pathLen}
	ret := []int{}

	// 从(0,0)开始从后向前倒推每一步的相对位置
	for i := pathLen - 1; i >= 0; i-- {
		row, col = row - direct[s[i]][0], col - direct[s[i]][1]
		num := pathLen - i
		fmt.Printf("row: %d, col: %d, num: %d\n", row, col, num)
		fmt.Printf("colNext: %v, rowNext: %v\n", colNext, rowNext)

		/*
		  检查当前位置出发是否会遇上停止点坐标
		    如没有则表示可以走到底(值为pathLen - i)
		    如遭遇则表示最远只能走到那一点(值取该点位置和当前位置之差); 如两个方向同时match,取其中较近的距离
		*/
		if v, ok := colNext[col-lmax]; ok {
			fmt.Printf("col[%d] - lmax[%d] = %d found in colNext!\n", col, lmax, col - lmax)
			num = min(num, v - i - 1)
		}

		if v, ok := rowNext[row-umax]; ok {
			fmt.Printf("row[%d] - umax[%d] = %d found in rowNext!\n", row, umax, row - umax)
			num = min(num, v - i - 1)
		}

		if v, ok := colNext[col+rmax]; ok {
			fmt.Printf("col[%d] + rmax[%d] = %d found in colNext!\n", col, rmax, col + rmax)
			num = min(num, v - i - 1)
		}

		if v, ok := rowNext[row+dmax]; ok {
			fmt.Printf("row[%d] + dmax[%d] = %d found in rowNext!\n", row, dmax, row + dmax)
			num = min(num, v - i - 1)
		}

		fmt.Printf("num after check: %d\n", num)
		// 更新终止点map
		rowNext[row], colNext[col] = i, i
		ret = append([]int{num}, ret...)
	}

	return ret
}

func main() {
	n := []int{3,2,1}
	startPos := [][]int{[]int{0,1}, []int{1,1}, []int{0,0}}
	s := []string{"RRDDLU","LURD","LRUD"}

	for i := range n {
		fmt.Println(executeInstructions(n[i], startPos[i], s[i]))
		fmt.Println("======")
	}
}
