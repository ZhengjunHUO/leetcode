package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

/*
参考了@gd303的思路，将最大多少跳的问题转化成有向无环图的最长路径问题
限于自身的能力和golang的机制代码有些重复，有待改进
*/
func maxJumps(arr []int, d int) int {
	n := len(arr)
	// 有向无环图
	dag := make(map[int][]int)
	max := 1

	s := datastruct.NewStack([]int{})
	for i := range arr {
		// 当前元素是大于栈顶元素的最近的值
		for !s.IsEmpty() && arr[s.Peek()] < arr[i] {
			temp := s.Pop()
			// 如在范围d内则加入有向图
			if i - temp <= d {
				dag[temp] = append(dag[temp], i)
			}
		}
		s.Push(i)
	}

	/* 
	从反方向重新执行一遍上述操作，最后dag中存储的是
	元素i左右两个方向“在范围d内”可以向上爬到的“最近的”高地
	*/
	s = datastruct.NewStack([]int{})
	for i := n-1 ; i>=0 ; i-- {
		for !s.IsEmpty() && arr[s.Peek()] < arr[i] {
			temp := s.Pop()
			if temp - i <= d {
				dag[temp] = append(dag[temp], i)
			}
		}
		s.Push(i)
	}

	pathlen := make(map[int]int)
	// 通过dag寻找从各元素出发可以走到的最长路径，记录最大值
	for i := range arr {
		if temp := findPathLen(dag, i, pathlen); temp > max {
			max = temp
		}else{
		}
	}

	return max
}

func findPathLen(dag map[int][]int, i int, pathlen map[int]int) int {
	// 使用cache，避免重复计算
	if v, ok := pathlen[i]; ok {
		return v
	}

	switch n := len(dag[i]); n {
	// 说明该值是局部max，从它出发在d以内没有更大的值
	case 0:
		pathlen[i] = 1
	case 1:
		rslt := 1+findPathLen(dag, dag[i][0], pathlen)
		pathlen[i] = rslt
	case 2:
		temp1, temp2 := findPathLen(dag, dag[i][0], pathlen), findPathLen(dag, dag[i][1], pathlen)
		if temp1 > temp2 {
			pathlen[i] = 1+temp1
		}else{
			pathlen[i] = 1+temp2
		}
	}

	return pathlen[i]
}


func main() {
	arr := []int{6,4,14,6,8,13,9,7,10,6,12}
	d := 2
	fmt.Println(maxJumps(arr, d))

	arr = []int{3,3,3,3,3}
	d = 3
	fmt.Println(maxJumps(arr, d))

	arr = []int{7,6,5,4,3,2,1}
	d = 1
	fmt.Println(maxJumps(arr, d))

	arr = []int{7,1,7,1,7,1}
	d = 2
	fmt.Println(maxJumps(arr, d))

	arr = []int{66}
	d = 1
	fmt.Println(maxJumps(arr, d))
}
