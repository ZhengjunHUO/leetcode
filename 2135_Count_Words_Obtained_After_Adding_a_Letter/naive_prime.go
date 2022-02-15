package main

import (
	"fmt"
	"github.com/ZhengjunHUO/gtoolkit"
)

func wordCount(startWords []string, targetWords []string) int {
	primeDict := gtoolkit.GetLowercasePrimeDict()

	startDict := make(map[int][]int)
	for i := range startWords {
		count, val := 0, 1
		for j := range startWords[i] {
			val *= primeDict[int(startWords[i][j])]
			count++
		}

		if _, ok := startDict[count]; !ok {
			startDict[count] = []int{val}
		}else{
			startDict[count] = append(startDict[count], val)
		}
	}	
	//fmt.Println(startDict)

	var ret int
	for i := range targetWords {
		count, val := 0, 1
		for j := range targetWords[i] {
			val *= primeDict[int(targetWords[i][j])]
			count++
		}
	
		//fmt.Println("count, val:", count, val)
		if v, ok := startDict[count-1]; ok {
			for k := range v {
				if val % v[k] == 0 {
					if temp := val/v[k]; v[k]%temp != 0 {
						//fmt.Printf("targetWords[%d] ok\n", i)
						ret++
					}
				}
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
