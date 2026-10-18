package main

import (
	"fmt"
	"unicode"
)

func IsPalindrome(s string) bool {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		if r[i] != r[j] {
			return false
		}
	}
	return true
}

func IsPalindromeClean(s string) bool {
	res := make([]rune, 0, len(s))
	for _, v := range s {
		if unicode.IsLetter(v) || unicode.IsDigit(v) {
			res = append(res, unicode.ToLower(v))
		}
	}
	return IsPalindrome(string(res))
}

func main() {
	cases := []struct {
		s         string
		want      bool
		wantClean bool
	}{
		{"level", true, true},
		{"racecar", true, true},
		{"hello", false, false},
		{"", true, true},
		{"a", true, true},
		{"ab", false, false},
		{"казак", true, true},
		{"Level", false, true},
		{"0P", false, false},
		{".,!", false, true},
		{"Panama", false, false},
		{"race a car!", false, false},
	}

	for _, tc := range cases {
		got := IsPalindrome(tc.s)
		gotClean := IsPalindromeClean(tc.s)
		status := "OK"
		statusClean := "OK"
		if got != tc.want {
			status = "FAIL"
		}
		if gotClean != tc.wantClean {
			statusClean = "FAIL"
		}
		fmt.Printf("%-4s %q: got: %t want: %t\n", status, tc.s, got, tc.want)
		fmt.Printf("%-4s %q: got: %t want: %t\n", statusClean, tc.s, gotClean, tc.wantClean)
	}
}
