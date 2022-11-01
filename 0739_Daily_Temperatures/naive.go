package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

func dailyTemperatures(temperatures []int) []int {
	mono := datastruct.NewStack([]int{})
	rslt := make([]int, len(temperatures))

	for i := len(temperatures) - 1; i >= 0; i-- {
		for !mono.IsEmpty() && temperatures[i] >= temperatures[mono.Peek()] {
			mono.Pop()
		}

		if !mono.IsEmpty() {
			rslt[i] = mono.Peek() - i
		}

		mono.Push(i)
	}

	return rslt
}

func main() {
	temps := [][]int{[]int{73,74,75,71,69,72,76,73}, []int{30,40,50,60}}
	for i := range temps {
		fmt.Println(dailyTemperatures(temps[i]))
	}
}
