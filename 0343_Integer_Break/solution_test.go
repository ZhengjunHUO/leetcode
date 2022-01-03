package main

import (
	"testing"
)

func TestIntegerBreak(t *testing.T) {
	n := []int{2,10}
	expected := []int{1,36}

	for i := range n {
		if rslt := integerBreak(n[i]); rslt != expected[i] {
			t.Errorf("integerBreak(%d) returned %d but expect %d\n", n[i], rslt, expected[i])
		}
	}
}
