package main

import (
	"testing"
)

func TestGetMaximumGenerated(t *testing.T){
	n := []int{7,2,3}
	expected := []int{3,1,2}

	for i := range n {
		if rslt := getMaximumGenerated(n[i]); rslt != expected[i] {
			t.Errorf("getMaximumGenerated(%d) returned %d but expect %d", n[i], rslt, expected[i])
		}
	}
}
