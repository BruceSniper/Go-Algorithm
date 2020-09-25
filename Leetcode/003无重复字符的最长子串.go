package Leetcode

import "strings"

func lengthOfLongestSubstring(s string) int {
	num := len(s)
	start, end := 0, 0
	for i := 0; i < num; i++ {
		index := strings.Index(s[start:i], string(s[i]))
		if index == -1 {
			if i+1 > end {
				end = i + 1
			}
		} else {
			start += index + 1
			end += index + 1
		}
	}
	return end - start
}
