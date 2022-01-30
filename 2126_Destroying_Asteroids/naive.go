package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func asteroidsDestroyed(mass int, asteroids []int) bool {
	pq := godtype.NewPQ([]int{}, []int{}, true)

	for i := range asteroids {
		if mass >= asteroids[i] {
			mass += asteroids[i]
		// 如果暂时不能破坏小行星，将其加入一个最小堆
		}else{
			pq.Push(asteroids[i], asteroids[i])
		}
	}

	for pq.Size() > 0 {
		// 如果质量小于最小堆的堆顶，那堆中剩下的小行星也肯定不能被消灭
		if pq.Peek().(int) > mass {
			return false
		}

		mass += pq.Pop().(int)
	}

	// 如果处理完了堆则表示所有小行星都被消灭
	return true
}

func main() {
	ast := [][]int{[]int{3,9,19,5,21}, []int{4,9,23,4}}
	mass := []int{10,5}

	for i := range ast {
		fmt.Println(asteroidsDestroyed(mass[i], ast[i]))
	}
}
