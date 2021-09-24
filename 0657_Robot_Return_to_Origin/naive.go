package main

import (
	"fmt"
)

func judgeCircle(moves string) bool {
	dict := map[byte][2]int{byte('U'): [2]int{0,1}, byte('D'): [2]int{0,-1}, byte('L'): [2]int{-1,0}, byte('R'): [2]int{1,0}}
	x, y := 0, 0
	for i := range moves {
		curr := dict[moves[i]]
		x, y = x+curr[0], y+curr[1]
	}

	return (x == 0) && (y == 0)
}

func main() {
	str := []string{"UD", "LL", "RRDD", "LDRRLRUULR", "LDRRLRUULD"}
	for i := range str {
		fmt.Println(judgeCircle(str[i]))
	}
}
