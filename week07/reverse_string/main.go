package main

import "fmt"

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	cases := []struct {
		s    string
		want string
	}{
		{"abc", "cba"},
		{"Привет", "тевирП"},
		{"", ""},
		{"a", "a"},
		{"Go, мир!", "!рим ,oG"},
		{"🙂ok", "ko🙂"},
	}

	for i, tc := range cases {
		got := Reverse(tc.s)
		status := "OK"
		if got != tc.want {
			status = "FAIL"
		}
		fmt.Printf("%-4s %d: got: %q, want: %q\n", status, i, got, tc.want)
	}
}
