package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	rslt := strs[0]

	for i:=1; i<len(strs); i++ {
		for j:=0; j<len(rslt) && j<len(strs[i]); j++ {
			if strs[i][j] != rslt[j] {
				if j == 0 {
					return ""
				}
				rslt = rslt[:j]
				break
			}
		}
	}

	return rslt
}

func main() {
	str := [][]string{[]string{"flower","flow","flight"}, []string{"dog","racecar","car"}}

	for _, v := range str {
		fmt.Println(longestCommonPrefix(v))
	}
}
