package main

import (
	"testing"
)

func TestChange(t *testing.T) {
	coins := [][]int{[]int{1,2,5}, []int{2}, []int{10}}
	as := []int{5,3,10}
	want := []int{4,0,1}

	for i := range coins {
		rslt := change(as[i], coins[i])
		if rslt != want[i] {
			t.Errorf("change(%d, %v) return %d but want %d", as[i], coins[i], rslt, want[i])
		}
	}
}
