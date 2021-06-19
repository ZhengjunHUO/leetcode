package main

import (
	"fmt"
)

/* 
  本方案支持中文
*/

func reverse(s string) string {
        r := []rune(s)

        for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
                r[i], r[j] = r[j], r[i]
        } 

        return string(r)
}

// 遍历字符串，记录从index 0到每个index i为止，最长的即是前缀也是后缀的pattern的长度
func buildKMPTable(s string) []int {
	r := []rune(s)
	n := len(r)

	tab := make([]int, n)
	tab[0] = 0
	
	for i:=1; i<n; i++ {
		tmp := tab[i-1]

		for tmp > 0 && r[i] != r[tmp] {
			tmp = tab[tmp-1]
		}
		if r[i] == r[tmp] {
			tmp++
		}

		tab[i] = tmp
	}

	return tab
}

/*
  将原字符串以"#"符号为分隔做成镜像，使用kmp算法找到从字符串头开始最长的回文组
  把原字符串去掉该回文组后剩下的部分倒序接在原字符串开头即可
*/
func shortestPalindrome(s string) string {
	symmetric := s + "#" + reverse(s)
	tab := buildKMPTable(symmetric)
	return reverse(string([]rune(s)[tab[len(tab)-1]:])) + s
}

func main() {
	fmt.Println(shortestPalindrome("aacecaaa"))
	fmt.Println(shortestPalindrome("abcd"))
	fmt.Println(shortestPalindrome("jkbacabdf"))
	fmt.Println(shortestPalindrome("为我我为人人"))
}
