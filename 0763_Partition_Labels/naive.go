package main

import (
	"fmt"
	"sort"
)

func partitionLabels(s string) []int {
	rslt, rsltlen := [][]int{}, []int{}
	interval := make(map[byte][]int)

	// 统计每个字母出现的始末index
	for i := range s {
		if _, ok := interval[s[i]]; !ok {
			interval[s[i]] = []int{i, i}
		}else{
			interval[s[i]][1] = i
		}
	}

	list := make([][]int, 0, len(interval))
	for _, v := range interval {
		list = append(list, v)
	}

	// 将区间按照起始index排序，然后合并重叠区间
	sort.SliceStable(list, func(i, j int) bool { return list[i][0] < list[j][0] })	
	rslt = append(rslt, list[0])	
	for i := 1; i < len(list); i++ {
		if rslt[len(rslt)-1][1] < list[i][0] {
			rslt = append(rslt, list[i])
		}else{
			if rslt[len(rslt)-1][1] < list[i][1] {
				rslt[len(rslt)-1][1] = list[i][1] 
			}
		}
	}

	for i := range rslt {
		rsltlen = append(rsltlen, rslt[i][1] - rslt[i][0] + 1)
	}

	return rsltlen    
}

func main() {
	str := []string{"ababcbacadefegdehijhklij", "eccbbbbdec"}
	
	for i := range str {
		fmt.Println(partitionLabels(str[i]))
	}
}
