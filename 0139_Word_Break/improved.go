package main

import (
	"fmt"
	zstr "github.com/ZhengjunHUO/goutil/strings"
	"sort"
)

func wordBreak(s string, wordDict []string) bool {
	// 对wordDict中所有pattern计算LPS表
	pf := make([]*zstr.PatternFinder, len(wordDict))
	for i := range wordDict {
		pf[i] = zstr.NewPatternFinder(wordDict[i])
	}

	// 对每个pattern使用KMP算法找出所有匹配
	rgs := [][2]int{}
	for i := range pf {
		if rslt := pf[i].FindIn(s); len(rslt) > 0 {
			for j := range rslt {
				rgs = append(rgs, [2]int{rslt[j], rslt[j] + len(pf[i].Pattern) - 1})
			}
		}

	}

	// 对结果升序排序
	sort.SliceStable(rgs, func(i, j int) bool{
		if rgs[i][0] == rgs[j][0] {
			return rgs[i][1] < rgs[j][1]
		}

		return rgs[i][0] < rgs[j][0]
	})

	//fmt.Println(rgs)

	// 存放使用pattern可以拼凑出的子字符串的起始index, 在开始padding一个空格
	table := make([]int, len(s)+1)
	for i := range table {
		table[i] = -1
	}

	for i := range rgs {
		// 观察起始index前一个格子的值，如果为-1则不能接上之前的子字符串
		if table[rgs[i][0]] == -1 {
			table[rgs[i][1]+1] = rgs[i][0]+1
		// 如果不为-1则表示可以接上子字符串
		}else{
			table[rgs[i][1]+1] = table[rgs[i][0]]
		}
	}

	//fmt.Println(table)
	return table[len(table)-1] == 1
}

func main() {
	s := []string{"leetcode", "applepenapple", "catsandog"}
	wordDict := [][]string{[]string{"leet","code"}, []string{"apple","pen"}, []string{"cats","dog","sand","and","cat"}}
	for i := range s {
		fmt.Println(wordBreak(s[i], wordDict[i]))
	}
}
