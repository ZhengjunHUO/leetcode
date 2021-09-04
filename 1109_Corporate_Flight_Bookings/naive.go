package main

import (
	"fmt"
)

func corpFlightBookings(bookings [][]int, n int) []int {
	diff := make([]int, n)

	for i := range bookings {
		diff[bookings[i][0]-1] += bookings[i][2]
		if bookings[i][1] < n {
			diff[bookings[i][1]] -= bookings[i][2]
		} 
	}	

	for i := 1; i < len(diff); i++ {
		diff[i] += diff[i-1]	
	}

	return diff
}

func main() {
	bks := [][][]int{[][]int{[]int{1,2,10},[]int{2,3,20},[]int{2,5,25}}, [][]int{[]int{1,2,10}, []int{2,2,15}}}
	ns := []int{5,2}
	for i := range bks {
		fmt.Println(corpFlightBookings(bks[i], ns[i]))
	}
}
