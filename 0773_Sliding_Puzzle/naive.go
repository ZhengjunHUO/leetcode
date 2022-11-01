package main

import (
	"fmt"
	"strings"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

func slidingPuzzle(board [][]int) int {
	// 需要到达的状态
	goal := "123450"

	// 初始状态
	str := ""
	for i := range board {
		str += strings.Trim(strings.Join(strings.Split(fmt.Sprint(board[i]), " "), ""), "[]")
	}

	// 每个格子可以和周围交换的格子列表
	neighbour := [][]int{[]int{1,3}, []int{0,2,4}, []int{1,5}, []int{0,4}, []int{1,3,5}, []int{2,4}}

	// 记录已经出现过的状态，避免重复计算
	dict := make(map[string]bool)
	dict[str] = true

	// BFS标配
	queue := datastruct.NewQueue([]string{})
	queue.Push(str)

	// 记录步数
	rslt := 0

	for !queue.IsEmpty() {
		size := queue.Size()
		for i := 0; i < size; i++ {
			curr := queue.Pop()
			if curr == goal {
				return rslt
			}

			// 寻找当前状态0的位置
			temp, idx := []byte(curr), 0
			for i := range temp {
				if temp[i] == byte('0') {
					idx = i
					break
				}
			}

			// 通过查表遍历交换
			for _,v := range neighbour[idx] {
				temp[idx], temp[v] = temp[v], temp[idx]
				tempstr := string(temp)
				// 跳过已经访问过的，记得还原temp
				if _, ok := dict[tempstr]; ok {
					temp[idx], temp[v] = temp[v], temp[idx]
					continue
				}
				// push新状态到队列，标记为见过
				queue.Push(tempstr)
				dict[tempstr] = true
				// 还原
				temp[idx], temp[v] = temp[v], temp[idx]
			}
		}

		// 完成一层，尚未找到，步数加1
		rslt++
	}

	return -1
}

func main() {
	bs := [][][]int{[][]int{[]int{1,2,3},[]int{4,0,5}}, [][]int{[]int{1,2,3},[]int{5,4,0}}, [][]int{[]int{4,1,2},[]int{5,0,3}}, [][]int{[]int{3,2,4},[]int{1,5,0}}}
	for i := range bs {
		fmt.Println(slidingPuzzle(bs[i]))
	}
}
