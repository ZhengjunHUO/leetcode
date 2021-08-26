package main

import (
	"fmt"
)

var dict map[[2]int]int

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// 空间复杂度为O(kn)：函数个数为k*n，计算所得的状态需要存储
// 时间复杂度为O(kn^2)：每个函数中都有一个loop
func eggDrop(k int, n int) int {
	//fmt.Printf("Inside k: %d; n: %d\n", k, n)
	// 只有一个鸡蛋，只能从底楼往上逐层尝试，最坏的情况为n
	if k == 1 {
		return n
	}
	// 在底楼没必要尝试
	if n == 0 {
		return 0
	}

	if v, ok := dict[[2]int{k,n}]; ok {
		//fmt.Println("Hit cache, return ", v)
		return v
	}

	rslt := 10001
	for i:=1; i<=n; i++ {
		//fmt.Printf("compare eggDrop(%d, %d) and eggDrop(%d, %d)\n", k, n-i, k-1, i-1)
		// 找到需要最少操作数的楼层
		// 每次行动有两种可能结果，如果鸡蛋碎了则往i下方寻找(k-1, i-1)；反之则向上寻找(k, n-i)
		rslt = min(rslt, max(eggDrop(k, n-i), eggDrop(k-1, i-1))+1)
	}

	dict[[2]int{k,n}] = rslt
	//fmt.Printf("Leave k: %d; n: %d, return: %d\n", k, n, rslt)
	return rslt
}

func superEggDrop(k int, n int) int {
	dict = make(map[[2]int]int)
	/*
	defer func() {
		fmt.Println(dict)
	}()
	*/
	return eggDrop(k, n)
}

func main() {
	ks := []int{1, 2, 3}
	ns := []int{2, 6, 14}

	for i := range ks {
		fmt.Println(superEggDrop(ks[i], ns[i]))
	}
}
