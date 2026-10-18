package main

import "fmt"

func BinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if target == nums[mid] {
			return mid
		}

		if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	nums := make([]int, 0, 100)

	for i := 1; i <= 100; i++ {
		nums = append(nums, i)
	}

	cases := []struct {
		nums   []int
		target int
		want   int
	}{
		{nums, 73, 72},
		{nums, 101, -1},
		{[]int{}, 10, -1},
		{[]int{1}, 1, 0},
		{[]int{1}, 2, -1},
	}

	for i, tc := range cases {
		got := BinarySearch(tc.nums, tc.target)
		status := "OK"
		if got != tc.want {
			status = "FAIL"
		}
		fmt.Printf("%-4s %d: got: %d, want: %d\n", status, i, got, tc.want)
	}
}
