package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

const MaxInt = int(^uint(0) >> 1)

func networkDelayTime(times [][]int, n int, k int) int {
        // 1. 生成字典形式的邻接表，value为[邻接结点编号，花费]
        dict := make(map[int][][]int)
        for i := range times {
                dict[times[i][0]] = append(dict[times[i][0]], []int{times[i][1], times[i][2]})
        }

	// 2. 初始化costTo
	costTo := make([]int, n+1)
	for i := range costTo {
		costTo[i] = MaxInt
	}
	costTo[k] = 0

	// 3. 准备优先队列，值为节点ID，优先级为出发点到该节点的花费，优先Pop出花费小的
	pq := datastruct.NewPQ([]int{k}, []int{0}, true)

	for pq.Size() > 0 {
                idCurr, costToCurr := pq.PopWithPriority()

		if costToCurr > costTo[idCurr] {
			continue
		}

		if v, ok := dict[idCurr]; ok {
			for i := range v {
				if temp := costToCurr + v[i][1]; temp < costTo[v[i][0]] {
					costTo[v[i][0]] = temp
					pq.Push(v[i][0], temp)
				}
			}
		}
	}

	rslt := -1
	for i := 1; i < n+1; i++ {
		if costTo[i] == MaxInt {
			return rslt
		}

		if costTo[i] > rslt {
			rslt = costTo[i]
		}
	}

	return rslt
}

func main() {
	times := [][][]int{[][]int{[]int{2,1,1},[]int{2,3,1},[]int{3,4,1}}, [][]int{[]int{1,2,1}}, [][]int{[]int{1,2,1}}}
	size := []int{4,2,2}
	start := []int{2,1,2}

	for i := range times {
		fmt.Println(networkDelayTime(times[i], size[i], start[i]))
	}
}
