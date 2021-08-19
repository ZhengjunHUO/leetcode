package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	Elems	[]int    
}

func Constructor(n int, blacklist []int) Solution {
	newsize := n - len(blacklist)

	rslt := make([]int, n)
	for i := range rslt {
		rslt[i] = i
	}

	// 将黑名单加入hashmap便于查询
	dict := make(map[int]bool)
	for i := range blacklist {
		dict[blacklist[i]] = true
	}

	// 将所有黑名单中的元素swap到队伍末
	last := len(rslt) - 1
	for i := range blacklist {
		// 跳过已经在队伍后部无需swap的黑名单元素
		if blacklist[i] >= newsize {
			continue	
		}

		for {
			// 避免原来就在后部的黑名单元素被swap到前面
			if _, ok := dict[rslt[last]]; ok {
				delete(dict, rslt[last])
				last--
			}else{
				break
			}
		}
		
		rslt[blacklist[i]], rslt[last] = rslt[last], rslt[blacklist[i]]	
		delete(dict, blacklist[i])
		last--
	}

	// 只返回数列前部
	return Solution{rslt[:newsize]}
}

func (this *Solution) Pick() int {
	return this.Elems[rand.Int()%len(this.Elems)]
}

func main() {
	obj := Constructor(7, []int{2, 3, 5})
	for i:=0; i<12; i++ {
		fmt.Println(obj.Pick())
	}
}
