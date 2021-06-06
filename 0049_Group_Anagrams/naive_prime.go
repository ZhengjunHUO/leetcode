package main

import (
	"fmt"
)

func lowercasePrimeDict() map[int]int {
	m := make(map[int]int)
	primes := []int{2,3,5,7,11,13,17,19,23,29,31,37,41,43,47,53,59,61,67,71,73,79,83,89,97,101}

	for i:=0; i<26; i++ {
		m[i+97] = primes[i]
	} 

	return m
}

func groupAnagrams(strs []string) [][]string {
	dict :=  lowercasePrimeDict()
	reg := make(map[int][]string)
	rslt := [][]string{}

	for i := range(strs) {
		prod := 1
		for j := range strs[i] {
			prod *= dict[int(strs[i][j])]
		}

		reg[prod] = append(reg[prod], strs[i])
	}

	for _,v := range reg {
		rslt = append(rslt, v)
	}

	return rslt
}

func main() {
	strs := []string{"eat","tea","tan","ate","nat","bat"}
	fmt.Println(groupAnagrams(strs))

	strs = []string{}
	fmt.Println(groupAnagrams(strs))

	strs = []string{"a"}
	fmt.Println(groupAnagrams(strs))
}
