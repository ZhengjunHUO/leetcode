package main

import (
	"testing"
)

func TestHashtable(t *testing.T) {
	pattern, s := []string{"abba", "abba", "aaaa"}, []string{"dog cat cat dog", "dog cat cat fish", "dog cat cat dog"}
	expected := []bool{true, false, false}
	for i := range pattern {
		if rslt := wordPattern(pattern[i], s[i]); rslt != expected[i] {
			t.Errorf("wordPattern(%s, %s) return %v, expected %v\n", pattern[i], s[i], rslt, expected[i])
		}
	}
}
