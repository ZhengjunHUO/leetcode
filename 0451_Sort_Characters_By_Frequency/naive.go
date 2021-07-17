package main

import (
	"fmt"
)

func frequencySort(s string) string {
	dict := make(map[byte]int)
	for i := range s {
		dict[s[i]] += 1
	}

	bucket := make([][]byte, len(s))
	for k, v := range dict {
		if bucket[v] == nil {
			bucket[v] = []byte{k}
		}else{
			bucket[v] = append(bucket[v], k)
		}
	}

	rslt := make([]byte, 0, len(s))
	for i := len(bucket)-1; i>=0; i-- {
		if bucket[i] != nil {
			for _, v := range bucket[i] {
				for j := 0; j < i; j++ {
					rslt = append(rslt, v)
				}
			}
		}
	}

	return string(rslt)
}

func main() {
	strs := []string{"tree", "cccaaa", "Aabb"}
	for _, v := range strs {
		fmt.Println(frequencySort(v))
	}
}
