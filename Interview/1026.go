package main

import (
	"fmt"
	"math"
)

func main() {
	var str1, str2 string
	fmt.Scanf("%s", &str1)
	fmt.Scanf("%s", &str2)
	fmt.Println(minSubStr(str1, str2))
}

func minSubStr(str1 string, str2 string) int {
	var result int
	if len(str2) > len(str1) {
		return 0
	}
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(str2); i++ {
		ori[str2[i]]++
	}
	sLen := len(str1)
	Len := math.MaxInt32
	ansL, ansR := -1, -1
	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && ori[str1[r]] > 0 {
			cnt[str1[r]]++
		}
		for check() && l <= r {
			if r-l+1 < Len {
				Len = r - l + 1
				ansL, ansR = l, l+Len
			}
			if _, ok := ori[str1[l]]; ok {
				cnt[str1[l]] -= 1
			}
			l++
		}
	}

	if ansL == -1 {
		return 0
	}

	result = len(str1[ansL:ansR])

	return result
}
