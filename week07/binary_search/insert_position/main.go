package main

import "fmt"

func InsertSearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	nums := []int{1, 3, 5, 7, 9}

	cases := []struct {
		nums   []int
		target int
		want   int
	}{
		{nums, 7, 3},
		{nums, 4, 2},
		{nums, 0, 0},
		{nums, 10, 5},
		{[]int{}, 10, 0},
	}

	for i, tc := range cases {
		got := InsertSearch(tc.nums, tc.target)
		status := "OK"
		if got != tc.want {
			status = "FAIL"
		}
		fmt.Printf("%-4s %d: got: %d, want: %d\n", status, i, got, tc.want)
	}
}
