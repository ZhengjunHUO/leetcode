package main

import (
	"testing"
)

func TestMinimumTotal(t *testing.T) {
	triangles := [][][]int{[][]int{[]int{2},[]int{3,4},[]int{6,5,7},[]int{4,1,8,3}}, [][]int{[]int{-10}}}
	expected := []int{11,-10}

	for i := range triangles {
		if rslt := minimumTotal(triangles[i]); rslt != expected[i] {
			t.Errorf("minimumTotal(%v) returned %d but expect %d", triangles[i], rslt, expected[i])
		}
	}
}
