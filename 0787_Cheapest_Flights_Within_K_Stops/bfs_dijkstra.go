package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)


func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	/* 制作邻接表
	tab := make([][]int, n)
	for i := range tab {
		tab[i] = make([]int, n)
	}
	for i := range flights {
		tab[flights[i][0]][flights[i][1]] = flights[i][2]
	}
	*/

	// 字典形式的邻接表，value为[邻接结点编号，花费]
	dict := make(map[int][][]int)
	for i := range flights {
		dict[flights[i][0]] = append(dict[flights[i][0]], []int{flights[i][1], flights[i][2]})
	}

	// 值为[结点编号，剩余可中转数]，权重为累计花费
	pq := godtype.NewPQ([][]int{[]int{src, k+1}}, []int{0}, true) 
	for pq.Size() > 0 {
		// 优先pop出累计花费最少的点
		curr := pq.PopWithPrio()
		currPos, currK, currCost := (curr[0].([]int))[0], (curr[0].([]int))[1], curr[1].(int)
		if currPos == dst {
			return currCost
		}
		
		if currK > 0 {
			/*
			for i := range tab[currPos] {
				if tab[currPos][i] != 0 {
					pq.Push([]int{i, currK-1}, currCost + tab[currPos][i])
				}
			}
			*/
			if v, ok := dict[currPos]; ok {
				for i := range v {
					pq.Push([]int{v[i][0], currK-1}, currCost + v[i][1])
				}
			}
		}
	}
	
	return -1
}

func main() {
	f := [][]int{[]int{0,1,100},[]int{1,2,100},[]int{0,2,500}}
	fmt.Println(findCheapestPrice(3, f, 0, 2, 1))
	fmt.Println(findCheapestPrice(3, f, 0, 2, 0))
}
