package Leetcode

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	prefix := MinStr(strs)
	for _, v := range strs {
		for strings.Index(v, prefix) != 0 {
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}

func MinStr(strs []string) string {
	minStr := strs[0]
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < len(minStr) {
			minStr = strs[i]
		}
	}
	return minStr
}
