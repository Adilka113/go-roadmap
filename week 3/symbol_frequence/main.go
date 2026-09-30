package main

import (
	"fmt"
	"unicode"
)

func symbolFreq(s string) map[string]int {
	res := make(map[string]int, len(s))
	for _, symbol := range s {
		if !unicode.IsLetter(symbol) {
			continue
		}
		res[string(unicode.ToLower(symbol))]++
	}
	return res
}

func main() {
	s := "hello world"
	m := symbolFreq(s)
	fmt.Println(m)
}
