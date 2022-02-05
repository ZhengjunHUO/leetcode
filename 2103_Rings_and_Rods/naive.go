package main

import (
	"fmt"
)

// 以标杆编号为key, value为可能包含'R', 'G', 'B'的set
func countPoints(rings string) int {
	dict := make(map[int]map[byte]struct{})
	ret := 0

	for i := 0; i < len(rings); i += 2 {
		if _, ok := dict[int(rings[i+1])]; !ok {
			dict[int(rings[i+1])] = make(map[byte]struct{})
		}
		dict[int(rings[i+1])][rings[i]] = struct{}{}
	}

	for _, v := range dict {
		if len(v) == 3 {
			ret += 1
		}	
	}
	
	return ret
}

func main() {
	rings := []string{"B0B6G0R6R0R6G9", "B0R0G0R9R0B0G0", "G4"}
	for i := range rings {
		fmt.Println(countPoints(rings[i]))
	}
}
