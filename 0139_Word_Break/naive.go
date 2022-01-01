package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func wordBreak(s string, wordDict []string) bool {
	// 使用KMP算法，对wordDict中所有pattern计算LPS表
	pf := make([]*godtype.PatternFinder, len(wordDict))
	for i := range wordDict {
		pf[i] = godtype.NewPatternFinder(wordDict[i]) 	
	}

	// 存放使用pattern可以拼凑出的子字符串的起始index, 在开始padding一个空格
	table := make([]int, len(s)+1)
	for i := range table {
		table[i] = -1
	}

	for i := range s {
		// 随着slice的增长使用kmp算法更新table, O(n^2)存在很多重复计算待优化
		for j := range pf {
			if rslt := pf[j].FindIn(s[:i+1]); len(rslt) > 0 {
				// 如果pattern存在于当前slice，记录起点终点
				beginAt := rslt[len(rslt)-1]
				endAt := beginAt + len(pf[j].Pattern()) - 1
				//fmt.Printf("%v find in %s at %d to %d\n",string(pf[j].Pattern()), s[:i+1], beginAt, endAt)
				// 观察起始index前一个格子的值，如果为-1则不能接上之前的子字符串
				if table[beginAt] == -1 {
					table[endAt+1] = beginAt+1
				}else{
				// 如果不为-1则表示可以接上子字符串
					table[endAt+1] = table[beginAt]
				}
			}
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
