package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

func minJumps(arr []int) int {
	n, jmpCnt := len(arr), 0
	if n <= 1 {
		return 0
	}

	posDict := make(map[int][]int)
	for i := range arr {
		posDict[arr[i]] = append(posDict[arr[i]], i)
	}

	s := datastruct.NewQueue([]int{})
	s.Push(0)

	for !s.IsEmpty() {
		jmpCnt += 1

		// 非常重要，i的值本身不会用到，但是可以控制for配合queue的FIFO遍历当前level（相当于树的BFS）
		// 换成stack即不成立
		currLvlSize := s.Size()
		for i:=0; i<currLvlSize; i++ {
			idx := s.Pop()

			if v, ok := posDict[arr[idx]]; ok {
				for _,j := range v {
					if idx != j {
						if j == n - 1 {
							return jmpCnt
						}
						s.Push(j)
					}

				}
				delete(posDict, arr[idx])
			}

			if idx - 1 > 0 {
				if _, ok := posDict[arr[idx-1]]; ok {
					s.Push(idx-1)
				}
			}

			if idx + 1 < n {
				if _, ok := posDict[arr[idx+1]]; ok {
					if idx+1 == n-1 {
						return jmpCnt
					}

					s.Push(idx+1)
				}
			}

		}
	}

	return jmpCnt
}

func main() {
	arr := []int{100,-23,-23,404,100,23,23,23,3,404}
	fmt.Println(minJumps(arr))

	arr = []int{7}
	fmt.Println(minJumps(arr))

	arr = []int{7,6,9,6,9,6,9,7}
	fmt.Println(minJumps(arr))

	arr = []int{6,1,9}
	fmt.Println(minJumps(arr))

	arr = []int{11,22,7,7,7,7,7,7,7,22,13}
	fmt.Println(minJumps(arr))
}
