package main

import (
	"fmt"
)

func isNumber(s string) bool {
	// 把合法的字符加入map
	valid := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.'} 	
	dict := make(map[byte]int)

	for _, b := range valid {
		dict[b] = 1
	}

	i := 0

	// 处理开头可能出现的符号
	if s[i] == '-' || s[i] == '+' {
		if len(s) == 1 {
			return false
		}
		i += 1
	}

	// 符号之后第一个字符只可能是0-9和.
	if _, ok := dict[s[i]]; ok {
		// [+|-].和[+|-].e* 应该也是没意义的
		if s[i] == '.' {
			if i+1==len(s) || s[i+1] == 'e' || s[i+1] == 'E' {
				return false
			}
		}

		i += 1

		// 在该字符后e和E也是合法字符，加入hashmap
		dict['e'], dict['E'] = 1, 1
	}else{
		return false
	}

	for i < len(s) {
		if _, ok := dict[s[i]]; ok {
			// 点出现后就不能再出现，从字典中删除
			if s[i] == '.' {
				delete(dict, '.')
			}

			// e和E后面可以有符号；接下来就只能出现0-9了，hashmap中移除e,E和.
			if s[i] == 'e' || s[i] == 'E' {
				if i + 1 == len(s) {
					return false
				}

				if s[i+1] == '+' || s[i+1] == '-' {
					i += 1
				}
				delete(dict, 'e')
				delete(dict, 'E')
				delete(dict, '.')
			}
		}else{
			return false
		}

		i += 1
	}	

	return true    
}

func main() {
	str := []string{"2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789", "abc", "1a", "1e", "e3", ".e3", "99e2.5", "--6", "-+3", "95a54e53", "."}
	for _, s := range str {
		fmt.Printf("%s \tis string ? %v\n", s, isNumber(s))
	}
}
