package main

import (
	"fmt"
)

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	i, j, rslt := 0, 0, [][]int{}

	for i < len(firstList) && j < len(secondList) {
		if firstList[i][0] <= secondList[j][1] && secondList[j][0] <= firstList[i][1] {
			begin, end := 0, 0
			if firstList[i][0] < secondList[j][0] {
				begin = secondList[j][0]
			}else{
				begin = firstList[i][0]
			}

			if firstList[i][1] < secondList[j][1] {
				end = firstList[i][1]
			}else{
				end = secondList[j][1]
			}

			rslt = append(rslt, []int{begin, end})
		}

		if firstList[i][1] < secondList[j][1] {
			i++
		}else{
			j++
		}
	}

	return rslt
}

func main() {
	firstList := [][]int{[]int{0,2},[]int{5,10},[]int{13,23},[]int{24,25}}
	secondList := [][]int{[]int{1,5},[]int{8,12},[]int{15,24},[]int{25,26}}
	fmt.Println(intervalIntersection(firstList, secondList))

	fmt.Println(intervalIntersection([][]int{[]int{1,3},[]int{5,9}}, [][]int{}))
	fmt.Println(intervalIntersection([][]int{}, [][]int{[]int{4,8},[]int{10,12}}))
	fmt.Println(intervalIntersection([][]int{[]int{1,7}}, [][]int{[]int{3,10}}))
}
