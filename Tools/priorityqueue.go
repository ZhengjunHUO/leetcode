package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func main() {
	//values := []string{"banana", "apple", "pear"}
	values := []int{500, 404, 1000}
	prios := []int{3, 2, 4}

	pq := godtype.InitPQ(values, prios, false)
	(&pq).Insert("orange", 1)
	
	fmt.Println("Priority queue: ")
	for i := range pq {
		fmt.Printf("%.2d:%v [index: %v]\n", pq[i].Priority, pq[i].Value, pq[i].Index)
	}
	fmt.Println("")

	fmt.Println("Update elem.")
	(&pq).Update("orange", 10)

	fmt.Println("Priority queue: ")
	for i := range pq {
		fmt.Printf("%.2d:%v [index: %v]\n", pq[i].Priority, pq[i].Value, pq[i].Index)
	}
	fmt.Println("")

	fmt.Println("Peek: ", (&pq).Peek())

	fmt.Println("Pop:")
	for pq.Len() > 0 {
		fmt.Println((&pq).Pull())
	}
}
