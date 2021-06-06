package main

import (
	"fmt"
)

func doSomething(nums []int)  {
	fmt.Printf("In son, before append: %p\n", &nums)
	fmt.Printf("In son, before append: %p\n", &nums[0])
	nums[2] = 4
	fmt.Println(nums)
	nums = append(nums, 100)
	nums = append(nums, 200)
	fmt.Println(nums)
	fmt.Printf("In son, after append: %p\n", &nums)
	fmt.Printf("In son, after append: %p\n", &nums[0])
}

func main() {
	nums := []int{0,1,0,3,12}
	fmt.Printf("In main: %p\n", &nums)
	fmt.Printf("In main: %p\n", &nums[0])
	doSomething(nums)
	fmt.Println(nums)
}
