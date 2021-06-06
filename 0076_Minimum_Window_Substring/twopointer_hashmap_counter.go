package main

import (
	"fmt"
)

// char in t is unique 
func minWindow(s string, t string) string {
	// 需要找到的各元素，及对应的次数（值减为0=》该元素已找齐）
	elemCounter := make(map[int]int)
	for i := range t {
		elemCounter[int(t[i])] += 1
	}

	// 需要找到的元素个数（值减为0=> valid状态）
	counter := len(t)
	
	// 左右指针
	lp, rp := 0,0
	// 存储结果
	startPos, minSize := -1, 100000

	// 右指针遍历到底
	for (rp<len(s)) {
		//右指针行动
		if elemCounter[int(s[rp])] > 0 {
			counter -= 1
		}
		elemCounter[int(s[rp])] -= 1
		rp += 1

		for (counter == 0) {
			//更新较优解
			if (rp-lp) < minSize {
				startPos = lp
				minSize = rp-lp
			}

			//左指针行动（仅在valid状态下移动）
			elemCounter[int(s[lp])] += 1
			if elemCounter[int(s[lp])] > 0 {
				counter += 1
			}
			lp += 1
		}
	}

	if minSize < 100000 {
		return s[startPos:startPos+minSize]
	}

	return ""
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s,t))
}
