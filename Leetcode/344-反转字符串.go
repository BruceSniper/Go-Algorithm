package Leetcode

func reverseString(s []byte) {
	n := len(s)
	head, tail := 0, n-1
	for head < tail {
		s[head], s[tail] = s[tail], s[head]
		head++
		tail--
	}
}
