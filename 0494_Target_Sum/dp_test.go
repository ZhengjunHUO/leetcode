package main

import (
	"testing"
)

func TestFindTargetSumWays(t *testing.T) {
	rslt := findTargetSumWays([]int{1,1,1,1,1}, 3)
	if rslt != 5 {
		t.Errorf("findTargetSumWays([]int{1,1,1,1,1}, 3) return %d want %d", rslt, 5)
	}
}
