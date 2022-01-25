package main

import (
	"testing"
)

func TestCapitalizeTitle(t *testing.T) {
	titles := []string{"capiTalIze tHe titLe", "First leTTeR of EACH Word", "i lOve leetcode"}
	wanted := []string{"Capitalize The Title", "First Letter of Each Word", "i Love Leetcode"}
	for i := range titles {
		if rslt := capitalizeTitle(titles[i]); rslt != wanted[i] {
			t.Errorf("capitalizeTitle(%s) get \"%s\" but wait for \"%s\"", titles[i], rslt, wanted[i])
		}
	}
}
