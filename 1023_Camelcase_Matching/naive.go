package main

import (
	"fmt"
)

func camelMatch(queries []string, pattern string) []bool {
	rslt := make([]bool, len(queries))
	
	// 对每个单词和pattern进行匹配
	for i := range queries {
		ptr_p, ptr_q := 0, 0
		// 遍历单词中的字母
		loop: for ptr_q < len(queries[i]) {
			// 如果匹配pattern中的字母，两个指针都前移
			if queries[i][ptr_q] == pattern[ptr_p] {
				ptr_q += 1
				ptr_p += 1
			}else{
				// 如果不匹配，但是单词中字母为大写字母，则匹配失败
				if int(queries[i][ptr_q]) < 91 {
					break loop
				}
				ptr_q += 1
			}

			// 如果pattern已经匹配完成，只需要检查单词剩下的字母中没有大写字母即可
			if ptr_p == len(pattern) {
				for ptr_q < len(queries[i]) {
					if int(queries[i][ptr_q]) < 91 {
						break loop				
					}
					ptr_q += 1
				}
				rslt[i] = true
			}
		}
	}

	return rslt
}

func main() {
	queries := [][]string{{"FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"}, {"FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"}, {"FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"}}
	pattern := []string{"FB", "FoBa", "FoBaT"}

	for i := range pattern {
		fmt.Println(camelMatch(queries[i], pattern[i]))
	}
}
