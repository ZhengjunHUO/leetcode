package main

import (
	"fmt"
	"math"
)

func stationsNeeded(stations []int, dist float64) int {
	rslt := 0
	for i := 0; i < len(stations)-1; i++ {
		rslt += int(math.Ceil(float64(stations[i+1] - stations[i])/dist) - 1.0)
	}
	return rslt
}

func minmaxGasDist(stations []int, K int) float64 {
	var l, r, m float64 = 0, float64(stations[len(stations)-1] - stations[0]), 0
	for l + 0.000001 <= r {
		m = l + (r-l)/2
		if stationsNeeded(stations, m) <= K {
			r = m
		}else{
			l = m
		}
	}

	return m
}

func main() {
	fmt.Println(minmaxGasDist([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 9))
} 
