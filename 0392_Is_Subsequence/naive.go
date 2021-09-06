package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	l, r := 0, len(nums)
	
	for l < r {
		m := l + (r - l)/2
		if nums[m] >= target {
			r = m
		}else{
			l = m + 1
		}
	}

	return l
}

func isSubsequence(s string, t string) bool {
	// 分析t
	dict := make(map[byte][]int)
	for i := range t {
		if _, ok := dict[t[i]]; ok {
			dict[t[i]] = append(dict[t[i]], i)
		}else{
			dict[t[i]] = []int{i}
		}
	}
	
	// 指向t的指针，在t中找到和s中的字符最靠前匹配的index
	j := 0
	for i := range s {
		if _, ok := dict[s[i]]; ok {
			k := search(dict[s[i]], j)
			// 匹配失败
			if k == len(dict[s[i]]) {
				return false
			}
			// 更新j
			j = dict[s[i]][k] + 1
		}else{
			// s中字符在t中不存在
			return false
		}
	}

	return true
}

func main() {
	ss := []string{"abc", "axc"}
	t := "ahbgdc"

	for i := range ss {
		fmt.Println(isSubsequence(ss[i], t))
	}
}
