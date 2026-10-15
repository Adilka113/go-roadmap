package main

import "fmt"

func removeDuplicates(s []int) []int {
	seen := make(map[int]struct{}, len(s))
	res := make([]int, 0, len(s))

	for _, elem := range s {
		if _, ok := seen[elem]; !ok {
			seen[elem] = struct{}{}
			res = append(res, elem)
		}
	}

	return res
}

func main() {
	s := []int{1, 2, 2, 3, 1, 4}

	fmt.Println(removeDuplicates(s))
}
