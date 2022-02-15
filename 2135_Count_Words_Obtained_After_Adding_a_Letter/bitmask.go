package main

import (
	"fmt"
)

func wordCount(startWords []string, targetWords []string) int {
	var ret int
	dict := make(map[int]struct{})

	// 由于题目规定没有重复字符，所以对字符串中出现的字符在a-z对应0-25的位置置1
	for i := range startWords {
		w := 0
		for j := range startWords[i] {
			w ^= 1 << (int(startWords[i][j]) - 97)
		}
		// 存放于集合中
		dict[w] = struct{}{}
	}

	for i := range targetWords {
		// 同样的办法分析目标单词
		w := 0
		for j := range targetWords[i] {
			w ^= 1 << (int(targetWords[i][j]) - 97)
		}

		// 目标单词减少一个字符，看是否能匹配初始单词
		for j := range targetWords[i] {
			if _, ok := dict[w ^ (1 << (int(targetWords[i][j]) - 97))]; ok {
				//fmt.Printf("targetWords[%d] ok\n", i)
				ret++
			}
		}
	}

	return ret
}

func main() {
	startWords := [][]string{[]string{"ant","act","tack"}, []string{"ab","a"}}
	targetWords := [][]string{[]string{"tack","act","acti"}, []string{"abc","abcd"}}

	for i := range startWords {
		fmt.Println(wordCount(startWords[i], targetWords[i]))
	}
}
