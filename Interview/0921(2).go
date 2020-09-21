package main

import "fmt"

func main() {
	var s string
	fmt.Scanf("%s", &s)
	fmt.Println(LongestSubString(s))
}

func LongestSubString(s string) string {
	var res int
	low, fast := 0, 1
	switch {
	case len(s) < 2:
		if len(s) == 0 {
			return ""
		}
		if len(s) == 1 {
			return s
		}
	default:

		for i := 0; i < len(s)-1; i++ {
			for j := low; j < fast; j++ {
				if s[j] == s[fast] {
					low = j + 1
				}
			}
			res = Max(res, fast-low+1)
			fast++
		}
	}
	return s[low:fast]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
