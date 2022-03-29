package main

import (
	"fmt"
	"strings"
)

func decodeCiphertext(encodedText string, rows int) string {
	if rows == 1 {
		return encodedText
	}

	ret := []byte{}
	cols := len(encodedText)/rows
	var curr_pos, start_col, offset int
	for start_col < cols {
		curr_pos = start_col + cols*offset + offset
		if curr_pos > len(encodedText)-1 {
			start_col++
			offset=0
			continue
		}
		ret = append(ret, encodedText[curr_pos])
		offset++
	}

	return strings.TrimRight(string(ret), " ")
}

func main() {
	encodedText := []string{"ch   ie   pr", "iveo    eed   l te   olc", "coding"}
	rows := []int{3, 4, 1}

	for i := range rows {
		fmt.Printf("[%s]\n", decodeCiphertext(encodedText[i], rows[i]))
	}
}
