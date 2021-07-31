package main

import (
	"fmt"
)

func partitionLabels(s string) []int {
	// 只记录每个字母最后出现的index
	dict := make(map[byte]int)
	for i := range s {
		dict[s[i]] = i
	}

	// 当前分片的起始位置
	begin, end, rslt := 0, 0, []int{}
	for i := range s {
		if dict[s[i]] > end {
			end = dict[s[i]]	
		}

		// 相当于merge完了从begin开始的重叠的interval
		if i == end {
			rslt = append(rslt, end - begin + 1)
			begin = end + 1
		}
	}

	return rslt	
}

func main() {
	str := []string{"ababcbacadefegdehijhklij", "eccbbbbdec"}
	
	for i := range str {
		fmt.Println(partitionLabels(str[i]))
	}
}
