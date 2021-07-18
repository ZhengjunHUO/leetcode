package main

import (
	"fmt"
)

func topKFrequent(words []string, k int) []string {
	dict := make(map[string]int)
	for _,v := range words {
		dict[v] += 1
	}

	bucket := make([][]string, len(words))
	for k, v := range dict {
		if bucket[v] != nil {
			l := len(bucket[v])
			for i:=0; i<l; i++ {
				if k < bucket[v][i] {
					bucket[v] = append(bucket[v], "")
					copy(bucket[v][i+1:], bucket[v][i:l])
					bucket[v][i] = k
					break
				}else{
					if i == l-1 {
						bucket[v] = append(bucket[v], k)
					}
				}
			}	
		}else{
			bucket[v] = []string{k}
		}
	}

	rslt := make([]string, 0, k)
	loop: for i:=len(bucket)-1; i>=0; i-- {
		if bucket[i] != nil {
			for j:=0; j<len(bucket[i]); j++ {
				rslt = append(rslt, bucket[i][j])
				if len(rslt) == k {
					break loop
				}
			}
		}
	}

	return rslt
}

func main() {
	strs := [][]string{[]string{"i", "love", "leetcode", "i", "love", "coding"}, []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}}
	ks := []int{2, 4}

	for i := range strs {
		fmt.Println(topKFrequent(strs[i], ks[i]))
	}
}
