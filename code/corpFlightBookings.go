package main

import (
	"fmt"
	"sort"
)

func corpFlightBookings(bookings [][]int, n int) []int {
	nums := make([]int, n)
	for _, booking := range bookings {
		nums[booking[0]-1] += booking[2]
		if booking[1] < n {
			nums[booking[1]] -= booking[2]
		}
	}
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
	}
	sort.Ints(nums)
	return nums
}

func main() {
	bookings := [][]int{{1,2,10},{2,3,20},{2,5,25}}
	n := 5
	nums := corpFlightBookings(bookings,n)
	fmt.Println("nums:",nums)
}
