package main

import (
	"fmt"
)

func abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}

func containsNearbyDuplicate(nums []int, k int) bool {
	// 值为value出现的index列表
	dict := make(map[int][]int)

	for i, v := range nums {
		if list, ok := dict[v]; ok {
			// 值已经存在 && 比较是否满足abs(delta) <= k的条件
			for j := range list {
				if abs(list[j] - i) <= k {
					return true
				}  
			}

			// 不要忘了将index插入列表
			dict[v] = append(dict[v], i)
		}else{
			// 如果值不存在于字典中新建slice
			dict[v] = []int{i}
		}
	}

	return false
}

func main() {
	nums := [][]int{[]int{1,2,3,1}, []int{1,0,1,1}, []int{1,2,3,1,2,3}}
	k := []int{3, 1, 2}

	for i := range nums {
		fmt.Println(containsNearbyDuplicate(nums[i], k[i]))
	}
}
