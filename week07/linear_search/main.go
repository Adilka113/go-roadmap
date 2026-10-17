package main

import "fmt"

func LinearSearch(nums []int, target int) int {
	for i, elem := range nums {
		if elem == target {
			return i
		}
	}
	return -1
}

func main() {
	cases := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{10, 20, 30, 40, 50}, 30, 2},
		{[]int{}, 10, -1},
		{[]int{10, 20, 30}, 40, -1},
	}

	for i, tc := range cases {
		got := LinearSearch(tc.nums, tc.target)
		status := "OK"
		if got != tc.want {
			status = "FAIL"
		}
		fmt.Printf("%-4s %d: got : %d, want: %d\n", status, i, got, tc.want)
	}
}
