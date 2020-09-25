package Leetcode

func longestPalindrome(s string) string {
	length := len(s)
	if length == 0 {
		return ""
	}
	var start, end, res = 0, 1, 1
	for i := 0; i < length; i++ {
		for j := length - 1; j > i; j-- {
			if s[i] == s[j] {
				var ok = true
				for p1, p2 := i+1, j-1; p1 < p2; p1, p2 = p1+1, p2-1 {
					if s[p1] != s[p2] {
						ok = false
						break
					}
				}
				if ok && res < j-i+1 {
					res = j - i + 1
					start, end = i, j+1
				}
			}
		}
	}
	return s[start:end]
}
