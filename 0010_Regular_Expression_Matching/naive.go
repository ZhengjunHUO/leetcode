package main

import (
	"fmt"
)

func isMatch(s string, p string) bool {
	m, n, i, j := len(s), len(p), 0, 0
	var curr byte

	for i < n {
		// s已经匹配完成
		if j == m {
			return true
		}

		// 存放pattern中的当前字符
		curr = p[i]

		// 已经到了pattern最后一个字符
		if i + 1 == n {
			// s的当前字符不匹配pattern最后一个字符
			if curr != s[j] && curr != '.' {
	 			return false
			// 匹配，如果s当前正好也是最后一个字符的话最终返回true
			}else{
				j+=1
				break
			}
		}  	

		// 已知pattern当前字符不是最后一个字符，可以向后看一个字符
		switch p[i+1]{
		case '*':
			// 检查s是否有等于curr的字符，一直前进至不匹配为止
			// 如果一上来就不匹配的话就不用前进
			for j < m {
				if curr != s[j] && curr != '.' {
					break
				}
				j += 1
			}
			i += 2
		default:
			if curr != s[j] && curr != '.' {
				return false
			}
			i, j = i+1, j+1
		}
	}    

	return j == m
}

func main() {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("ab", ".*"))
	fmt.Println(isMatch("aab", "c*a*b"))
	fmt.Println(isMatch("mississippi", "mis*is*p*."))
	fmt.Println(isMatch("mississippi", "mis*is*ip*."))
}
