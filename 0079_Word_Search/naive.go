package main

import (
	"fmt"
)

func backtrace(board [][]byte, word string, i int, j int, currIdx int) bool {
	// 表示已经完成匹配
	if currIdx == len(word) {
		//fmt.Println("Got it !")
		return true
	}

	// 如果寻找的坐标超出界限，或者坐标上的值和字符串当前值不匹配则返回false
	// 后者可以快速略过board中一开始就不匹配的情况，大大减少需要深入寻找的工作量
	if i < 0 || i > len(board)-1 || j < 0 || j > len(board[i])-1 || board[i][j] != word[currIdx] {
		return false
	}

	//fmt.Println(i, j)
	// 用byte的max value来暂时遮挡当前值
	board[i][j] = 255
	// 向四个方向递归寻找(dfs)从当前元素开始是否可能有匹配值
	rslt := backtrace(board, word, i - 1, j, currIdx + 1) || backtrace(board, word, i + 1, j, currIdx + 1) || backtrace(board, word, i, j - 1, currIdx + 1) || backtrace(board, word, i, j + 1, currIdx + 1 )
	//还原当前值
	board[i][j] = word[currIdx]
	return rslt
}

func exist(board [][]byte, word string) bool {
	for i := range board {
		for j := range board[i] {
			// 遍历矩阵，尝试从每个格子出发寻找是否有匹配word的存在
			// 理论上时间复杂度达到O(i*j*4^len(word)), 但是很多和word首字母不同的元素会被很快略过
			// 所以实际花费时间还是在可以接受范围
			if backtrace(board, word, i, j, 0) {
				return true
			}
		}
	}

	return false 
}

func main() {
	board := [][]byte{[]byte{'A', 'B', 'C', 'E'}, []byte{'S', 'F', 'C', 'S'}, []byte{'A', 'D', 'E', 'E'}}
	str := []string{"ABCCED", "SEE", "ABCB"}

	for i := range str {
		fmt.Println(exist(board, str[i]))
	}
}
