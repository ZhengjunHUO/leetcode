package main

import (
	"testing"
)

func TestCoinChange(t *testing.T) {
	coins := [][]int{[]int{1,2,5}, []int{2}, []int{1}, []int{1}, []int{1}}
	amounts := []int{11,3,0,1,2}
	wants := []int{3,-1,0,1,2}

	for i := range coins {
		rslt := coinChange(coins[i], amounts[i])
		if wants[i] != rslt {
			t.Errorf("coinChange(%v, %d) return %d but want %d", coins[i], amounts[i], rslt, wants[i])
		}
	}
}
