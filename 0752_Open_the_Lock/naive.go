package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func rotate(str string, idx int, isUp bool) string {
	// '0': 48, '9': 57
	var changed byte
	if isUp {
		if str[idx] == 57 {
			changed = 48
		}else{
			changed = str[idx] + 1
		}
	}

	if !isUp {
		if str[idx] == 48 {
			changed = 57
		}else{
			changed = str[idx] - 1
		}
	}
	
	return str[:idx] + string(changed) + str[idx+1:]
}

func openLock(deadends []string, target string) int {
	visited := make(map[string]bool)
	for _,v := range deadends {
		visited[v] = true
	}

	if _, ok := visited["0000"]; ok {
		return -1
	}

	q := godtype.NewQueue()
	q.Push("0000")
	visited["0000"] = true
	step := 0

	for !q.IsEmpty() {
		size := q.Size()
		for i:=0; i<size; i++ {
			curr := q.Pop().(string)
			if curr == target {
				return step
			}			

			for i:=0; i<4; i++ {
				up := rotate(curr, i, true)
				if _, ok := visited[up]; !ok {
					q.Push(up)
					visited[up] = true
				}

				down := rotate(curr, i, false)
				if _, ok := visited[down]; !ok {
					q.Push(down)
					visited[down] = true
				}
			}
		}

		step++
	}

	return -1
}

func main() {
	deads := [][]string{[]string{"0201","0101","0102","1212","2002"}, []string{"8888"}, []string{"8887","8889","8878","8898","8788","8988","7888","9888"}, []string{"0000"}}
	targets := []string{"0202", "0009", "8888", "8888"}
	for i := range deads {
		fmt.Println(openLock(deads[i], targets[i]))
	}
}
