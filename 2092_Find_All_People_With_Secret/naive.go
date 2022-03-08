package main

import (
	"fmt"
	"sort"
	"github.com/ZhengjunHUO/godtype"
)

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	// 按照meetings的时间发生顺序从低到高排序
	sort.SliceStable(meetings, func (i, j int) bool {
		return meetings[i][2] < meetings[j][2]
	})

	// 创建一个unionfind对象，连接0和firstPerson表示秘密被分享
	uf := godtype.NewUF(n)
	uf.Union(0, firstPerson)

	i := 0
	for i < len(meetings) {
		dict := make(map[int]struct{})	
		time := meetings[i][2]

		// 对于同一时刻发生的meetings，连接会议的两方，记录于一个set中
		for i < len(meetings) && meetings[i][2] == time {
			uf.Union(meetings[i][0], meetings[i][1])
			dict[meetings[i][0]], dict[meetings[i][1]] = struct{}{}, struct{}{}
			i++
		}

		// 移除该时刻中没有获得秘密的节点的连接
		for k,_ := range dict {
			if !uf.IsLinked(0, k) {
				// 该操作会使uf中的size和count失效，但是本题并未用到所以没有影响
				p := uf.Parent()
				p[k] = k
				uf.SetParent(p)
			}
		}
	}

	ret := []int{0}
	// 选出和0连接的节点
	for j:=1; j<n; j++ {
		if uf.IsLinked(0, j) {
			ret = append(ret, j)
		}
	}

	return ret
}

func main() {
	n := []int{6,4,5}
	meetings := [][][]int{[][]int{[]int{1,2,5},[]int{2,3,8},[]int{1,5,10}},[][]int{[]int{3,1,3},[]int{1,2,2},[]int{0,3,3}},[][]int{[]int{3,4,2},[]int{1,2,1},[]int{2,3,1}}}
	firstPerson := []int{1,3,1}

	for i := range n {
		fmt.Println(findAllPeople(n[i], meetings[i], firstPerson[i]))
	}
}
