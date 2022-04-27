package main

import (
	"fmt"
)

func finalValueAfterOperations(operations []string) int {
	dict := map[string]int{"--X": -1, "X--": -1, "++X": 1, "X++": 1}
	var ret int

	for i := range operations {
		ret += dict[operations[i]]
	}

	return ret
}

func main() {
	ops := [][]string{[]string{"--X","X++","X++"}, []string{"++X","++X","X++"}, []string{"X++","++X","--X","X--"}}
	for i := range ops {
		fmt.Println(finalValueAfterOperations(ops[i]))
	}
}
