package main

import "fmt"

func Subslice(s []int, start, end int) []int {
	if start > end || len(s) < end || start < 0 {
		return nil
	}

	res := make([]int, end-start)

	copy(res, s[start:end])

	return res

}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	a := Subslice(s, 2, 7)

	b := Subslice(s, 7, 2)

	fmt.Println(a, b)
}
