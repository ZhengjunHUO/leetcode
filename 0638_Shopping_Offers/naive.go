package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// 将needs []int转换为map可以接受的键[3]int
func sliceToTriplet(s []int) [3]int {
	triplet := [3]int{}
	if len(s) > 3 {
		fmt.Println("[WARN] Can't conver slice %v to triplet", s)
		return triplet
	}

	for i := range s {
		triplet[i] = s[i]
	}

	return triplet
}

func dfs(price []int, special [][]int, needs []int, memo map[[3]int]int) int {
	if v, ok := memo[sliceToTriplet(needs)]; ok {
		return v
	}

	var ret int
	// (1) 不考虑special offer需要花费的金额
	for i := range needs {
		ret += price[i] * needs[i]
	}
	//fmt.Printf("%v's cost without offer: %d\n", needs, ret)

	// (2) 对当前needs，判断每一个special offer判断是否可以应用
	needs_updated := make([]int, len(needs))
	for _,v := range special {
		
		for i := range needs {
			if needs[i] < v[i] {
				// 不满足条件，检查下一个offer
				goto toNextOffer
			}
		}

		// offer可用，needs中减去offer中的商品数
		for i := range needs {
			needs_updated[i] = needs[i] - v[i]
		}
		//fmt.Printf("Compare %d and %d+dfs(...%v...)\n", ret, v[len(v)-1], needs_updated)
		// 比较不采用offer和采用的花费，取较小者
		ret = min(ret, v[len(v)-1] + dfs(price, special, needs_updated, memo))

		toNextOffer:
	}

	// 保存结果
	memo[sliceToTriplet(needs)] = ret

	return ret
}

func shoppingOffers(price []int, special [][]int, needs []int) int {
	// 用来存储已经计算过的needs对应的最小cost
	memo := make(map[[3]int]int)
	return dfs(price, special, needs, memo)
}

func main() {
	price := [][]int{[]int{2,5}, []int{2,3,4}}
	special := [][][]int{[][]int{[]int{3,0,5},[]int{1,2,10}}, [][]int{[]int{1,1,0,4},[]int{2,2,1,9}}}
	needs := [][]int{[]int{3,2}, []int{1,2,1}}

	for i := range price {
		fmt.Println(shoppingOffers(price[i], special[i], needs[i]))
	}
}
