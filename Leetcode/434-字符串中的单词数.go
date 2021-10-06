package Leetcode

import (
	"strings"
)

func countSegments(s string) int {
	count := strings.Fields(s)
	return len(count)
}

func countSegments2(s string) int {
	var count int = 0
	for i := 0; i < len(s); i++ {
		if (i == 0 || s[i-1] == ' ') && s[i] != ' ' {
			count++
		}
	}
	return count
}
