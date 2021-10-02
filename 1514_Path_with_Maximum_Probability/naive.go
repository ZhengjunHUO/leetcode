package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	// 构建邻接表
	dict := make(map[int][][2]interface{})
	for i := range edges {
		dict[edges[i][0]] = append(dict[edges[i][0]], [2]interface{}{edges[i][1], succProb[i]})
		dict[edges[i][1]] = append(dict[edges[i][1]], [2]interface{}{edges[i][0], succProb[i]})
	}

	// 准备概率表
	probTo := make([]float64, n)
	probTo[start] = 1.0

	// 优先队列
	pq := godtype.NewPQ([]int{start}, []float64{1.0}, false)
	for pq.Size() > 0 {
		curr := pq.PopWithPrio()
		currNode, probToCurr := curr[0].(int), curr[1].(float64)

		if probToCurr < probTo[currNode] {
			continue
		}

		if v, ok := dict[currNode]; ok {
			for i := range v {
				if temp := probToCurr * v[i][1].(float64); temp > probTo[v[i][0].(int)] {
					probTo[v[i][0].(int)] = temp
					pq.Push(v[i][0].(int), temp)
				}
			}
		}
	}

	return probTo[end]
}

func main() {
	edges := [][][]int{[][]int{[]int{0,1},[]int{1,2},[]int{0,2}}, [][]int{[]int{0,1},[]int{1,2},[]int{0,2}}, [][]int{[]int{0,1}}}
	succProb := [][]float64{[]float64{0.5,0.5,0.2}, []float64{0.5,0.5,0.3}, []float64{0.5}}

	for i := range edges {
		fmt.Println(maxProbability(3, edges[i], succProb[i], 0, 2))
	}
}
