package main

import (
	"fmt"
	"math"
)

/*
  一个灯最后保持亮的状态需要被拨动奇数次
  根据规则灯是否被开关取决于其分解的因数
  如9可以分解为1*9和3*3(只算一个3)，故9最终会亮
  值为某个数的平方的灯最后都是亮的，对n求sqrt就能知道1-n内有几个这样的数
*/
func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n)))
}

func main() {
	ns := []int{3,0,1,10}
	for i := range ns {
		fmt.Println(bulbSwitch(ns[i]))
	}
}
