package main

import "fmt"

func FindString(s string) int {
	m := map[byte]int{}
	n := len(s)
	l, ans := -1, 0
	for j := 0; j < n; j++ {
		if j != 0 {
			delete(m, s[j-1])
		}
		for l+1 < n && m[s[l+1]] == 0 {
			m[s[l+1]]++
			l++
		}
		ans = max(ans, l-j+1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	var str string
	fmt.Scanf("%s", &str)
	fmt.Println(FindString(str))
}
