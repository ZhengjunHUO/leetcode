package main

import (
	"fmt"
)

func exist(board [][]byte, word string) bool {
	return false 
}

func main() {
	board := [][]byte{[]byte{'A', 'B', 'C', 'E'}, []byte{'S', 'F', 'C', 'S'}, []byte{'A', 'D', 'E', 'E'}}
	str := []string{"ABCCED", "SEE", "ABCB"}

	for i := range str {
		fmt.Println(exist(board, str[i]))
	}
}
