package main

import (
	"fmt"
)

func findPoisonedDuration(timeSeries []int, duration int) int {
	rslt := 0

	for i := range timeSeries {
		if i == len(timeSeries) - 1 {
			rslt += duration
			break
		}

		if timeSeries[i] + duration < timeSeries[i+1] {
			rslt += duration
		}else{
			rslt += timeSeries[i+1] - timeSeries[i]
		}
	}

	return rslt
}

func main() {
	ts := [][]int{[]int{1, 4}, []int{1, 2}, []int{1, 3, 9}}
	d := []int{2, 2, 4}

	for i := range ts {
		fmt.Println(findPoisonedDuration(ts[i], d[i]))
	}
}
