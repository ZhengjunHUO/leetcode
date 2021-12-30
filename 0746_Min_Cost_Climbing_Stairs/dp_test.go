package main

import (
	"testing"
)

func TestMinCostClimbingStairs(t *testing.T) {
	cost := [][]int{[]int{10,15,20}, []int{1,100,1,1,1,100,1,1,100,1}}
	wanted := []int{15, 6}

	for i := range cost {
		if rslt := minCostClimbingStairs(cost[i]); rslt != wanted[i] {
			t.Errorf("Execute minCostClimbingStairs(%v) returned %d but expect %d", cost[i], rslt, wanted[i])
		}
	}
}
