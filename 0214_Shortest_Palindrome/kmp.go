package main

import (
	"fmt"
	"github.com/ZhengjunHUO/gtoolkit"
)

/*
  将原字符串以"#"符号为分隔做成镜像，使用kmp算法找到从字符串头开始最长的回文组
  把原字符串去掉该回文组后剩下的部分倒序接在原字符串开头即可
  本方案支持中文
*/
func shortestPalindrome(s string) string {
	symmetric := s + "#" + gtoolkit.ReverseStr(s)
	tab := gtoolkit.GetKmpTable(symmetric)
	return gtoolkit.ReverseStr(string([]rune(s)[tab[len(tab)-1]:])) + s
}

func main() {
	fmt.Println(shortestPalindrome("aacecaaa"))
	fmt.Println(shortestPalindrome("abcd"))
	fmt.Println(shortestPalindrome("jkbacabdf"))
	fmt.Println(shortestPalindrome("为我我为人人"))
}
