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

// 空间复杂度为O(kn)：函数个数为k*n，计算所得的状态需要存储
// 时间复杂度为O(kn*logn)：每个函数中都有一个二分搜索
func eggDrop(k int, n int) int {
	if k == 1 {
		return n
	}
	if n == 0 {
		return 0
	}

	if v, ok := dict[[2]int{k,n}]; ok {
		return v
	}

	rslt := 10001
	l, r := 1, n
	// 找到需要最少操作数的楼层
	for l <= r {
		m := l + (r-l)/2
		// 如果鸡蛋碎了则往i下方寻找(k-1, i-1) <= 随着i的增加呈现递增关系
		brk := eggDrop(k-1, m-1) 
		// 反之则向上寻找(k, n-i) <= 随着i的增加呈现递减关系
		nbk := eggDrop(k, n-m) 
		// 需要找到两个值相等的那一点
		// 参考暴力破解中的rslt = min(rslt, max(eggDrop(k, n-i), eggDrop(k-1, i-1))+1)
		if brk > nbk {		
			// 递增函数的值大于递减函数的值，需要向左搜索
			r = m - 1	
			rslt = min(rslt, brk+1)
		}else{
			l = m + 1
			rslt = min(rslt, nbk+1)
		}
	}

	dict[[2]int{k,n}] = rslt
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
