package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	pq := datastruct.NewPQ([]int{}, []int{}, false)

	for k > 0 {
		i := 0
		for i < len(capital) {
			// 将成本小于等于目前资本的项目加入PQ
			if capital[i] <= w {
				pq.Push(profits[i], profits[i])

				// 将该项目从两个列表中移除
				copy(profits[i:], profits[i+1:])
				profits = profits[:len(profits)-1]
				copy(capital[i:], capital[i+1:])
				capital = capital[:len(capital)-1]
				i--
			}
			i++
		}

		if pq.Size() == 0 {
			return w
		}else{
			// 将当前可承受的最赚钱项目的利润加入资本
			w += pq.Pop()
			k--
		}
	}

	return w
}

func main() {
	k, w := 2, 0
	profits := []int{1, 2, 3}
	capital := []int{0, 1, 1}
	fmt.Println(findMaximizedCapital(k, w, profits, capital))

	k, w = 3, 0
	profits = []int{1, 2, 3}
	capital = []int{0, 1, 2}
	fmt.Println(findMaximizedCapital(k, w, profits, capital))
}
