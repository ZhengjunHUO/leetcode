package main

import (
	"fmt"
)

// 因为每次可以拿走1/2/3个石子，所以只需要确保留给对手4以及4的倍数的石子即可获胜
// 如果开局有4的倍数个石子，由于玩家是先手所以必输
func canWinNim(n int) bool {
	return n%4 != 0
}

func main() {
	ns := []int{4,1,2}
	for i := range ns {
		fmt.Println(canWinNim(ns[i]))
	}	
}
