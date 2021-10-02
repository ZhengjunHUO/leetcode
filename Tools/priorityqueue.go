package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func main() {
	//values := []string{"banana", "apple", "pear"}
	values := []int{500, 404, 1000}
	prios := []int{3, 2, 4}

	pq := godtype.NewPQ(values, prios, false)
	pq.Push("orange", 1)
	
	fmt.Println("Priority queue: ")
	for _,v := range pq.Data {
		fmt.Printf("%.2d:%v [index: %v]\n", v.Priority, v.Value, v.Index)
	}
	fmt.Println("")

	/*
	fmt.Println("Update elem.")
	pq.Update("orange", 10)

	fmt.Println("Priority queue: ")
	for _,v := range pq.Data {
		fmt.Printf("%.2d:%v [index: %v]\n", v.Priority, v.Value, v.Index)
	}
	fmt.Println("")
	*/

	fmt.Println("Peek: ", pq.Peek())

	fmt.Println("Pop:")
	for pq.Data.Len() > 0 {
		fmt.Println(pq.Pop())
	}
}
