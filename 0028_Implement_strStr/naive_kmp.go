package main

import (
	"fmt"
	"github.com/ZhengjunHUO/gtoolkit"
)

func strStr(haystack string, needle string) int {
	n, m := len(haystack), len(needle)

	if m == 0 {
		return 0
	}

	tab := gtoolkit.GetKmpTable(needle) 
	i, j := 0, 0	

	for i < n {
		// 如匹配则都向前进1
		if haystack[i] == needle[j] {
			i, j = i+1, j+1
		}

		// pattern已经匹配完成，返回第一次出现的位置(即i-len(pattern))
		if j == m {
			// 如果需要找到所有index的话，这里可以令j = tab[j-1]然后继续匹配
			return i - j
		}

		if i < n && haystack[i] != needle[j] {
			// 部分匹配的情况，此时i不需要移动，j回到tab中记录的上一个元素最长prefix处，进行下一轮比较
			if j > 0 {
				j = tab[j-1]
			}else{
			// 不匹配，i向后移动1，重新和needle从头尝试匹配
				i++
			} 	
		}
	}

	return -1

}

func main() {
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("aaaaa", "bba"))
	fmt.Println(strStr("", ""))
}
